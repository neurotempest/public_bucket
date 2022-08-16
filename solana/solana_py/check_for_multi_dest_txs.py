from solana.rpc.api import Client
from solana.exceptions import SolanaRpcException
import time
import json

blocks_to_scan = 50

output_dir='temp_blocks'

duplicate_output='duplicates.log'

def make_block_filepath(block_height):
  return output_dir + '/' + str(block_height) + '.json'

def read_block_from_file(block_height):
  with open(make_block_filepath(block_height), 'r') as f:
    return json.load(f)

def get_block_from_cli_and_write_to_file(cli, block_height):

  time.sleep(1)
  resp = cli.get_block(block_height)

  with open(make_block_filepath(block_height), 'w') as f:
    json.dump(resp, f)

  return resp

def scan_block(cli, block_height):

  resp = {}
  try:
    resp = read_block_from_file(block_height)
  except FileNotFoundError:
    resp = get_block_from_cli_and_write_to_file(cli, block_height)


  if 'error' in resp and resp['error']['code'] == -32009:
    print(block_height, 'skipped')
    return

  txs = resp['result']['transactions']

  for tx in txs:
    instructions = tx['transaction']['message']['instructions']

    if len(instructions) > 1:

      first_transfer_acc = []

      for i in instructions:

        if i['programIdIndex'] == 2:

          if first_transfer_acc == [] and len(i['accounts']) == 2:
            first_transfer_acc = i['accounts']
          elif first_transfer_acc == i['accounts']:
            print('duplicate tx:', tx['transaction']['signatures'][0])
            with open(duplicate_output, 'w') as f:
              f.write(tx['transaction']['signatures'][0] + '\n')
            break


if __name__ == '__main__':

  cli = Client("https://api.mainnet-beta.solana.com")

  resp_bh = cli.get_block_height()['result']

  current_block = resp_bh
  while 1:
    try:
      print('scanning block', current_block)
      scan_block(cli, current_block)
      current_block -= 1
    except SolanaRpcException:
      print('got rate limit error - backing off for 20...')
      time.sleep(20)
      print('continuing...')
      continue


