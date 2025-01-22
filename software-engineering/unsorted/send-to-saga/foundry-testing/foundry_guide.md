How to create a new project:

Create a folder, then cd into that folder

```
forge init
```

---

How to run a test:

```
forge test
```

Note: If you have a console.log in your test file then to see these logs
what you need to do is run the command like this:

```
forge test -vvv
```

By default "forge test" will run every single test function in your test
script. If you want to only run a specific test, you can do the following. 

forge test --mt testPriceFeedVersionIsAccurate -vv

--mt stands for match test
-vvv increases the verbosity (the amount of info to be displayed)

You would use -vvv to view console.log ouputs

---

## Type of Tests

1. Unit Test - Testing a specific part of your code
2. Integration Test - Testing how your code works with other parts of code
3. Forked Test - Testing how your code works on a simulated environment
4. Stagging - Testing how your code works on a real environment that is not
production.

---

## How to check how many lines of your code are actually tested.

forge coverage --fork-url $SEPOLIA_RPC_URL

You will get an output that looks like this:

| File                      | % Lines       | % Statements  | % Branches    | % Funcs     
 |
|---------------------------|---------------|---------------|---------------|-------------
-|
| script/DeployFundMe.s.sol | 0.00% (0/3)   | 0.00% (0/3)   | 100.00% (0/0) | 0.00% (0/1) 
 |
| src/FundMe.sol            | 16.67% (2/12) | 23.53% (4/17) | 0.00% (0/4)   | 25.00% (1/4)
 |
| src/PriceConverter.sol    | 0.00% (0/6)   | 0.00% (0/11)  | 100.00% (0/0) | 0.00% (0/2) 
 |
| Total                     | 9.52% (2/21)  | 12.90% (4/31) | 0.00% (0/4)   | 14.29% (1/7)

---

How to install a dependency like the one below:

https://github.com/smartcontractkit/chainlink-brownie-contracts

```
forge install smartcontractkit/chainlink-brownie-contracts@0.6.1 --no-commit
```
@0.6.1 is the version number

This will install:


    Installed chainlink-brownie-contracts 
0.6.1

You should see the folder in the "lib" folder of your Foundry project.

Next, add this line to your "foundry.toml" file:

remappings = ["@chainlink/contracts/=lib/chainlink-brownie-contracts/contracts"]

Then run this command from the terminal:

```
forge build
```

---

## How to write a test

Go to the "test" folder of your Foundry project

create your test file

This is the correct way to name it:

NameOfTheContractYouWantToTest-Test.t.sol

E.g. FundMeTest.t.sol

---
