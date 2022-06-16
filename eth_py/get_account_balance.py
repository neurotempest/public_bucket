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

balance = web3.eth.get_balance(account.address)

print("Balance for account:", account.address)
print(balance)
