// SECTION: Custom Data Types
// In Rust custom data types are formed mainly through the two keywords:
// 1. struct 
// 2. enum

// SECTION: Constants 
// These are variables that can be created using two keywords:
// 1. const
// 2. static

// NOTE: const
// 1. A const is always immutable. Its value can't is set at compile time
// and it can't be changed at run time.
// 2. Can be declared in any scope. So it does not have to be in a function.
const MAX_HEALTH: u32 = 1500;

fn main() {
    println!("Each fighter has a max of: {} health points", MAX_HEALTH);
    // Each fighter has a max of: 1500 health points
}


// TODO: Learn how `static` works
