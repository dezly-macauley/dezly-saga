// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

// NOTE: This is how to import a library that you created

import {PriceConverter} from "./PriceConverter.sol";

// import {AggregatorV3Interface} from "@chainlink/contracts/src/v0.8/shared/interfaces/AggregatorV3Interface.sol";

contract FundMe {

    // NOTE: How to attach the functions of a library to all variable types
    // E.g. In this example the PriceConverter library and all its functions
    // have been attached to all uint256 variables in this contract.

    using PriceConverter for uint256;
   
    uint256 public minimumUsd = 5e18;

    // An array to keep track of who is depositing money into the smart contract
    address[] public funders;

    // A mapping to make it convienient to keep track of how much money
    // each person has deposited by their name
    // You can give names to the mappings
    mapping(address funder => uint256 amountFunded) public addressToAmountFunded;

    function fund() public payable {

        // NOTE: How to use the imported library


        // NOTE: This is how it would be written without using the library
        // require(getConversationRate(msg.value) >= minimumUsd, "The deposit amount was too low"); 

        // NOTE: This is how it is written when you use the library

        // The first input variable in the library is going to be the type
        // that you are using with the library.
        
        require(msg.value.getConversationRate() >= minimumUsd, "The deposit amount was too low"); 

        // Update the array with each deposit
        funders.push(msg.sender);

        // Update the mapping with each deposit
        // msg.value = 

        addressToAmountFunded[msg.sender] = addressToAmountFunded[msg.sender] + msg.value;

        // msg.sender = The public address of the person who is making the call.
        // msg.value = Number of Wei sent with the message

    }

}
