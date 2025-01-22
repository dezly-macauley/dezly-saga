package main

import "fmt"

//=============================================================================

// NOTE: Go does not have a specific keyword for enums
// You can create this data structure by first creating a custom variable 
// type and another const keyword

// NOTE: Step 1: First create a custom type
type WeaponType int

// NOTE: Step 2: Then create the enum using "const"

/*
    const (
        Axe WeaponType = 0
        Sword WeaponType = 1
        WoodenStick WeaponType = 2
        Knife WeaponType = 3
    )
*/

// There is a more convienient way to write this using the keyword "iota"
// "iota" basically gives each one a number starting from 0
// So Axe WeaponType = 0, Sword WeaponType = 1

const (
    Axe WeaponType = iota
    Sword
    WoodenStick
    Knife 
)

//=============================================================================

func getDamage(weaponType WeaponType) int {

    switch weaponType {
    case Axe:
        return 100
    case Sword:
        return 90 
    case WoodenStick:
        return 1
    case Knife:
        return 40 
    default:
        // panic will cauase the program to stop
        panic("Not a valid weapon")
    }

}

// NOTE: Over here I am using a function to create a custom method 
// for the custom variable type WeaponType
// This will allow you to print out the name of the enum

func (weaponTypeInstance WeaponType) WeaponName() string {

    switch weaponTypeInstance {
    case Axe:
        return "Axe"
    case Sword:
        return "Sword" 
    case WoodenStick:
        return "WoodenStick"
    case Knife:
        return "Knife" 
    default:
        // panic will cauase the program to stop
        panic("Not a valid weapon")
    }
}

// NOTE: It the function above did not exist
// Doing this: fmt.Printf("%s\n", Sword) 
// would lead to this
// %!s(main.WeaponType=1)

// But with the function... 
// fmt.Printf("%s\n", Sword.WeaponName()) 
// Leads to this
// Sword

// NOTE: You naming the function String will allow you to do this
// fmt.Printf("%s\n", Sword) 
func (weaponTypeInstance WeaponType) String() string {

    switch weaponTypeInstance {
    case Axe:
        return "Axe"
    case Sword:
        return "Sword" 
    case WoodenStick:
        return "WoodenStick"
    case Knife:
        return "Knife" 
    default:
        // panic will cauase the program to stop
        panic("Not a valid weapon")
    }
}


func main() {
    fmt.Println("Your attack did", getDamage(Axe), "damage!")
    // Your attack did 100 damage!

    fmt.Printf("%s\n", Sword.WeaponName()) 
    // Sword

    fmt.Printf("%s\n", Sword) 
    // Sword

}
