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

nonce = web3.eth.get_transaction_count(account.address)

func_selector = web3.keccak(text='mint(address,uint256)')[0:4].hex()
to_address=account.address[2:].zfill(64)
amount=hex(web3.toWei(1, 'ether'))[2:].zfill(64)

signed_tx = account.sign_transaction({
  'from': account.address,
  'to': '0xFab46E002BbF0b4509813474841E0716E6730136',
  'nonce': nonce,
  'gas': 60000,
  'maxFeePerGas': web3.toWei(3, 'gwei'),
  'maxPriorityFeePerGas': web3.toWei(1, 'gwei'),
  'chainId': web3.eth.chain_id,
  'data': func_selector + to_address + amount,
})

tx_hash = web3.eth.send_raw_transaction(signed_tx.rawTransaction)

print("Broadcast tx:", tx_hash.hex())


