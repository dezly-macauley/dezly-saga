// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import {PriceConverter} from "./PriceConverter.sol";

contract FundMe {
    
    address public owner;
    
    // NOTE: A contructor function is a special function that is immediately
    // called when you deploy your contract.
    
    constructor() {
        
        // NOTE: Step 1: Set the owner of the contract

        // The address of the the person who is sending the contract is the owner
        owner = msg.sender;
    }


    using PriceConverter for uint256;
   
    uint256 public minimumUsd = 5e18;

    address[] public funders;

    mapping(address funder => uint256 amountFunded) public addressToAmountFunded;

    function fund() public payable {

        require(msg.value.getConversationRate() >= minimumUsd, "The deposit amount was too low"); 

        funders.push(msg.sender);

        addressToAmountFunded[msg.sender] = addressToAmountFunded[msg.sender] + msg.value;

    }

    function withdraw() public {

       // NOTE: Step 2: Setting up the withdraw function so that it can only be called 
       // by the owner

       require(msg.sender == owner, "You are not the authorized to withdraw!");

        for (uint256 funderIndex = 0; funderIndex < funders.length; funderIndex += 1) {
           
            address funder = funders[funderIndex];
            addressToAmountFunded[funder] = 0;

        }

        (bool callSuccess, /*bytes memory dataReturned*/) = payable(msg.sender).
            call{value: address(this).balance}("");
        require(callSuccess, "Call failed");

    }

    //=========================================================================

}
