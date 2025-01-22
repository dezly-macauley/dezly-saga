// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

// NOTE: How to import a smart contract from another file

import {SimpleStorage} from "./SimpleStorage.sol";

contract StorageFactory {

    // NOTE: Step 1: Create a placeholder for the contract when it is stored

    // This is a variable to store the address of the contract 
    // when it is created.
    // This line of code does NOT create the contract, it simply creates
    // a placeholder for it.

    SimpleStorage public simpleStorageContract;

    function createSimpleStorageContract() public {

        // NOTE: Step2: Create the new simpleStorage contract

        // The new contract created will follow the structure and functionality
        // of the "SimpleStorage" contract that was inherited from the
        // SimpleStorage.sol file

        simpleStorageContract = new SimpleStorage();


    }
}

