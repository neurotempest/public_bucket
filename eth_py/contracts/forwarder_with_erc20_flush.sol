// SPDX-License-Identifier: MIT
pragma solidity >0.8.0;

abstract contract ERC20Interface {
  function transfer(address _to, uint256 _value)
    public
    virtual
    returns (bool success);

  function balanceOf(address _owner)
    public
    virtual
    view
    returns (uint256 balance);
}

contract Forwarder {
  address payable addr;

  constructor() {
    addr = payable(msg.sender);
  }

  receive() external payable {
    addr.transfer(msg.value);
  }

  function flushTokens(address tokenAddr) external payable {
    require(addr == msg.sender);

    ERC20Interface instance = ERC20Interface(tokenAddr);
    address forwarderAddr = address(this);
    uint256 forwarderBal = instance.balanceOf(forwarderAddr);
    if(forwarderBal == 0) {
      return;
    }

    bool success = instance.transfer(addr, forwarderBal);
    require(success);
  }
}
