// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

contract SimpleStorage {
    
    uint256 myFavoriteNumber = 0;

    struct Person {
        uint256 favoriteNumber;
        string name;
    }

    Person[] public listOfPeople;

    mapping(string => uint256) public nameToFavoriteNumber;

    function store(uint256 _favoriteNumber) public {
        myFavoriteNumber = _favoriteNumber;
    }

    function retrieve() public view returns(uint256) {
        return myFavoriteNumber;
    }
    function addPerson(string memory _name, uint256 _favoriteNumber) public {
        listOfPeople.push(Person(_favoriteNumber, _name));

        nameToFavoriteNumber[_name] = _favoriteNumber;

    }

}

contract StorageFactory {

    // NOTE: How to create a function that creates a new smart contract
    // Step 1 - Create a variable that will store the new contract when it is
    // created.

    // "simpleStorage is the variable name"
    // SimpleStorage is the variable type

    // This is essentially a struct. 
    // A template to create a custom variable type.
    // In this case the template is the SimpleStorage contract above.

    // The variable type is "address" becasu SimpleStorage is a Smart Contract
    SimpleStorage public simpleStorage;

    // NOTE: Step 2 - create a function that will create the contract and 
    // and store it in the variable above.

    function CreateSimpleStorageContract() public {

        simpleStorage = new SimpleStorage();

    }
}

