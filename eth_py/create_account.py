from web3 import Web3, HTTPProvider
from web3.middleware import geth_poa_middleware

with open('key_info.json') as key_info_file:
  key_info = json.load(key_info_file)

web3 = Web3(Web3.HTTPProvider(key_info['client']))
web3.middleware_onion.inject(geth_poa_middleware, layer=0)

if web3.isConnected():
  print('Connected!')

acc = web3.eth.account.create('Nobody expects the Spanish Inquisition!')

print(acc._address)
print(acc._private_key)

with open('secret.key', 'w') as f:
  f.write(kp.secret_key.hex())
