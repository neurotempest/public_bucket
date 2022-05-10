from pycardano import *
from pathlib import Path

network = Network.TESTNET

psk = PaymentSigningKey.load("payment.skey")
ssk = StakeSigningKey.load("stake.skey")

pvk = PaymentVerificationKey.from_signing_key(psk)
svk = StakeVerificationKey.from_signing_key(ssk)

address = Address(pvk.hash(), svk.hash(), network)

bf_key = Path('blockfrost.key').read_text().rstrip()
context = BlockFrostChainContext(bf_key, network)

utxos = context.utxos(str(address))

builder = TransactionBuilder(context)
builder.add_input_address(address)
builder.add_input(utxos[0])

builder.add_output(
  TransactionOutput(
    Address.from_primitive(
      #"addr_test1qp070yckjmp05u0n3xl90fs5ydfmz2avly7aexzjclz3rm6802xeya5yqh02crg8exj2k2y4639nfjg5xgvuh6ynyyfqpxw67w"
      "addr_test1qp07eka2e6s6kv973t4js5jhl88dk268s5vk7zxamsycfqj802xeya5yqh02crg8exj2k2y4639nfjg5xgvuh6ynyyfqnjdcuj"
    ),
    Value.from_primitive([3100000]),
  )
)

builder.add_output(
  TransactionOutput(
    Address.from_primitive(
      "addr_test1qp07eka2e6s6kv973t4js5jhl88dk268s5vk7zxamsycfqj802xeya5yqh02crg8exj2k2y4639nfjg5xgvuh6ynyyfqnjdcuj"
    ),
    Value.from_primitive([3200000]),
  )
)

signed_tx = builder.build_and_sign([psk], change_address=address)
context.submit_tx(signed_tx.to_cbor())

print("Broadcast tx:", signed_tx.id)
