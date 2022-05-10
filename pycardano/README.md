## Usage

Rough steps to start sending some ADA txs using these scripts (note by default they are for test net)

1. `pip install pycardano`

2. Create a Cardano wallet with the `create_wallet.py` script.
  * This will create payment and staking private keys in `payment.skey` and `stake.skey` files (it will fail if they already exist) and output the address
  * Fund this address from the cardano faucet https://testnets.cardano.org/en/testnets/cardano/tools/faucet/

3. You need an api key for https://blockfrost.io/ which provides a free, REST api which wraps a cardano node. To create + use the api key:
  * Create account with them and create a project in their dasahboard
  * Paste the key into `blockfrost.key`
  * Test that this works using the `list_utxos.py` script


4. `python3 send_tx.py`

5. profit
