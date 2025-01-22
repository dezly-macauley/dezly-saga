// SPDX-License-Identifier: MIT

pragma solidity ^0.8.18;

import {AggregatorV3Interface} from "@chainlink/contracts/src/v0.8/shared/interfaces/AggregatorV3Interface.sol";

// NOTE: To create a library you would use this keyword instead

library PriceConverter {

    // NOTE: For each function make sure that you use the keyword internal

    function getPrice() internal view returns(uint256) {
    
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

    function getConversationRate(uint256 ethAmount) internal view returns(uint256) {

        uint256 ethPrice = getPrice();
        uint256 ethAmountInUSD = (ethPrice * ethAmount) / 1e18;

        return ethAmountInUSD;
        
    }

    function getVersion() internal view returns(uint256) {
        return AggregatorV3Interface(0x694AA1769357215DE4FAC081bf1f309aDC325306).version();
    }



}
