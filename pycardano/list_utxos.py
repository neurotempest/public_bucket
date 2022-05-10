from pycardano import *
from pathlib import Path

network = Network.TESTNET

psk = PaymentSigningKey.load("payment.skey")
ssk = StakeSigningKey.load("stake.skey")

pvk = PaymentVerificationKey.from_signing_key(psk)
svk = StakeVerificationKey.from_signing_key(ssk)

address = Address(pvk.hash(), svk.hash(), network)

bf_key = Path('blockfrost.key').read_text().rstrip()
context = BlockFrostChainContext(bf_key, network)

utxos = context.utxos(str(address))

print("Uxos for address:", address)

print(utxos)
