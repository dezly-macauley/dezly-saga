// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import {AggregatorV3Interface} from "@chainlink/contracts/src/v0.8/shared/interfaces/AggregatorV3Interface.sol";

contract FundMe {
   
    uint256 public minimumUsd = 5e18;

    // An array to keep track of who is depositing money into the smart contract
    address[] public funders;

    // A mapping to make it convienient to keep track of how much money
    // each person has deposited by their name
    // You can give names to the mappings
    mapping(address funder => uint256 amountFunded) public addressToAmountFunded;

    function fund() public payable {

        require(getConversationRate(msg.value) >= minimumUsd, "The deposit amount was too low"); 

        // Update the array with each deposit
        funders.push(msg.sender);

        // Update the mapping with each deposit
        // msg.value = 

        addressToAmountFunded[msg.sender] = addressToAmountFunded[msg.sender] + msg.value;

        // msg.sender = The public address of the person who is making the call.
        // msg.value = Number of Wei sent with the message


    }

    function getPrice() public view returns(uint256) {
    
        AggregatorV3Interface dataFeed = AggregatorV3Interface(0x694AA1769357215DE4FAC081bf1f309aDC325306);

        (
            /* uint80 roundID */,
            int answer,
            /*uint startedAt*/,
            /*uint timeStamp*/,
            /*uint80 answeredInRound*/
        ) = dataFeed.latestRoundData(); 
       
        // NOTE: Step 3: Getting the variable types and the decimals to align

        return uint256(answer * 1e10);

    }

    function getConversationRate(uint256 ethAmount) public view returns(uint256) {

        uint256 ethPrice = getPrice();
        uint256 ethAmountInUSD = (ethPrice * ethAmount) / 1e18;

        return ethAmountInUSD;
        
    }

    function getVersion() public view returns(uint256) {
        return AggregatorV3Interface(0x694AA1769357215DE4FAC081bf1f309aDC325306).version();
    }

}
