// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {Test, console} from "forge-std/Test.sol";
import {Counter} from "../src/Counter.sol";

contract CounterTest is Test {
    Counter public counter;

    function setUp() public {
        counter = new Counter();
        counter.setNumber(0);
    }

    function test_Increment() public {
        
        // This calls the "increment" function from the counter contract
        // The "increment" function is supposed to increase the value of 
        // the the variable called "number" by 1
        counter.increment();

        // This line of code will then check that the value of the variable
        // "number" has been increased by 1

        // This is the syntax:
        // assertEq(the-value that you want to check,
        // what that value should currently be);
        assertEq(counter.number(), 1);
    }

    function testFuzz_SetNumber(uint256 x) public {
        counter.setNumber(x);
        assertEq(counter.number(), x);
    }
}
