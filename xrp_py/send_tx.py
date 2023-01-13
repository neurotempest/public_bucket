import json
import xrpl
from xrpl.wallet import Wallet

# Created from instructions here: https://xrpl.org/send-xrp.html

with open('creds.json') as creds_file:
  creds = json.load(creds_file)

creds_wallet = Wallet(seed=creds['secret'], sequence=creds['sequence'])

testnet_url = 'https://s.altnet.rippletest.net:51234'
client = xrpl.clients.JsonRpcClient(testnet_url)

my_payment = xrpl.models.transactions.Payment(
    account=creds_wallet.classic_address,
    amount=xrpl.utils.xrp_to_drops(890),
    destination='rHuv8LnNCvfAuCXfeapHMZhqEydgSq7imz',
    destination_tag=0,
)

signed_tx = xrpl.transaction.safe_sign_and_autofill_transaction(
        my_payment, creds_wallet, client)

try:
    tx_response = xrpl.transaction.send_reliable_submission(signed_tx, client)
except xrpl.transaction.XRPLReliableSubmissionException as e:
    exit(f"Submit failed: {e}")

print('Submitted Tx:', signed_tx.get_hash())
