// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import {SimpleStorage} from "./SimpleStorage.sol";

contract StorageFactory {

    // Create an array to store the addresses of each simple storage
    // contract that will be created
    SimpleStorage[] public listOfSimpleStorageContracts;

    function createSimpleStorageContract() public {
       
        // NOTE: Part 1

        // Create a placeholder a placeholder to store the single contract
        // and then actually create the contract.
        // So this is doing two things at once, declaring and intializing.
        SimpleStorage newSimpleStorageContract = new SimpleStorage();

        // NOTE: Part 2

        // After the contract has been created its address needs to be 
        // added to the array "listOfSimpleStorageContracts" so that you
        // can keep track of what contracts have been created.

        listOfSimpleStorageContracts.push(newSimpleStorageContract); 

    }

}
