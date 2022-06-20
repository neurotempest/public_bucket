// SPDX-License-Identifier: MIT
pragma solidity >0.8.0;

contract Forwarder {
  address payable addr;

  constructor() {
    addr = payable(msg.sender);
  }

  receive() external payable {
    addr.transfer(msg.value);
  }
}
