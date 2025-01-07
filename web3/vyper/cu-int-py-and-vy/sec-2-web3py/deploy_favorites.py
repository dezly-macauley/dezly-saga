# This will allow this program to use the `compile_code` function from the
# vyper module to compile `favorites.vy` into bytecode.
from vyper import compile_code
from web3 import Web3

# This will give this program a way to create a smart contract from bytecode

def main():

    #__________________________________________________________________________
    # SECTION: Read the Vyper file and use the `compile_code` funtion from 
    # the `vyper` compile it into Python dictionary

    print("Let's read in the Vyper code and deploy it")

    # `r` means open the file in read only mode
    with open("favorites.vy", "r") as favorites_file:
        favorites_code = favorites_file.read()

        # compile_code takes a file, and creates a dictionary from that file.
        # output_formats is where you specify what key-value pairs 
        # that you want to be returned.
        # In summary I want to create a dictionary called `compliation_details`
        # and I want the dictionary to have a key called "bytecode" and the 
        # value will be the bytecode of the compiled smart contract

        # It will also have a key-value pair called the "abi"
        # The abi (Application Binary Interface) in simple terms is a list
        # of functions that people can use to interact with a smart contract

        compliation_details = compile_code(
                favorites_code, 
                output_formats=["bytecode", "abi"]
        )

        # NOTE: compliation_details is a Python dictionary
        # The output should look something like this:

        # print(compliation_details)

        # {'bytecode': '0x3461001a5760075f5561021661001e61000039610216610000f35b5f80fd5f3560e01c60026005820660011b610
        # 20c01601e395f51565b636057361d811861003457602436103417610208576004355f55005b632578e0f08118610204576024361034
        # 17610208576004356005811015610208576001015460405260206040f35b632e64cec1811861007d5734610208575f5460405260206
        # 040f35b635ec0e9c9811861020457602436103417610208576020806040526006600435600581101561020857026006018160400160
        # 40825482528060208301526001830181830160208254015f81601f0160051c600581116102085780156100f457905b8085015481600
        # 51b8501526001018181186100de575b5050508051806020830101601f825f03163682375050601f19601f8251602001011690509050
        # 8101905090509050810190506040f35b6305a8568681186102045760443610341761020857600435600401803560648111610208575
        # 060208135018082604037505060243560e0526020604051018060406101005e5060066024546005811015610208570260060160e051
        # 815560206101005101600182015f82601f0160051c600581116102085780156101c257905b8060051b6101000151818401556001018
        # 181186101aa575b50505050506024356024546005811015610208576001015560245460018101818110610208579050602455602435
        # 60256040516060206020525f5260405f2055005b5f5ffd5b5f80fd006200180204012a020484190216810a00a165767970657283000
        # 4000014', 'abi': [{'stateMutability': 'nonpayable', 'type': 'function', 'name': 'store', 'inputs': [{'name'
        # : 'favorite_number', 'type': 'uint256'}], 'outputs': []}, {'stateMutability': 'view', 'type': 'function', '
        # name': 'retrieve', 'inputs': [], 'outputs': [{'name': '', 'type': 'uint256'}]}, {'stateMutability': 'nonpay
        # able', 'type': 'function', 'name': 'add_person', 'inputs': [{'name': 'name', 'type': 'string'}, {'name': 'f
        # avorite_number', 'type': 'uint256'}], 'outputs': []}, {'stateMutability': 'view', 'type': 'function', 'name
        # ': 'list_of_numbers', 'inputs': [{'name': 'arg0', 'type': 'uint256'}], 'outputs': [{'name': '', 'type': 'ui
        # nt256'}]}, {'stateMutability': 'view', 'type': 'function', 'name': 'list_of_people', 'inputs': [{'name': 'a
        # rg0', 'type': 'uint256'}], 'outputs': [{'name': '', 'type': 'tuple', 'components': [{'name': 'favorite_numb
        # er', 'type': 'uint256'}, {'name': 'name', 'type': 'string'}]}]}, {'stateMutability': 'nonpayable', 'type': 
        # 'constructor', 'inputs': [], 'outputs': []}]}

    #__________________________________________________________________________
    # SECTION: Turn the compliation_details output into a contract object

    # NOTE: You need to have anvil (part of the Foundry Framework) 
    # for this part to work

    # Run `anvil` in a separate terminal
    # Look for the line `Listening on 127.0.0.1:8545`
    # This is the RPC url (It's basically the address of the chain you want
    # to deploy the smart contract to)

    # First create a web3 object to use its functionality
    # To initialize a web3 object. You have to connect to a blockchain node 
    w3 = Web3(Web3.HTTPProvider("http://127.0.0.1:8545"))
    
    # This will return a contract object
    # The bytecode of this contract will be set to the bytecode field of 
    # the compliation_details dictionary
    favorites_contract = w3.eth.contract(
            bytecode=compliation_details["bytecode"],
            abi=compliation_details["abi"]
    )
    # print(favorites_contract)

#______________________________________________________________________________
# SECTION: Build a transaction that deploys the smart contract

    # This is essentially what the `Remix Ethereum IDE` does when you click
    # the deploy button.

    print("Building the transaction...")

    # This is will use the `web3` module to automatically 
    # populate a lot of the fields that a transaction would have.
    
    # NOTE: The `bytecode` and `abi` field were already filled in when the
    # `favorites_contract` variable was created
    transaction = favorites_contract.constructor().build_transaction()
    print(transaction)

    # {'value': 0, 'gas': 191170, 'maxFeePerGas': 3000000000, 'maxPriorityFeePerGas': 1000000000, 'chainId': 3133
    # 7, 'data': '0x3461001a5760075f5561021661001e61000039610216610000f35b5f80fd5f3560e01c60026005820660011b61020
    # c01601e395f51565b636057361d811861003457602436103417610208576004355f55005b632578e0f0811861020457602436103417
    # 610208576004356005811015610208576001015460405260206040f35b632e64cec1811861007d5734610208575f546040526020604
    # 0f35b635ec0e9c981186102045760243610341761020857602080604052600660043560058110156102085702600601816040016040
    # 825482528060208301526001830181830160208254015f81601f0160051c600581116102085780156100f457905b808501548160051
    # b8501526001018181186100de575b5050508051806020830101601f825f03163682375050601f19601f825160200101169050905081
    # 01905090509050810190506040f35b6305a856868118610204576044361034176102085760043560040180356064811161020857506
    # 0208135018082604037505060243560e0526020604051018060406101005e5060066024546005811015610208570260060160e05181
    # 5560206101005101600182015f82601f0160051c600581116102085780156101c257905b8060051b610100015181840155600101818
    # 1186101aa575b5050505050602435602454600581101561020857600101556024546001810181811061020857905060245560243560
    # 256040516060206020525f5260405f2055005b5f5ffd5b5f80fd006200180204012a020484190216810a00a16576797065728300040
    # 00014', 'to': b''}

#______________________________________________________________________________

if __name__ == "__main__":
    main()
