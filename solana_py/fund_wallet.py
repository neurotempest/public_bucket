from solana.keypair import Keypair
from solana.rpc.api import Client

with open('secret.key') as f:
  secret_key = bytes.fromhex(f.read())
kp = Keypair().from_secret_key(secret_key)

cli = Client("https://api.devnet.solana.com")

resp = cli.request_airdrop(
  kp.public_key,
  1000000000,
)

funding_tx_id = resp['result']

print('Funding address:', kp.public_key)
print('Broadcast tx:', funding_tx_id)
