pragma solidity ^0.4.11;

contract SimpleStorage {
    uint data;

    function setData(uint x) public {

        data = x;
    }

    function getData() public constant returns (uint) {

        return data;
    }
}