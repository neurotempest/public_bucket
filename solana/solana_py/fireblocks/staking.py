import json
import base64
from fireblocks_sdk import *
from solana.rpc.api import Client
from solana.transaction import Transaction, TransactionInstruction
from solana.system_program import AdvanceNonceParams, nonce_advance
from solana.publickey import PublicKey
from solana.keypair import Keypair

import stake_program as ssp


VAULT_ID=""
ASSET="SOL_TEST"

with open('secret.key', 'r') as f:
    apiSecret = f.read()

with open('api.key', 'r') as f:
    apiKey = f.readline().strip()

with open('nonceAuth/testNonceAccounts.json', 'r') as f:
  test_nonce_accounts = json.load(f)

fb_sdk = FireblocksSDK(apiSecret, apiKey)

# use Solana Devnet
sol_sdk = Client("https://api.devnet.solana.com")

# if vault id is not set, try to get it from Fireblocks
if not VAULT_ID :
    res = fb_sdk.get_vault_accounts(name_suffix='SOL Staking')
    if not len(res):
        print('Staking Vaut not found')
        exit(-1)

    VAULT_ID=res[0]['id']

# find deposite address for the Solana Assset
solAsset = fb_sdk.get_vault_account_asset(VAULT_ID, ASSET)
deposite_address = fb_sdk.get_deposit_addresses(VAULT_ID, ASSET)[0]['address']

# get public key
res=fb_sdk.get_public_key_info_for_vault_account(ASSET, VAULT_ID, 2, 0, True)
stake_address=bytes.fromhex(res['publicKey'])
# make the key to be used in Solana stake_public_key  = PublicKey(stake_address)

# get nonce accounts
nonce_account = Keypair.from_secret_key(bytes(test_nonce_accounts['accounts'][2]))
authority_keypair = Keypair.from_secret_key(bytes(test_nonce_accounts['authority']))
nonce_account_info = sol_sdk.get_account_info(nonce_account.public_key)
res=nonce_account_info['result']
nonceHash = base64.b64decode(res['value']['data'][0])


balance = float(solAsset['balance'])
if not balance:
    print(f'Balance is {balance} Deposit some SOL at {deposite_address}')

print(dir(TransactionInstruction))
tx = Transaction()
tx.add(
        nonce_advance(
            AdvanceNonceParams(
                nonce_pubkey = nonce_account.public_key,
                authorized_pubkey = authority_keypair.public_key
                )
        )
    )
tx.add(
        ssp.create_stake_account(
            ssp.CreateStakeAccountParams(
                from_pubkey = stake_address,
                stake_pubkey = stake_address,
                lamports = 200)
            )
    )

print(tx)
