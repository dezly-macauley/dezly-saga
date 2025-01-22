// SPDX-License-Identifier: MIT

pragma solidity 0.8.19;

// SECTION: How to use this file to deploy a contract

// NOTE: TLDR... Use this code

/*
    forge script script/deploy_simple_storage.s.sol:DeploySimpleStorage \
    --rpc-url http://127.0.0.1:8545 \
    --broadcast \
    --account anvil_test_account_3
*/

// Step 1: Make sure the local Anvil blockchain is running if you are deploying
// to Anvil
// Open a seperate terminal and run the command `anvil`

// Step 2: Make sure that you are in the root of the project.
// Open up the command line and then run this:

/*
    forge script script/deploy_simple_storage.s.sol:DeploySimpleStorage \
    --rpc-url http://127.0.0.1:8545 \
    --broadcast \
    --private-key 0x7c852118294e51e653712a81e05800f419141751be58f605c371e15141b007a6
*/

// The rpc-url used here is for the Anvil local blockchain

// The private key used is for the anvil test account (3)

// Public address:
// (3) 0x90F79bf6EB2c4f870365E785982E1f101E93b906

// Private key:
// (3) 0x7c852118294e51e653712a81e05800f419141751be58f605c371e15141b007a6

// NOTE: NEVER expose a real private key of an account with 
// real funds like this

//_____________________________________________________________________________

// NOTE: After you run this command make sure to copy the public address where
// the smart contract was deployed to

// It will look something like this
// 0: contract SimpleStorage 0x057ef64E23666F000b34aE31332854aCBd1c8544

// It you don't specify an RPC URL in Foundry, Foundry will temporarily start
// the `Anvil` blockchain, deploy your contract there and then shut down.

//_____________________________________________________________________________

// SECTION: How to to use the `cast wallet` command to hide your private keys

// For this example I will use the Anvil test account (3)

// Public address:
// (3) 0x90F79bf6EB2c4f870365E785982E1f101E93b906

// Private key:
// (3) 0x7c852118294e51e653712a81e05800f419141751be58f605c371e15141b007a6

/*

    Open the terminal and run:

    cast wallet import anvil_test_account_3 --interactive

    Then paste the private key

    You will be asked to enter a password (This is if you ever need to 
    retrieve the private key)
   
    Again since this is only a test account, so I will use a simple password like:

    To view the accounts that you have stored use this command:
    cast wallet list

    To delete private keys that you have saved, you can find them in:
    cd ~
    cd .config/foundry/keystores

*/

// NOTE: Use this code

/*
    forge script script/deploy_simple_storage.s.sol:DeploySimpleStorage \
    --rpc-url http://127.0.0.1:8545 \
    --broadcast \
    --account anvil_test_account_3 \ 
    --vvvv
*/

// --broadcast means that this will actually deploy to to blockchain or testnet
// without this flag Foundry will simply perform a simulation.

// --sender is the public address that matches the account
// vvvv This increases the verbosity. In simple terms, you will get more
// detailed information after you run the command

//_____________________________________________________________________________

// NOTE: Step 1: Import the contract that will be deployed by this script
import {SimpleStorage} from "../src/simple_storage.sol";

//_____________________________________________________________________________

// NOTE: Step 2: Import the "Script" contract
// from the "forge" standard library

import {Script} from "forge-std/Script.sol";

//_____________________________________________________________________________

// DeploySimpleStorage will now inherit all the functionality of the Script
// contract

contract DeploySimpleStorage is Script {
    // "run" is a function that will return a SimpleStorage contract
    function run() external returns (SimpleStorage) {
        // NOTE: This is optional, but you can also set some boilerplate code
        // before you start the transaction
        // uint256 startingValue = 1;

        // NOTE: vm is a special keyword that is part of the "forge" standard
        // library. It is not part of the Solidity langauge
        vm.startBroadcast();

        // Everything between startBroadcast and stopBroadcast is what
        // will be deployed

        // This will create a new SimpleStorage contract and store it in
        // the variable simpleStorage
        SimpleStorage simpleStorage = new SimpleStorage();

        vm.stopBroadcast();
        return simpleStorage;
    }
}
