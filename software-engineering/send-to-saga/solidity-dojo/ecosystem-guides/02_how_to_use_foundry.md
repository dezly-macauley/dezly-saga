# How to use Foundry

### How to compile your smart contract

There are two commmands depending on the sitution
Regardless of which method you use, you will have a folder in your project
called "out". This is where you can find your compiled code.

___

Command 1:

`forge compile`

Use this when you are developing your smart contract, to check if you code
will actually compile. This command is what you use when you are deploying 
your contract to development environmment for testing.

___

`forge build`

Use this when you are are done developing your code and you want to deploy it
to a production environmment. Forge build does everything that `forge compile`
does, plus it adds some extra build steps that will optimize your code for
release.

___

### How to deploy your contract

There are two steps to this

#### Step 1: Decide what blockchain you will be deploying your contract to 



#### Step 2: Setup an account that will be used to deploy the contract 

#### Step 3: Now that you have all the information, use `forge create`

Tip: `forge create` commands can get pretty long, 
so I use the `\` to break them up into shorter lines
so that they are easier to read. 

When you run this command the terminal will process it as one long command.

#

This is the format of a `forge create` command:

forge create `The name of the contract in your .sol file` \
--rpc-url `the endpoint of the blockchain that you are trying to connect to` \
(This command will be finished later once I go over private key saftey)

#

___

