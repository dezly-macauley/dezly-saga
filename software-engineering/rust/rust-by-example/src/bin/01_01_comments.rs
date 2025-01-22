// SECTION: Single line comments

// This is a single line comment

//_____________________________________________________________________________
// SECTION: Multi-line comments

/*
    This is a block comment.
    For longer comments that take multiple lines
    
    It is especially useful for making paragraphs

*/

//_____________________________________________________________________________
// SECTION: Module-level (aka Crate-level) documentation comments `//!`
// Use this to create a high level overview documentation of the entire 
// module / crate.

//! This is the documentation for the entire module or crate.
//! It provides a brief overview of the program's functionality.

//_____________________________________________________________________________
// SECTION: Item-level documentation comments `///`
// Use this when you want to create documentation for a block of code

// This is used to document the item right below it

/// This is the main function.
/// It prints two lines to the console.

fn main() {
    println!("Hello World");
    println!("I'm a Rustacean");
}

// NOTE: To generate documentation use `cargo doc`

//_____________________________________________________________________________
