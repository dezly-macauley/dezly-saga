// NOTE: Attributes tell the Rust compiler 
// how to treat specific parts of your code,
// or how to treat the whole file.

//_____________________________________________________________________________
// NOTE: There are two types of attributes
// -------------------------------------------------------
// 1. `#[attribute]` - This tells Rust how to treat item 
// that is below this line
// -------------------------------------------------------
// 2. `#![attribute]` - This tells Rust how to treat
// all code inside the file

//_____________________________________________________________________________
// SECTION: Example 1 - #![allow(unused_variables)]

// NOTE: This attributes tells Rust not show warnings about unused variables,
// and to apply this setting for the entire file

#![allow(unused_variables)]

//_____________________________________________________________________________

fn main() {
    
    // NOTE: Normally if you tried to compile this code,
    // the Rust compiler would give you a warning about unused code.
    // This is because by default the Rust compiler has a trait called:
    // `#[warn(unused_variables)]` enabled on a Rust file as a default
    // settting.
    let number_of_users: i32 = 87;

    //_________________________________________________________________________
    





}


//_____________________________________________________________________________
