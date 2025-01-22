# How to call a function on a smart contract from the command line:

_______________________________________________________________________________

## How to call a function that modifies the state of the smart contract

E.g. I want to call the `store` function from the `SimpleStorage` contract,
and I want to store the number `42`

```solidity

    function store(uint256 _favoriteNumber) public {
        myFavoriteNumber = _favoriteNumber;
    }

```

First make sure `Anvil` is running, then deploy your contract,
then make sure to copy the contract address (i.e. the public address where
the contract was deployed to)

Contract Address of my SimpleStorage Contract:

0x057ef64E23666F000b34aE31332854aCBd1c8544


### This is the command to use (If you are trying to interact with a local
blockchain like Anvil):

```
cast send 0x057ef64E23666F000b34aE31332854aCBd1c8544 \
"store(uint256)" 42 \
--rpc-url http://127.0.0.1:8545 \
--account anvil_test_account_3

```

_______________________________________________________________________________

## How to call a function that reads the the state of the smart contract

For example I want to call the `retrieve` function from `SimpleStorage`
contract that I deployed to the address:

0x057ef64E23666F000b34aE31332854aCBd1c8544

```solidity
    function retrieve() public view returns (uint256) {
        return myFavoriteNumber;
    }
```

I want to see what the value of myFavoriteNumber is, in order to confirm that
the amount of `42` I sent was properly stored

### This is the command to use:

```
cast call 0x057ef64E23666F000b34aE31332854aCBd1c8544 \
"retrieve()" 
```

You will get back the hex value:
0x000000000000000000000000000000000000000000000000000000000000002a

To get convert the hex value to a number that is human-readable:

```
cast --to-base \
0x000000000000000000000000000000000000000000000000000000000000002a \
dec
```
And you get 42

### For convience you can just make this one command:

```
cast --to-base \
$(cast call 0x057ef64E23666F000b34aE31332854aCBd1c8544 "retrieve()") \
dec
```


_______________________________________________________________________________


