from pycardano import *
from pathlib import Path
from blockfrost import BlockFrostApi, ApiError, ApiUrls


network = Network.TESTNET

ssk = StakeSigningKey.load("stake.skey")
svk = StakeVerificationKey.from_signing_key(ssk)

address = Address(staking_part=svk.hash(), network=network)
bf_key = Path('blockfrost.key').read_text().rstrip()

api = BlockFrostApi(project_id=bf_key, base_url=ApiUrls.testnet.value)

account_rewards = api.account_rewards(
        stake_address=address,
        count=20,
        gather_pages=True, # will collect all pages
)
print(f"###### Staking rewards for {address}")
print(account_rewards)
