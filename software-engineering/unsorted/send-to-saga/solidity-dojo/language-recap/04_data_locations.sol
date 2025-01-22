// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

contract SimpleStorage {
    
    // myFavoriteNumber is implicitly stored in "storage"
    // This is because it is not inside any function within the contract.
    uint256 myFavoriteNumber = 0;

    // NOTE: There are 6 places you can store data in Solidity
    
    // Storage, Memory, Calldata
    // Stack
    // Code, Logs

    // Memory = temporary variables that can be modified
    // Calldata = temporary variables that can't be modified
    // Storage = permanent variables that can be modified

    // NOTE: Data locations can only be specified for array, struct,
    // and mapping types. These variables are considered special types in
    // Solidity.
    // Strings are actually an array of bytes... see the note below

    struct Person {
        uint256 favoriteNumber;
        string name;
    }

    Person[] public listOfPeople;

    function store(uint256 _favoriteNumber) public {
        myFavoriteNumber = _favoriteNumber;
    }

    function retrieve() public view returns(uint256) {
        return myFavoriteNumber;
    }

    // NOTE: Strings are actually and array of bytes so you have to specify
    // a data location.
    function addPerson(string memory _name, uint256 _favoriteNumber) public {
        listOfPeople.push(Person(_favoriteNumber, _name));
    }

}
