// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract Incrementer {

    uint256 public number;
    event Increment(address who, uint256 value, uint256 updated_number);
    event Reset(address who);

    constructor(uint256 _initialNumber) {
        number = _initialNumber;
    }

    function increment(uint256 _value) public {
        number = number + _value;
        emit Increment(msg.sender, _value, number);
    }

    function reset() public {
        number = 0;
        emit Reset(msg.sender);
    }

}