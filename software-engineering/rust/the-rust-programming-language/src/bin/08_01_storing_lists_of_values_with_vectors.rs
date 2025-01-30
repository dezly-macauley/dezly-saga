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

    // Adding a new element onto the end of the vector might require 
    // allocating new memory and copying the old elements to the new space, 
    // if there isn’t enough room to put all the elements next to each other 
    // where the vector is currently stored
    
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

    // WARNING: The risk of this method is that if you try to access an
    // index that does not exist, the program will crash immediately.
    // E.g. there are 3 values in the Vector. 
    // You start counting values starting from index 0:
    // - The first value is 12, and is stored at index 0
    // - The second value is 26, and is stored at index 1
    // - The third value is 75, and is stored at index 2
    // ---------------------------------------------------------------------
    // Trying to access index 6 (which does not exist) will 
    // cause the program to `panic` (stop working and exit).
    // let second_score: &i32 = &player_scores[6];


    // index 0 = 12
    // index 1 = 26
    // index 2 = 73

    //_________________________________________________________________________
    // SECTION: How to read the elements of a Vector (using the get method)
    
    let combo_hit_scores: Vec<i32> = vec![10, 50, 35];
    println!("combo_hit_scores = {:?}", combo_hit_scores);
    // combo_hit_scores = [10, 50, 35]

    // NOTE: This time I will declare the the score as an Option type,
    // because there is a chance that the index does not exist.

    // Option<&i32> is a special enum in Rust.
    // It means that the variable third_hit can either be an i32,
    // or `None` (if it does not exist)
    // An easy way to remember the syntax is to say:
    // The value of third_hit being an i32 is optional
    // -----------------------------------------------------------
    //The `.get()` method returns an Option<&T> 
    // So it can return either a variable that is the same type of
    // the elements in the collection, in this case an &i32
    // or it can return `None`
    let third_hit: Option<&i32> = combo_hit_scores.get(2);

    match third_hit {
        // If the third hit returns an element, then print this message
        Some(&element) => println!("The third hit is {}", element),
        // If the third hit returns an element, then print this message
        None => println!(
            "Element not found. There are only {} hits in the combo",
            combo_hit_scores.len()
        )
    }
    
    //_________________________________________________________________________
    // SECTION: Modifying the contents of a Vector

    // WARNING: One of the Ownership rules in Rust is that you can't have
    // an immutable (read-only) borrow, and a mutable borrow in the same scope.
    // The following code will NOT work


    // NOTE: 1. Vector is declared
    let mut arrow_hits: Vec<i32> = vec![10, 20, 30, 40, 50];
    // ----------------------------------------
    // The `mut` keyword is required if you want to modify the Vector

    // NOTE: 2. An immutable borrow is performed
    // let first_hit: &i32 = &arrow_hits[0];
    // ----------------------------------------
    // println!("The first hit is {}", first_hit);

    // NOTE: 3. A mutable borrow is performed
    // arrow_hits.push(6);
    // println!("The first hit is {}", first_hit);
    // --------------------------------------------------
    // `.push()` does not take ownership of the Vector `arrow_hits` but it 
    // modifies the original Vector.

    // This code will not work because you are attempting to modify 
    // the original vector while first_hit needs to read from the Vector.

    //_________________________________________________________________________
    // SECTION: To fix this, make sure that the mutable borrow 
    // happens first

    // mutable borrow
    arrow_hits.push(60);
    let first_hit: &i32 = &arrow_hits[0];
    println!("The first hit is {}", first_hit);
    // The first hit is 10

    //_________________________________________________________________________
    // SECTION: How to iterate over the values in a Vector
    // Immutable borrow
    
    let mut team_health_bars: Vec<i32> = vec![100, 32, 57];

    // NOTE: Remember to use a `&` because you don't want the for loop to
    // take ownership of the the Vector or the elements inside it.
    // This is an immutable borrow (read-only)
    for element in &team_health_bars {
        println!("{}", element);
    }
    // 100
    // 32
    // 57

    //_________________________________________________________________________
    // SECTION: How to iterate over the values in a Vector
    // mutable borrow
   
   // NOTE: Use `&mut` for a mutable borrow
    // This tells the Rust compiler:
    // I don't want to the for loop to take ownership of the Vector or the 
    // elements inside it but I want to modify the elements in the original
    // Vector.

    for element in &mut team_health_bars {
        *element = *element + 300; 

        // You can also write it like this:
        // *element += 300; 
    }

    println!("The updated team_health_bars are {:?}", team_health_bars);
    // The updated team_health_bars are [400, 332, 357]


    // NOTE: The dereference operator `*`
    // When you want to modify the actual value in a Vector you need to 
    // use this `*`
    // It tells Rust that you want access to the actual value of the element
    // and not a memory address that points to that value.
    
    //_________________________________________________________________________
    // SECTION: Using an enum to store multiple types 

    // In this example, I want to store rows of data from an SQLite table.
    // Each row has values that are of different types.
    enum PlayerDataRow {
        Name(String),
        HighScore(i32),
        AccountBalance(f64)
    }

    let player_one_data_row: Vec<PlayerDataRow> = vec![
        PlayerDataRow::Name(String::from("Seth")),
        PlayerDataRow::HighScore(628),
        PlayerDataRow::AccountBalance(43.89),
    ];

    // NOTE: Make sure to use `&` because you don't want the for loop to take
    // ownership of the Vector or the elements inside it.

    for element in &player_one_data_row {
        match element {
            PlayerDataRow::Name(value) => {
                println!("Name: {}", *value)
            },
            PlayerDataRow::HighScore(value) => {
                println!("High Score: {}", *value)
            },
            PlayerDataRow::AccountBalance(value) => {
                println!("Account Balance: {}", *value)
            }
        }
    }

    // Name: Seth
    // High Score: 628
    // Account Balance: 43.89

    // NOTE: You do not need to dereference value in println! because
    // println! does it for you.
    // I just prefer to be explicit

    //_________________________________________________________________________

}
