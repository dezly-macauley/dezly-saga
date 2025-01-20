// NOTE: Vectors
// 1. These are basically arrays but they don't have a fixed size
// 2. Vectors are stored on the heap because their size is not known at
// compile time
// ---------------------------------------------------------------------
// A vector is represented using 3 parameters
// 1. A pointer to the data
// 2. length
// 3. capacity - how much memory is reserved for the vector. 
// A vector an grow as long as the length is smaller than the capacity. 
// When this threshold needs to be surpassed,
// the vector is reallocated with a larger capacity.
// ---------------------------------------------------------------------

fn main() {

    //-------------------------------------------------------------------------
    // SECTION: How to declare a vector


    let mut combo_damage_scores: Vec<i32> = Vec::new();
    
    //-------------------------------------------------------------------------
    // SECTION: How to add items to a Vector 

    combo_damage_scores.push(17);
    combo_damage_scores.push(48);
    combo_damage_scores.push(25);

    // NOTE: How to print Vectors
    // 1. Vectors already have the attribute #[derive(Debug)] by default
    // 2. So to print them all you have to do is use {:?},
    // or {:#?} if you want the pretty-print output.

    // This is the `debug` format.
    // It prints a data structure in one line.
    println!("combo_damage_scores: {:?}", combo_damage_scores);
    // combo_damage_scores: [17, 48, 25]
   
    // This is the `pretty-print` format.
    // It prints the data structure in multiple lines and uses indentation
    // to make it more readable.
    println!("combo_damage_scores: {:#?}", combo_damage_scores);
    // combo_damage_scores: [
    //     17,
    //     48,
    //     25,
    // ]

    //-------------------------------------------------------------------------
    // SECTION: How to remove items from the end of a Vector

    combo_damage_scores.pop();
    println!("combo_damage_scores: {:?}", combo_damage_scores);
    // combo_damage_scores: [17, 48]

    //-------------------------------------------------------------------------
    // SECTION: How to get the length of a Vector

    // NOTE: usize = unsigned size
    // 1. Remember that Vectors are stored on the heap
    // --------------------------------------------------------------------
    // 2. The variable `combo_damage_scores` stores a pointer to the 
    // memory address where the data of the Vector is stored, 
    // the length of the Vector, and the capacity of the Vector
    // --------------------------------------------------------------------
    // 3. usize is that variable type used to store the length of a Vector
    // because usize matches the size of the pointer on the machine that
    // is running the code.
    // --------------------------------------------------------------------
    // 4. On a `32 bit system` usize is 32 bits
    // 5. On a `64 bit system` usize is 64 bits
    // --------------------------------------------------------------------

    let num_hits_in_combo: usize = combo_damage_scores.len();
    println!("There are {} hits in the combo", num_hits_in_combo);
    // There are 2 hits in the combo

    //-------------------------------------------------------------------------
    // SECTION: How to read a specific index in a Vector

    println!("combo_damage_scores: {:?}", combo_damage_scores);
    // combo_damage_scores: [17, 48]

    // Indexing starts at 0
    // index 0 = 17
    // index 1 = 48

    println!(
        "The first hit in the combo did {} damage!", 
        combo_damage_scores[0]
    );

    // The first hit in the combo did 17 damage!

    //-------------------------------------------------------------------------
    // SECTION: How to create a Vector that already has items inside
    // using the `vec![]` macro

    /*  
        Normally you would create a Vector like this:

        let mut combo_damage_scores: Vec<i32> = Vec::new();

        And then you would use `.push` method to add each item:
        combo_damage_scores.push(17);
        combo_damage_scores.push(48);
        combo_damage_scores.push(25);

        But Rust has a way to do this all in one go.
    */

    let mut magic_attack_scores: Vec<i32> = vec![16, 78, 32];
    println!("magic_attack_scores: {:?}", magic_attack_scores);
    // magic_attack_scores: [16, 78, 32]
    
    //-------------------------------------------------------------------------
    // SECTION: How to iterate over a Vector 
    // and get the value of each element 
    // iterate = go through each element in the Vector

    // NOTE: `.iter()` - Use when you only want the value of each element
    // 1. This gives you an immutable reference to each element in the Vector
    // 2. In simple terms you get a read-only version of each element.

    for value in magic_attack_scores.iter() {
        println!("The value is {}", value);
    }
    // The value is 16
    // The value is 78
    // The value is 32

    //-------------------------------------------------------------------------

    // SECTION: How to iterate over a Vector and get the index and value
    // of each element

    for (index, value) in magic_attack_scores.iter().enumerate() {
        println!("The item at index {} is {}", index, value);
    }
    // The item at index 0 is 16
    // The item at index 1 is 78
    // The item at index 2 is 32

    // NOTE: `iter().enumerate()`
    // 1. Use when you want both the the index and the value
    // 2. This is returns a tuple that keeps track of both the index, and
    // the value of each element
    
    //-------------------------------------------------------------------------
    // SECTION: How to modify each element in a Vector 

    // NOTE: `.iter_mut()`
    // This gives you a mutable reference of each value in the Vector
    // `.iter_mut()` will actually allow you to change the values in the 
    // original.

    println!(
        "The damage of the original attack is {:?}",
        magic_attack_scores
    ); 
    // The damage of the original attack is [16, 78, 32]

    for value in magic_attack_scores.iter_mut() {

        // *value: First you have to deference the reference.
        // In simple terms you are saying "I want access to the value at
        // this reference"
        // The code below will add 10 to each value in the Vector
        *value = *value + 10
    }

    println!(
        "The damage of the original attack is {:?}",
        magic_attack_scores
    ); 
    // The damage of the original attack is [26, 88, 42]


    //-------------------------------------------------------------------------

}
