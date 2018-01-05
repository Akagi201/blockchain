pragma solidity ^0.4.4;

contract HelloWorld {
  function HelloWorld() public {
    // constructor
  }

  function sayHello() public pure returns (string) {
    return ("Hello World");
  }
}
