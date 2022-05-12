from solana.keypair import Keypair
from solana.rpc.api import Client
from solana.system_program import TransferParams, transfer
from solana.transaction import Transaction

with open('secret.key') as f:
  secret_key = bytes.fromhex(f.read())
kp = Keypair().from_secret_key(secret_key)

cli = Client("https://api.devnet.solana.com")

tx = Transaction()

tx.add(
  transfer(
    TransferParams(
      from_pubkey=kp.public_key,
      #to_pubkey='2XnzxSRiYPy5q72pYRUBDK8eH81BJbNbtN3eBCVbNeRL',
      #to_pubkey='12ZqUHFApp7QgZaddQr4ikxQWv8V3eXJjR8iyNLnZXZR',
      to_pubkey='135wbmL1D7gYDkdKkde515NaqCQGW8iUaVVhozR7SMrZ',
      lamports=21000000,
    ),
  ),
)

tx.add(
  transfer(
    TransferParams(
      from_pubkey=kp.public_key,
      to_pubkey='12ZqUHFApp7QgZaddQr4ikxQWv8V3eXJjR8iyNLnZXZR',
      lamports=22000000,
    ),
  ),
)

resp = cli.send_transaction(tx, kp)

print('Broadcast tx:', resp['result'])
