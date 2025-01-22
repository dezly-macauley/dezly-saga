// SPDX-License-Identifier: MIT

pragma solidity ^0.8.18;

contract FallbackExample {


    // NOTE: How to handle cases where someone sends money to the contract 
    // without calling the fund function

    uint256 public result;

    // NOTE: This is used when someone tries to interact with the contract,

    // E.g. Make a deposit without using any associated functions
    // in the contract
    // AND the transaction they send has no call data

    receive() external payable {
        result = 1;
    }

    // NOTE: Works the same way as receive() but this for when the transaction
    // sent has call data.

    fallback() external payable {
        result = 2;
    }

}
