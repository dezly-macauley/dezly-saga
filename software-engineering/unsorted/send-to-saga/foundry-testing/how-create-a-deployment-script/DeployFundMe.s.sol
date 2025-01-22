// SPDX-License-Identifier: MIT

// NOTE: Place this file in the "script" folder of your project

pragma solidity ^0.8.18;

// This will allow you to use the "vm" keyword from the Forge Standard Library
// to deploy your smart contract
import {Script} from "forge-std/Script.sol";

// Remember to import the smart contract that you want to test
import {FundMe} from "../src/FundMe.sol";

// NOTE: How to run this contract
// forge script script/DeployFundMe.s.sol

// Add "is Script" so that this deployment contract gains the functionality
// to use the "vm" keyword
contract DeployFundMe is Script {

    // Next thing you need is a "run" function to deploy your script
    function run() external {

        // This will deploy a new FundMe contract
        vm.startBroadcast();
        new FundMe();
        vm.stopBroadcast();
    }

}
