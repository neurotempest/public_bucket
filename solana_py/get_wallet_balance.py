from solana.keypair import Keypair
from solana.rpc.api import Client

with open('secret.key') as f:
  secret_key = bytes.fromhex(f.read())
kp = Keypair().from_secret_key(secret_key)

cli = Client("https://api.devnet.solana.com")

resp = cli.get_balance(kp.public_key)

print('Balance for address', kp.public_key, ':', resp['result']['value'])

