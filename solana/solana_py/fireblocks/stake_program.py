"""Library to interface with the stake program."""
from __future__ import annotations
from typing import NamedTuple, Union

from solana._layouts.system_instructions import SYSTEM_INSTRUCTIONS_LAYOUT, InstructionType
from solana.publickey import PublicKey
from solana.transaction import Transaction, TransactionInstruction
from solana.system_program import *


STAKE_PROGRAM_ID: PublicKey = PublicKey("StakeConfig11111111111111111111111111111111")
"""Public key that identifies the Stake program."""


# Instruction Params
class CreateStakeAccountParams(NamedTuple):
    """Create stake account transaction params."""

    from_pubkey: PublicKey
    """"""
    stake_pubkey: PublicKey
    """"""
    lamports: int
    """"""

def create_stake_account(params: CreateStakeAccountParams) -> TransactionInstruction:
    return  create_account(
                CreateAccountParams(
                    from_pubkey=params.from_pubkey, new_account=params.stake_pubkey,
                    lamports=params.lamport, space=200, program_id=STAKE_PROGRAM_ID
                )
            )
            )
