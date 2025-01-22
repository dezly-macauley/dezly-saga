// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;
// The line above tells the compiler what version of Solidity
// should be used to compile this Smart Contract

// E.g. ^0.8.18 means 0.18.18 and higher
// E.g. 0.8.18 means the contract can only be compiled with 0.18.18.
// E.g. >=0.8.18 <0.9.0 You can also use a range

// "SimpleStorage" is the name of the contract that I am creating.
// Note that it is possible to have multiple contracts in the same .sol file
contract SimpleStorage {
    
    // NOTE: Basic Solidity Types
    // bytes
    
    //=========================================================================
    
    // NOTE: Boolean
    bool hasFavoriteNumber = true;
    bool isOnline = false;

    //=========================================================================

    // NOTE: Integers
    
    // An unsigned integers of 256 bits.
    // Unsigned means that the number stored in the variable must be positive.
    
    // 256 is the default size in Solidity.
    // So uint favoriteNumber = 88; is the same as the declaration below:
    uint256 favoriteNumber = 88;

    // This is a signed integer so it can be positive
    int256 bankBalance = -50;

    //=========================================================================
    
    // NOTE: Address

    // Solidity also has a special variable type called "address",
    // for storing wallet addresses.

    address myMetaMaskTestWallet = 0xE783787AeD79aD4915a2c00c1f6220C477160A41;

    //=========================================================================

    // NOTE: Strings
    
    // Strings in Solidity are actually just bytes objects but specifically
    // for text

    string myMessage = "GM Friends!";

    //=========================================================================
    
    // NOTE: Bytes

    // Need more info on this
    // The largest variable type of this you can get is bytes32
    bytes32 favoriteBytes = "cat"; 

    //=========================================================================

}
