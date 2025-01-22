// NOTE: What is a pointer?
// A pointer is an 8 byte long integer
// It points to a memory address

package main

import "fmt"

type Player struct {
    HP int
}

// NOTE: "Pass by value" vs "Pass by reference"


// NOTE: Pass by value: (player Player) 

// If you had used (player Player) then this would be a "pass" by value
// Then function would simply get a COPY of the instance of the player struct 
// and then modify that copy. 
// This may not be what you want because any changes made to this copy will
// not affect the original

// NOTE: Pass by reference: (player *Player)
// player *Player means that the variable player is a pointer that points to 
// the memory address where the instance of the struct is pointing to
// In simply terms, the function gets a REFERENCE to the instance of the 
// player struct.
// So the function is actually modifying the ORIGINAL and not a copy of the
// instance
// To use the function:
// takeDamage(&player, 10)

func takeDamage(player *Player, amount int) {
    player.HP -= amount
    fmt.Println("The player has taken damage. Their HP is now:",
        player.HP)
}

// NOTE: You could also create what is called a "function reciever" in Go
// This is a method that will be available to any 
// instance of the Player struct

// (instanceOfPlayer *player) A pointer is used here because this method will
// be used to modify the original instanceOfPlayer
func (instanceOfPlayer *Player) takeDamage(amount int) {
    instanceOfPlayer.HP -= amount
    fmt.Println("The player has taken damage. Their HP is now:",
        instanceOfPlayer.HP)
}

func main() {

    player := Player {
        HP: 100,
    }

    takeDamage(&player, 10)
    // The player has taken damage. Their HP is now: 90
    fmt.Println("The player's health is:", player.HP)
    // The player's health is: 90

    player.takeDamage(10)
    fmt.Println("The player's health is:", player.HP)
    // The player's health is: 80

//=============================================================================

    myString := "Bob is cool"
    
    // var myStringPtr *string = &myString
    myStringPtr := &myString

    fmt.Println(*myStringPtr)
    // Bob is cool

//=============================================================================

}
