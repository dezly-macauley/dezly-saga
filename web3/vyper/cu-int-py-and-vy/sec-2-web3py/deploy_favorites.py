# This will allow this program to use the `compile_code` function from the
# vyper module to compile `favorites.vy` into bytecode.
import getpass
from vyper import compile_code
from web3 import Web3
from dotenv import load_dotenv

# NOTE: Import the `KEYSTORE_PATH` variable from the file `encrypt_key.py` 
from encrypt_key import KEYSTORE_PATH

# This is a built-in module
import os

load_dotenv()
# This is the address of the blockchain that you will be 
# deploying the contract to
RPC_URL = os.getenv("ANVIL_RPC_URL")

# Anvil - Available account (0)
ADDRESS_OF_DEPLOYER = os.getenv("ADDRESS_OF_DEPLOYER")
# Open up anvil in the terminal and choose any of the available accounts
# MY_ADDRESS = Web3.to_checksum_address(ADDRESS_OF_DEPLOYER)


# WARNING: This is bad practice!!! 
# Never have you private key in plain text like this!!
# This is fine for now because this is a fake account

# Open up anvil in the terminal and choose any of the available accounts
# PRIVATE_KEY = "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80" 
PRIVATE_KEY = decrypt_key()


# NOTE: to_checksum_address():

# Think of this like a standardized way to write Ethereum addresses that helps catch typos and errors. Here's what it does:
# Takes a regular Ethereum address (like 0xf39fd6e51aad88f6f4ce6ab8827279cfffb92266)
# Applies a special algorithm that mixes uppercase and lowercase letters based on the address's hash
# Returns something like: 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266

# Notice how some letters are uppercase and some are lowercase? That's not random - it's a checksum! If you accidentally mistype even one character, the checksum format would be wrong, and Web3.py would catch the error. It's like built-in error detection.


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
    w3 = Web3(Web3.HTTPProvider(RPC_URL))
    
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

#______________________________________________________________________________

    # SUB_SECTION: Setting the `nonce` field of the smart contract

    # The value inside the bracket is the public address of the account
    # that you will use to deploy the smart contract
    # You can get this from Anvil

    # Open up anvil in the terminal an go to available accounts 0

    nonce = w3.eth.get_transaction_count(MY_ADDRESS)

    # This is will use the `web3` module to automatically 
    # populate a lot of the fields that a transaction would have.
    
    # The `bytecode` and `abi` field were already filled in when the
    # `favorites_contract` variable was created

    transaction = favorites_contract.constructor().build_transaction(

        # NOTE: Setting the nonce and other fields manually
        {
            "nonce": nonce,
            "from": MY_ADDRESS,
            "gasPrice": w3.eth.gas_price
        }

    )

    # NOTE: How to modify a specific field

    # transaction["nonce"] = nonce
#______________________________________________________________________________
    
    print(transaction)

    private_key = decrypt_key()
    
    # This will output


    # {'value': 0, 'gas': 191170, 'chainId': 31337, 'nonce': 0, 'from': '0xf39Fd6e51aad88F6F4ce6aB8827279cffFb922
    # 66', 'gasPrice': 2000000000, 'data': '0x3461001a5760075f5561021661001e61000039610216610000f35b5f80fd5f3560e
    # 01c60026005820660011b61020c01601e395f51565b636057361d811861003457602436103417610208576004355f55005b632578e0
    # f0811861020457602436103417610208576004356005811015610208576001015460405260206040f35b632e64cec1811861007d573
    # 4610208575f5460405260206040f35b635ec0e9c9811861020457602436103417610208576020806040526006600435600581101561
    # 02085702600601816040016040825482528060208301526001830181830160208254015f81601f0160051c600581116102085780156
    # 100f457905b808501548160051b8501526001018181186100de575b5050508051806020830101601f825f03163682375050601f1960
    # 1f82516020010116905090508101905090509050810190506040f35b6305a8568681186102045760443610341761020857600435600
    # 401803560648111610208575060208135018082604037505060243560e0526020604051018060406101005e50600660245460058110
    # 15610208570260060160e051815560206101005101600182015f82601f0160051c600581116102085780156101c257905b8060051b6
    # 101000151818401556001018181186101aa575b50505050506024356024546005811015610208576001015560245460018101818110
    # 61020857905060245560243560256040516060206020525f5260405f2055005b5f5ffd5b5f80fd006200180204012a0204841902168
    # 10a00a1657679706572830004000014', 'to': b''}

    # NOTE: The `to` field is empty for a reason
    # When you want to deploy a contract to the EVM ecosystem, 
    # just leave the `to` field blank
    
    #__________________________________________________________________________
    # SECTION: Sign the transaction

    signed_transaction = w3.eth.account.sign_transaction(transaction, private_key=PRIVATE_KEY)
    print(signed_transaction)

    # NOTE: This is just a signed transaction. 
    # It has not been sent to the blockchain yet.

    # WARNING: Don't leave signed transactions like this lying around.
    # This is okay for now as this transaction does not contain real funds

#______________________________________________________________________________
    # SECTION: Send transaction

    # transaction hash
    tx_hash = w3.eth.send_raw_transaction(signed_transaction.rawTransaction)
    print(f"My transaction hash is {tx_hash}")

    # SUB_SECTION: Wait for the transaction to be completed after sending it
    tx_receipt = w3.eth.wait_for_transaction_receipt(tx_hash)

    print(f"The contract has been deployed to {tx_receipt.contractAddress}")

#______________________________________________________________________________
    # SECTION: Send transaction

    def decrypt_key() -> str:
        with open(KEYSTORE_PATH, "r") as fp:
            encrypted_account = fp.read()
            password = getpass.getpass("Enter your password: ")
            key = Account.decrypt(encrypted_account, password)
            print("Decrypted key!")
            return key

#______________________________________________________________________________

if __name__ == "__main__":
    main()
