// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

contract SimpleStorage {

    uint256 myFavoriteNumber = 0;

    //=========================================================================
    
    // NOTE: Arrays
   
    // listOfFavoriteNumbers is a variable type that is used to store arrays
    // uint256[] listOfFavoriteNumbers;

    // NOTE: How to create your own data types
    
    struct Person {
        uint256 favoriteNumber;
        string name;
    }

    // NOTE: How to create a variable from a struct

    // Person public myFriend = Person({favoriteNumber: 7, name: "Dezly"});

    // NOTE: How to create an array of structs (custom variable types)

    // This is a dynamic array. It's size can grow and shrink
    // Static arrays have a [] with no number inside
    Person[] public listOfPeople;

    // static array
    // The example below is a static array. It can only hold three structs
    // or three items of the same variable type.
    // Person[3] public listOfPeople;

    //=========================================================================
    function store(uint256 _favoriteNumber) public {
        myFavoriteNumber = _favoriteNumber;
    }

    function retrieve() public view returns(uint256) {
        return myFavoriteNumber;
    }


    //=========================================================================

    // NOTE: How to create a function that will add a struct to an array of 
    // structs
    
    function addPerson(string memory _name, uint256 _favoriteNumber) public {
        listOfPeople.push(Person(_favoriteNumber, _name));
    }


    //=========================================================================

}
