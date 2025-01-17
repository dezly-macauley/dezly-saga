# Import the external module called `titanoboa`
import boa

def main():
    print("Reading your Vyper code...")
    
    # This will read the contract and deploy it
    # favorites_contract = boa.load("favorites.vy")
    # Deploy the contract to a `pyevm` network!
    favorites_contract = boa.load("favorites.vy")
    starting_favorite_number = favorites_contract.retrieve()
    print(f"The starting favorite number is: {starting_favorite_number}")

    # NOTE: How to interact with the functions on a smart contract

    # print(f"The starting favorite number is: {starting_favorite_number}")

if __name__ == "__main__":
    main()
