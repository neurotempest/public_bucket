from pycardano import *
from pathlib import Path

# https://testnet.cardanoscan.io/pool/48f2c367cfe81cac6687c3f7c26613edfe73cd329402aa5cf493bb61
POOL_ID="48f2c367cfe81cac6687c3f7c26613edfe73cd329402aa5cf493bb61"
network = Network.TESTNET

psk = PaymentSigningKey.load("payment.skey")
ssk = StakeSigningKey.load("stake.skey")

pvk = PaymentVerificationKey.from_signing_key(psk)
svk = StakeVerificationKey.from_signing_key(ssk)

address = Address(pvk.hash(), svk.hash(), network)
bf_key = Path('blockfrost.key').read_text().rstrip()
context = BlockFrostChainContext(bf_key, network)

utxos = context.utxos(str(address))

builder = TransactionBuilder(context)
# stake everything
builder.add_input_address(address)
stake_credential = StakeCredential(
        svk.hash()
)
stake_registration = StakeRegistration(stake_credential)
pool_hash = PoolKeyHash(bytes.fromhex(POOL_ID))
stake_delegation = StakeDelegation(stake_credential, pool_keyhash=pool_hash)

builder.certificates = [stake_registration, stake_delegation]
signed_tx = builder.build_and_sign(
    [ssk, psk],
    address,
)

print("############### Submitting transaction ###############")
context.submit_tx(signed_tx.to_cbor())
print("Broadcast tx:", signed_tx.id)
