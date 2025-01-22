// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

contract SimpleStorage {

    // NOTE: Visibility "Public"

    // The default visibility of a variable in Solidity is "internal"
    // If you want people to be  able to view what the value,
    // of the variable is when the contract has been deployed, 
    // then you have to set the visibility to "public"
    // public also means that other contracts can make a call to view what
    // the value of this variable is
    uint256 public favoriteNumber = 0;

    function store(uint256 _favoriteNumber) public {
        favoriteNumber = _favoriteNumber;
    }

    // NOTE: Visibility "Private"

    // This means that only the current contract can acces this variable

    // NOTE: Visibility "External" (Only functions can use this keyword)

    // This means that the function is only visible externally and can only 
    // be called from outside the contract

    // NOTE: Visibility "Internal" 
    // (Only the current contract and its child contracts can call
    // the function)

    // Neeed to learn more about this

    // ========================================================================
    
    // NOTE: Modifiers

    //-------------------------------------------------------------------------
    
    // NOTE: "View"
    // This is for functions that do not modify the state of the contract
    
    // In simple terms, this is a keyword for functions that only read
    // information from the smart contract, and do not make any changes.

    function retrieve() public view returns(uint256) {
        return favoriteNumber;
    }

    //-------------------------------------------------------------------------
    
    // NOTE: "Pure"
    // This is for functions that not modify the state of the contract
    // AND do not read from the contract either.

    function returnTheNumber200() public pure returns(uint256) {
        return 200;
    }

    // ========================================================================
   
    // NOTE: "Payable"
    // Means that the this function can be called by someone
    // to deposit funds into the contract


    // ========================================================================

}
