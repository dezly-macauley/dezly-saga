// An attribute to hide warnings for unused code.
#![allow(dead_code)]

//_____________________________________________________________________________
// SECTION: 1. Creating a regular struct

// This is an attribute to make the struct printable
#[derive(Debug)]
struct Person {
    name: String,
    age: u8
}

//_____________________________________________________________________________
// SECTION: 2. Creating a tuple struct 
// This is basically a struct without named fields


// NOTE: I personally don't find these useful

// E.g. PlayerLocation variable is made up of two i32 varible types
// an x co-ordinate, and a y co-ordinate
struct PlayerLocation(i32, i32);

//_____________________________________________________________________________
// SECTION: 3. Creating a unit struct (A stuct with no fields)

struct MyEmptyStruct;

//_____________________________________________________________________________

fn main() {
    
    //_____________________________________________________________________________
    // SUB_SECTION: How to use a regular struct 

    let player_one: Person = Person {
        name: String::from("Peter"),
        age: 30
    };

    // {:?} This tells println! to print player_one using `Debug` formatting.
    // In simple terms, show me what is inside the player_one struct in a
    // human readable way.
    println!("player_one: {:?}", player_one);
    // player_one: Person { name: "Peter", age: 30 }

    println!("player one's name is: {}", player_one.name);
    // player one's name is: Peter

    //_____________________________________________________________________________
    
}

//_____________________________________________________________________________
