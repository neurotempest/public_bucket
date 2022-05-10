from pycardano import *

payment_key_pair = PaymentKeyPair.generate()
payment_key_pair.signing_key.save("payment.skey")

stake_key_pair = StakeKeyPair.generate()
stake_key_pair.signing_key.save("stake.skey")

network = Network.TESTNET

address = Address(
  payment_part=payment_key_pair.verification_key.hash(),
  staking_part=stake_key_pair.verification_key.hash(),
  network=network,
)

print("Created keys for address:", address)
