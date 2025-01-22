// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

contract SimpleStorage {
    
    uint256 myFavoriteNumber = 0;

    struct Person {
        uint256 favoriteNumber;
        string name;
    }

    Person[] public listOfPeople;

    // NOTE: How to create a "Mapping"

    // Mappings are similar to dictionaries in Python
    // The way this works is that you enter a key and it returns a 
    // value that is associated with that key.

    // If you enter a key that doesn't exist,
    // then the mapping will return the default value on the right
    // of the arrow. In this example, the default for uint256 = 0
    mapping(string => uint256) public nameToFavoriteNumber;

    // NOTE: That you still have to implement this.
    // Usually you would have a function that updates the mapping

    function store(uint256 _favoriteNumber) public {
        myFavoriteNumber = _favoriteNumber;
    }

    function retrieve() public view returns(uint256) {
        return myFavoriteNumber;
    }
    
    // This function takes a name and a favorite number,
    // uses it to create a new "Person" variable Person(_favoriteNumber, _name)
    // And then it adds that new "Person" variable type to a dynamic array
    // call "listOfPeople"
    function addPerson(string memory _name, uint256 _favoriteNumber) public {
        listOfPeople.push(Person(_favoriteNumber, _name));

        // NOTE: The implementation of the mapping

        // This line of code is creating a new entry to the mapping
        // variable to keep track of what values were updating in the array.
        
        // E.g. If the function was called like this: addPerson("Jin", 7)
        // The new entry will be created in the mapping.
        // Each entry has a key-value pair:
        // The key is "Jin", and the value that goes with that key is 7
        nameToFavoriteNumber[_name] = _favoriteNumber;

    }

}
