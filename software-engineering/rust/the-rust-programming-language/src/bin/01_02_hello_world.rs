fn main() {
    // This is a Rust macro. 
    // Macros in Rust end with `!` after their name
    // This prints a message to the terminal
    // The semicolon at the end is how you tell Rust that you are ready to
    // end the expression.
    println!("Hello, world!");
}

// NOTE: `fn main() {}`
// __________________________________________________________________
// 1. `fn` 
// This means function. This is how you declare functions in Rust
// __________________________________________________________________
// 2. `main` 
// This is a special function in Rust.
// The code inside the `main` function is always the first code 
// that gets executed in every Rust program.
// The main function is called automatically 
// when you run the program
// __________________________________________________________________
// 3. `()`
// Inside the round brackets is where you define the 
// function parameters.
// The function parameters are variables that the function needs 
// to perform its tasks. 
// In this example it is empty because the main function does not
// need any variables to complete task of printing "Hello world".
// __________________________________________________________________
// 4. {}
// The function body.
// This is where you write down the tasks that the function needs
// to perform when it is called.
// __________________________________________________________________
// 5. Functions do not end with a semicolon because a function is 
// a declaration of a block of code.
// __________________________________________________________________

//_____________________________________________________________________________
// SECTION: How to run your code

/*
    NOTE: First you have to compile your code to an executable binary:
    rustc 01_02_hello_world.rs

    This will create a binary executable that looks like this

    NOTE: Then to run the binary do the following:
    ./01_02_hello_world
*/


//_____________________________________________________________________________
