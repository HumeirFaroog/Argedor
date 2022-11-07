// SPDX-License-Identifier: MIT

pragma solidity >=0.4.22 <0.9.0;


contract token {

    uint balance;

    function updateBalance(uint _balance) external {
        balance = _balance;
    }

    function readBalance() external view returns(uint) {
        return balance;
    }

}