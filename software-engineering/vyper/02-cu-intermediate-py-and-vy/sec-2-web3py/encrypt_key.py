# NOTE: This thes are built-in Python modules
import getpass
import json
from pathlib import Path

# NOTE: This is from the module `web3`
from eth_account import Account

KEYSTORE_PATH = Path(".keystore.json") 

def main():

    # This is similar to `input()` in Python.
    # The user will be prompted in the terminal to enter their private key
    # What they type will be hidden.
    private_key = getpass.getpass("Enter your private key: ")

    # NOTE: Next the private key needs to be transformed into 
    # an eth account object.
    # This is basically a variable type that has useful methods like
    # `.encrypt()`


    my_account = Account.from_key(private_key)

    password = getpass.getpass(
            "Enter a password to encrypt/encrypt your key:\n"
    )

    encrypted_account = my_account.encrypt(password)

    print(f"Saving to {KEYSTORE_PATH}")

    # This will open the file specified in KEYSTORE_PATH in write mode.
    # If the file does not exist, it will be created.
    # `fp` stands for file path.
    with KEYSTORE_PATH.open("w") as fp:
        json.dump(encrypted_account, fp)

if __name__ == "__main__":
    main()
