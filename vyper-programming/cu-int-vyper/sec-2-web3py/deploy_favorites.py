# This will allow this program to use the `compile_code` function from the
# vyper module to compile `favorites.vy` into bytecode.
from vyper import compile_code
from web3 import Web3

# This will give this program a way to create a smart contract from bytecode

def main():
    print("Reading the Vyper code for deployment...")
    print("")
    
    # NOTE: This is how to open a file in Python in read-only mode

    with open("favorites.vy", "r") as favorites_file:
        favorites_code = favorites_file.read()
        compliation_details = compile_code(favorites_code, output_formats=["bytecode"])
        print(compliation_details)

    # NOTE: This will output this

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
    # 4000014'}

    
    # SECTION: Turn the compliation_details output into a contract object

    # First create a web3 object to use its functionality
    # To initialize a web3 instance. You have to connect to a blockchain node 
    w3 = Web3(Web3.HTTPProvider("https://virtual.sepolia.rpc.tenderly.co/ddaaa368-02a0-4f86-9f62-aa10cb01e7e9"))
    
    # NOTE: The link above is from a fake sepolia block node created on Tenderly
    # https://dashboard.tenderly.co/dezly-macauley/cyfrin-updraft-first-vyper-contract/testnet/df5ea98c-2b12-4755-b0b3-b0aeb410fa86

    # This will return a contract object
    favorites_contract = w3.eth.contract(bytecode=compliation_details["bytecode"])
    print(favorites_contract)

if __name__ == "__main__":
    main()
