// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import {SimpleStorage} from "./SimpleStorage.sol";

// NOTE: Using the is keyword will give the AddFiveStorage contract,
// all of the functionality of the SimpleStorage contract

contract AddFiveStorage is SimpleStorage {
    
    // NOTE: How to modify a function from another contract without 
    // changing the original
    
    /*  
   
        E.g. Here is a function from the SimpleStorage contract called
        "store"

        function store(uint256 _favoriteNumber) public {
            myFavoriteNumber = _favoriteNumber;
        }
        
        I want to use this in function in the "AddFiveStorage" contract,
        but I want a slight modification.
        When this function is called from this "AddFiveStorage" contract,
        I want it to do this instead:

        function store(uint256 _favoriteNumber) public {
            myFavoriteNumber = _favoriteNumber + 5;
        }

    */

        // NOTE: Step 1: Add the keyword "override" to let the compiler 
        // know that you want to use this version of the "store" function
        // and not the original

        function store(uint256 _favoriteNumber) public override {
            myFavoriteNumber = _favoriteNumber + 5;
        }

        // NOTE: Step 2: Now go to the "store" function in the original
        // imported contract and add the keyword "virtual" to the function

        // This will let Solidity know that the function has the ability
        // to be overidden when it is used in another contract.

        // function store(uint256 _favoriteNumber) public virtual {
        //     myFavoriteNumber = _favoriteNumber + 5;
        // }

}
