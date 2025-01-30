fn main() {
    // NOTE: Vector
    // 1. Allows you to store more than one value in a single data structure
    // that puts all of the values next to each other in memory
    // 2. Vectors are stored on the Heap 
    // 3. All the values in the Vector have to be of the same type 
    // 4. Example use case:
    // The lines of text in a file, the prices of items in a shopping cart

    //_________________________________________________________________________
   // SECTION: How to create an empty Vector 
    
    // This is a Vector that stores 32 bit integers 
    // Remember to add the keyword `mut` to make the variable mutable
    let mut item_costs: Vec<i32> = Vec::new(); 

    println!("Here is the cost of each item: {:?}", item_costs);
    // Here is the cost of each item: []

    //_________________________________________________________________________
    // SECTION: How to add value to a Vector

    item_costs.push(72);
    item_costs.push(15);
    item_costs.push(10);
    
    println!("Here is the cost of each item: {:?}", item_costs);
    // Here is the cost of each item: [72, 15, 10]

    //_________________________________________________________________________
    // SECTION: How to create a Vector with values 

    // index 0 = 12
    // index 1 = 26
    // index 2 = 73
    let player_scores: Vec<i32> = vec![12, 26, 73];
    println!("player_scores is {:?}", player_scores);
    // player_scores is [12, 26, 73]

    //_________________________________________________________________________
    // SECTION: How to read the elements of a Vector (using indexes)

    // NOTE: Remember to use `&`
    // This tells Rust that you do not want the variable second_score to
    // take ownership of the value in player_scores.
    // You simply want a reference to it (in simple terms, you just want to
    // view the element at that index). `&player_scores[1]`
    // _____________________________________________________________________
    // And that's why the variable type of second_score is `&i32`
    // second_score is a reference to an &i32 variable.
    // _____________________________________________________________________
    // `second_score` is stored on the stack, because second store is just
    // a memory address and the size of a memory address is known at compile
    // time. 
    // _____________________________________________________________________
    // The value that second_score points to, the original player_scores[1]
    // is stored on the heap
    // _____________________________________________________________________

    // [1] means index 1 which = 26
    let second_score: &i32 = &player_scores[1];
    println!("The second score is: {}", second_score);
    // The second score is: 26

    //_________________________________________________________________________
}
