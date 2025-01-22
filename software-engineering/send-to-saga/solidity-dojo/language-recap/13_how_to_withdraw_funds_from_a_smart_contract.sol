// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

// NOTE: This is how to import a library that you created

import {PriceConverter} from "./PriceConverter.sol";

contract FundMe {

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

        for (uint256 funderIndex = 0; funderIndex < funders.length; funderIndex += 1) {
           
            // NOTE: Reset the amount that each person deposited to 0 
            // when the witdraw function is called

            address funder = funders[funderIndex];
            addressToAmountFunded[funder] = 0;

        }

        // NOTE: How to reset an array in Solidity
        funders = new address[](0);


        //=========================================================================
        
        // NOTE: There are three different ways to send native blockchain currency

        // Method 1: Transfer

        // address(this).balance means the account balance of this smart contract
        // msg.sender is an address variable type,
        // To withdraw you need to type cast it to a payable address type  
        payable(msg.sender).transfer(address(this).balance);
        // Transfer function is capped at 2300 gas. If more gas is used, it
        // will throw an error, and revert the transaction.

        //=========================================================================

        // Method 2: Send
        // Also capped at 2300 gas. If more gas is used it returns a bool
        bool sendSuccess = payable(msg.sender).send(address(this).balance);

        // Dispaly an error message if the send was not successful.
        require(sendSuccess, "Send Failed");
        // payable 

        //=========================================================================
       
        // NOTE: Method 3: Call (This is the recommended way to send transactions)

        // This doesn't have a gas cap.

        // Call
        // () = something
        // This is the syntax you use when a function returns more than one
        // variable

        // The call functions returns two things. The success of the transaction
        // and data returned by the call
        // The "memory" keyword is used because dataReturned is a bytes object
        // It's commented out in this example because I only care
        // if the transaction was successful or not.
        
        // NOTE: Leave in the comma after callSuccess or it won't work.

        (bool callSuccess, /*bytes memory dataReturned*/) = payable(msg.sender).
            call{value: address(this).balance}("");
        require(callSuccess, "Call failed");

        //=========================================================================

    }


}
