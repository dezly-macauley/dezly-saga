// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import {PriceConverter} from "./PriceConverter.sol";

// NOTE: How to create custom errors
// Create the error outside the contract declaration

error NotOwner();

contract FundMe {

    // NOTE: You can make your contract use less gas by using the keywords
    // "constant" and "immutable"

    // The reason this uses less gas is because instead of being stored in 
    // inside of a storage slot, they will stored directly inside the byte
    // code of the contract

    // NOTE: Using the keyword "immutable"
    // This works the same way as the keyword "constant"
    // The difference is that you would use the keyword "immutable" when 
    // the variable is being declared on one line, and initialized in
    // a different line.

    address public immutable i_owner;

    constructor() {
                
        i_owner = msg.sender;
    }


    using PriceConverter for uint256;
  
    // NOTE: Using the keyword "constant"
    // When a variable is declared, and initialized
    // in the same line an you want to make it a constant 

    // you can use thi s if your have a variable that you know 
    // will never change after the the contract has been deployed.
    // Using the const keyword will actually make the contract use less gas.

    uint256 public constant MINIMUM_USD = 5e18;

    address[] public funders;

    mapping(address funder => uint256 amountFunded) public addressToAmountFunded;

    function fund() public payable {

        require(msg.value.getConversationRate() >= MINIMUM_USD, "The deposit amount was too low"); 

        funders.push(msg.sender);

        addressToAmountFunded[msg.sender] = addressToAmountFunded[msg.sender] + msg.value;

    }

    function withdraw() public onlyOwner {
        for (uint256 funderIndex = 0; funderIndex < funders.length; funderIndex += 1) {
           
            address funder = funders[funderIndex];
            addressToAmountFunded[funder] = 0;

        }

        (bool callSuccess, /*bytes memory dataReturned*/) = payable(msg.sender).
            call{value: address(this).balance}("");
        require(callSuccess, "Call failed");

    }

    modifier onlyOwner() {

        // NOTE: No custom error is used in this version

       // require(msg.sender == i_owner, "You are not the authorized to withdraw!");

       if(msg.sender != i_owner) { revert NotOwner(); }

       _;

    }

    // NOTE: How to handle cases where someone sends money to the contract 
    // without calling the fund function

    // These functions will trigger the fund function when that happens

    receive() external payable {
        fund();
    }

    fallback() external payable {
        fund();
    }

}
