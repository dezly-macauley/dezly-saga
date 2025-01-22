// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

contract FundMe {

    function fund() public payable {

        // 1e18 WEI = 1 ETH = 1 * 10 ** 18 WEI
        // ETH means Ether
        // If someone tries to deposit less than 1 ETH,
        // the transaction will be reverted.
        require(msg.value > 1e18, "The deposit amount was too low"); 

        // If the deposit amount does not meet the requirement,
        // The transaction will revert.
        // When a transaction is reverted,
        // any lines of code before the revert statement within the function,
        // will be reversed.
        // And the user who tried to send the failed transaction will incur
        // the gas fees.

        // If the gas fees are higher than the person's wallet balance,
        // then their wallet balance will simply be reduced to zero.
    
        // In most blochain systems a wallet balance can't be negative.

    }

    function withdraw() public {

    }

}
