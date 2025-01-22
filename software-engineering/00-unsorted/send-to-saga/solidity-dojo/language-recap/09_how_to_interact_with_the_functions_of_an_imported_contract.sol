// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import {SimpleStorage} from "./SimpleStorage.sol";

contract StorageFactory {

    SimpleStorage[] public listOfSimpleStorageContracts;

    function createSimpleStorageContract() public {
       
        SimpleStorage newSimpleStorageContract = new SimpleStorage();

        listOfSimpleStorageContracts.push(newSimpleStorageContract); 

    }

    // "Storage Factory Store"
    // This will allow this "Storage Factory" contract to use the "store"
    // function from the {SimpleStorage} contract.
    function sfStore(uint256 _simpleStorageIndex,
                     uint256 _favoriteNumber) public {
        
        // NOTE: In order to interact with a contract you will need two things.

        // The address of the contract that you want to interact with

        // The ABI of the contract = Application Binary Interface
        // The ABI is what tells your code how to interact with another
        // contract

        // To get the API of a contract when using the Remix IDE,
        // go to the "Compile" tab. 

        // Make sure to compile your contract first.

        // Then look for the little button at the bottom that says
        // "copy API to clipboard"
        // You can also find the ABI in the menu that
        // says "Compilation Details"

        // The ABI basically gives you a list of all the functions that 
        // your smart contract contains

        // The ABI of the contract is actually what Remix Ethereum IDE,
        // uses to create the interactive buttons that you press.

        // When you are trying to use the functionality of another contract,
        // You would use a line like this, so that the solidity compiler,
        // knows what the ABI is: 
        // import {SimpleStorage} from "./SimpleStorage.sol";

        // NOTE: You don't really need the ABI, you just need the function
        // selector but this will only be taught way later in the Cyfrin
        // Updraft course

        //=====================================================================
        
        // NOTE: How to interact with another smart contract.

        // Because this contract actually creates multiple
        // SimpleStorage contracts and adds them to a list,
        // the first step is actually selecting which one the 
        // contract you want to interact with.
    
        // So when the "sfStore" function is called, you will give it the
        // index of the contract that you want to interact with.

        SimpleStorage mySimpleStorage =
            listOfSimpleStorageContracts[_simpleStorageIndex];

        // NOTE: Using the function of the other contract.

        // Now you can interact with the store function 
        // of that SimpleStorage contract, from the "StorageFactory" contract
        // Notice how the original function from the SimpleStorage contract, 
        // is now being avaiblabe as a method for the varible, 
        // that was created using the SimpleStorage
        // template

        mySimpleStorage.store(_favoriteNumber);

    }

    // This function will allow you to use the "retrive" function that 
    // SimpleStorage contracts come with.
    // Notice that I try to keep this function as close to the original
    // "retrive" function as possible. 
    // The only difference is that because the StorageFactory contract,
    // creates multiple "SimpleStorage" contracts, I first need to target
    // the one I want to interact with.

    function sfRetrieve(uint256 _simpleStorageIndex)
    public view returns(uint256) {
      
        // Select the contract
        SimpleStorage mySimpleStorage =
            listOfSimpleStorageContracts[_simpleStorageIndex];

        // Notice how the "retrieve" function from the original contract,
        // is now avaiblabe as a method for the SimpleStorage function.
        return mySimpleStorage.retrieve(); 

    }

}
