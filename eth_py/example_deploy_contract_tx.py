import json
from web3 import Web3, HTTPProvider
from web3.middleware import geth_poa_middleware

with open('key_info.json') as key_info_file:
  key_info = json.load(key_info_file)

web3 = Web3(Web3.HTTPProvider(key_info['client']))
web3.middleware_onion.inject(geth_poa_middleware, layer=0)

if not web3.isConnected():
  print('Failed to connect to ETH client!')
  exit(1)

with open(key_info['path']) as keyfile:
    encrypted_key = keyfile.read()
    private_key = web3.eth.account.decrypt(encrypted_key, key_info['pw'])

account = web3.eth.account.from_key(private_key)

nonce = web3.eth.get_transaction_count(account.address)

with open('contract.bin') as contract_file:
  contract_data = contract_file.read().replace('\n', '')

signed_tx = account.sign_transaction({
  'from': account.address,
  'nonce': nonce,
  'gas': 3000000,
  'maxFeePerGas': web3.toWei(10, 'gwei'),
  'maxPriorityFeePerGas': web3.toWei(10, 'gwei'),
  'chainId': web3.eth.chain_id,
  'data': contract_data,
})

tx_hash = web3.eth.send_raw_transaction(signed_tx.rawTransaction)

print("Broadcast tx:", tx_hash.hex())


