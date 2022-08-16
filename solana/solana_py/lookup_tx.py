from solana.rpc.api import Client
import json


cli = Client("https://api.devnet.solana.com")

resp = cli.get_transaction('5LMQKVMVjVmVZqNNwAPaccSEgegwnDuprSkc73EkonAqTqgeAMVy5tdA7xqY7koZxHGEL6mERChEmG4SN41T6ra8')

print(json.dumps(resp))
