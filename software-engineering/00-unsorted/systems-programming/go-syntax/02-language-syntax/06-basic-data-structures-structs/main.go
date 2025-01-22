package main

import "fmt"

// NOTE: What is a struct

// Think of it as a way as creating your own custom data type that even has
// its own custom methods


// NOTE: How to declare a struct

type User struct {
    // If you have several value that are of the same type, you could do this:
    // Name, Role, Email string
    first_name string
    last_name string
    role string
    email string
    age int
}

// NOTE: How to create a method for a struct

// (user_instance User) is just a way to indicate that this is not a seperate
// function, but a method that is part of the design of the 
// User struct
// SetEmail can only be used on an instance of the user struct 
// *User means that it is referencing the original

// NOTE: When to use (user_instance *User) vs (user_instance User)
// User passes a copy of user_instance
// *User is a pointer that is referencing user_instance

// Use (user_instance User) when you are creating a method that 
// only reads information from the struct

// Use (user_instance *User) when you are creating a method that 
// actually updates information in the struct

// NOTE: In Go these functions are called "function recievers"

func (user_instance *User) SetEmail(_email string) {
    user_instance.email = _email 
}

func (user_instance User) GetFullName() {
    // fmt.Println will automatically add a space between values seperated
    // by a comma
    fmt.Println(user_instance.first_name, user_instance.last_name) 
}

//=============================================================================

// NOTE: Creating an annoymous struct inside of another struct

type fightingGameCharcter struct {
    Name string
    Attack float64
    Defence float64
    HP int
   
    // annoymous struct
    specialMoves struct {
        grab string
        finisher string
        ultimate string
    }

} 

//=============================================================================


func main() {

    // NOTE: How to create an instance of a struct

    user1 := User{
        first_name: "Kevin", 
        last_name: "Miller", 
        role: "Programmer",
        email: "kevmiller@gmail.com",
        age: 30,
    }

    // NOTE: How to create an empty instance of a struct
    user2 := User{}
    fmt.Println(user2)

    // NOTE: How to get a value from a struct
    println(user1.age)
    // 30

    // NOTE: How to use a method that a struct has
    user1.SetEmail("kmiller@proton.me") 

    fmt.Println(user1.email)
    // kevmiller@gmail.com

    user1.GetFullName()

    // NOTE: How to print a struct
    fmt.Printf("user1's details are %v\n", user1)
    // user1's details are {Kevin Miller Programmer kmiller@proton.me 30}

    // You can also use %+v to make the output more verbose
    fmt.Printf("user1's details are %+v\n", user1)
    // user1's details are {first_name:Kevin last_name:Miller role:Programmer email:kmiller@proton.me age:30}

    // NOTE: How to create an annoymous struct
    // This is for situations where you have no reason to create more than
    // one instance of the struct

    myVillain := struct {
        Name string 
        Power string 
    } {
        Name: "Kenny",
        Power: "Ice",
    } 

    fmt.Println(myVillain)
    // struct { Name string; Power string }

    newFighter := fightingGameCharcter {}
    newFighter.specialMoves.ultimate = "Twin Dragon Slice"
    fmt.Println(newFighter)

}
