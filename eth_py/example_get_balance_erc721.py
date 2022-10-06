import json
from web3 import Web3, HTTPProvider
from web3.middleware import geth_poa_middleware

with open('key_info_creator.json') as key_info_file:
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

func_selector = web3.keccak(text='balanceOf(address)')[0:4].hex()
address='0x0EF8E223e4Df04E6A7483875b2728Af5f30FafF0'[2:].zfill(64)

signed_tx = account.sign_transaction({
  'from': account.address,
  'to': '0xA99dc99c0C6F3A126c5DcdCEA7c1d2B5B27a5c37',
  'nonce': nonce,
  'gas': 80000,
  'maxFeePerGas': web3.toWei(10, 'gwei'),
  'maxPriorityFeePerGas': web3.toWei(10, 'gwei'),
  'chainId': web3.eth.chain_id,
  'data': func_selector + address,
})

tx_hash = web3.eth.send_raw_transaction(signed_tx.rawTransaction)

print("Broadcast tx:", tx_hash.hex())


