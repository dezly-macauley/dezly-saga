// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

// NOTE: Need more info on this

// This line requires npm to work
// https://www.npmjs.com/package/@chainlink/contracts
// This is required to interact with the Price Feed contract on chainlink
import {AggregatorV3Interface} from "@chainlink/contracts/src/v0.8/shared/interfaces/AggregatorV3Interface.sol";

contract FundMe {
   
    uint256 public minimumUsd = 5e18;
    // same as 
    // uint256 public minimumUsd = 5 * 1e18;

    function fund() public payable {

        require(getConversationRate(msg.value) >= minimumUsd, "The deposit amount was too low"); 

    }

    // function withdraw() public {}

    // This function will get the latest price of Ethereum in terms of USD
    function getPrice() public view returns(uint256) {
    
        // NOTE: How to get the latest price of Ethereum using Chainlink
       
        // NOTE: Step 1: First you need the to address of the price feed
        // contract for the specific data that you are trying to get.

        // https://docs.chain.link/data-feeds/price-feeds/addresses

        // If you are looking for the Priced Feeds for the Sepolia testnet,
        // then use this link:
        // NOTE:  https://docs.chain.link/data-feeds/price-feeds/addresses?network=ethereum&page=1#sepolia-testnet

        // Then copy the contract address in the Sepolia section that says "ETH / USD"

        // 0x694AA1769357215DE4FAC081bf1f309aDC325306

        // NOTE: Step 2: Now you need a way to interact with the functions 
        // of the smart contract.


        // NOTE: https://docs.chain.link/data-feeds/using-data-feeds#solidity

        // Add this line to your smart contract to import the Price Feed:
        // import {AggregatorV3Interface} from "@chainlink/contracts/src/v0.8/shared/interfaces/AggregatorV3Interface.sol";

        // NOTE: Next create a variable to hold the interface of dataFeed contract

        // The address in the brackets is actually the contract address of the Sepolia "ETH / USD"
        // price feed
        AggregatorV3Interface dataFeed = AggregatorV3Interface(0x694AA1769357215DE4FAC081bf1f309aDC325306);

        // NOTE: Next paste this code from the AggregatorV3Interface

        // The data feed is set up to return multiple values, however for this
        // contract, I only care abou the "answer" which in this case is the price of ETH / USD
        // However even when you comment out the other variables that you don't need,
        // it is important to leave in the commas at the end to show that you are skipping those other values.

        (
            /* uint80 roundID */,
            int answer,
            /*uint startedAt*/,
            /*uint timeStamp*/,
            /*uint80 answeredInRound*/
        ) = dataFeed.latestRoundData(); 
       
        // NOTE: Step 3: Getting the variable types and the decimals to align

        // VARIABLE TYPES are different
        // msg.value is a uint256
        // int answer is an int
        // So "answer" has to be converted to an int 

        // The Number of decimals are different
        // msg.value has 8 decimal places
        
        // "int answer" will return the price of ETH in USD
        // This will return 18 decimal places


        // So to get the decimals to add up

        return uint256(answer * 1e10);
        // Now price has 10 extra decimal places

    }

    // The purpose of this function is to get msg.value in terms of dollars
    // using the get price function
    function getConversationRate(uint256 ethAmount) public view returns(uint256) {

        // So first you get the price of Ethereum in USD
        uint256 ethPrice = getPrice();
        uint256 ethAmountInUSD = (ethPrice * ethAmount) / 1e18;
        // You divide by 1e18 because both of these have 18 decimal places.
        // So if you don't divide by 1e18 you will get this massive number 
        // of decimals

        // NOTE: In Solidity you always want to multiply before you divide
        
        return ethAmountInUSD;
        
    }

}
