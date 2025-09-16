// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

contract Add {

    event Added(uint indexed oldValue, uint indexed newValue);

    uint public a = 1;

    function add() public returns (uint) {
        emit Added(a, a+1);
        return a++;
    }
}