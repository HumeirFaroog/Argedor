// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract MyCrypto is ERC20 {

    constructor() ERC20("My Custom Token", "MCT") {
        _mint(msg.sender, 100 * 10 ** ERC20.decimals());
    }
}