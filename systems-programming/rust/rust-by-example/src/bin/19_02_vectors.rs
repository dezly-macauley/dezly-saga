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

    println!("combo_damage_scores: {}", combo_damage_scores);


    //-------------------------------------------------------------------------

}
