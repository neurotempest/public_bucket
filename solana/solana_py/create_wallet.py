from solana.keypair import Keypair

kp = Keypair.generate()

with open('secret.key', 'w') as f:
  f.write(kp.secret_key.hex())

print("Created address:", kp.public_key)
