fn main() {

    //_________________________________________________________________________
    // SECTION: Scalar Types

    //_________________________________________________________________________
    // SUB_SECTION: Signed Integers
    // 1. A positive or negative number with no decimals
    // 2. The available types are:
    // i8, i16, i32, i64, i128 and isize (pointer size)
    // 3. Signed integers default to `i32` if you don't specify

    let player_score: i32 = 500;
    println!("player_score: {}", player_score);
    // player_score: 500

    let city_temperature: i32 = -5;
    println!("city_temperature: {}", city_temperature);
    // city_temperature: -5

    //_________________________________________________________________________
    // SUB_SECTION: Unsigned Integers
    // 1. Only positive numbers with no decimals 
    // 2. The available types are:
    // u8, u16, u32, u64, u128 and usize (pointer size)
    // 3. Unsigned integers default to `f64` if you don't specify

    let account_balance: f64 = 821.50;
    println!("account_balance: {}", account_balance);
    // account_balance: 821.5
    
    //_________________________________________________________________________
    // SUB_SECTION: Boolean (A variable that can be true or false)

    let is_premium_user: bool = true;
    println!("is_premium_user: {}", is_premium_user);
    // is_premium_user: true
    
    //_________________________________________________________________________
    // SUB_SECTION: The char type 
    // Any single Unicode character

    // Letters: 'A', 'z', 'é'
    // Digits: '1', '9'
    // Punctuation: '!', ',', '.'
    // Symbols: '@', '#', '$'
    // Whitespace: ' ', '\n' (newline), '\t' (tab)
    // Emojis: '🙂', '🌟'
    // Unicode Characters: '中', 'λ', 'Ω'
    
    // char Unicode scalar values like 'a', 'α' and '∞' (4 bytes each)

    // NOTE: Put char values inside of single qoutes (and not double qoutes)

    let emoji_choice: char = '🍑';
    println!("emoji_choice: {}", emoji_choice);
    // emoji_choice: 🍑

    //_________________________________________________________________________
    // SUB_SECTION: The unit type `()`
    // The unit type is basically an explicit variable type that means 
    // "This thing does not return anything". 
    // In other programming languages like C or Zig, 
    // this concept would be called `void`

    // The unit type (), whose only possible value is an empty tuple: ()

    // NOTE: Despite the value of a unit type being a tuple, 
    // it is not considered a compound type 
    // because it does not contain multiple values.

    // This is not the most practical use of a unit type.
    // The unit type is usually used by functions

    // TODO: I will discuss this further in the `functions section`

    let does_not_return_anything: () = ();
    println!("does_not_return_anything: {:?}", does_not_return_anything);
    // does_not_return_anything: ()

    // NOTE: The unit type does not implement `Display` so to print it,
    // you need to use the `debug` formater `{:?}`
    
    //_________________________________________________________________________
    // SECTION: Compound Types

    //_________________________________________________________________________
    // SUB_SECTION: Arrays
    // 1. All values in the array must be of the same type
    // 2. An array can only have values of the same type.
   
    // This is an array of that can hold five i32 numbers
    let list_of_numbers: [i32; 5] = [20, 12, 23, 41, 15];
    println!("list_of_numbers: {:?}", list_of_numbers);
    // list_of_numbers: [20, 12, 23, 41, 15]

    // NOTE: The array type does not implement `Display` so to print it,
    // you need to use the `debug` formater `{:?}`

    //_________________________________________________________________________
    // SUB_SECTION: Tuples

    // NOTE: 1. Tuples do not implement `Display` so to print it,
    // you need to use the `debug` formater `{:?}`
    
    let player_info: (i32, bool, f64) = (192, true, 43.20);
    println!("player_info: {:?}", player_info);
    // player_info: (192, true, 43.2)

    //_________________________________________________________________________

    // NOTE: 2. Destructuring a tuple
    // You can assign the values inside of a tuple to indivual variable
    // types

    let number_of_swords: i32 = 0; 
    println!("number_of_swords: {}", number_of_swords);
    // number_of_swords: 0

    let can_dual_wield: bool = false; 
    println!("can_dual_wield: {}", can_dual_wield);
    // can_dual_wield: false

    let player_money: f64 = 0.0; 
    println!("player_money: {}", player_money);
    // player_money: 0
    
    let (number_of_swords, can_dual_wield, player_money) = player_info;

    println!("number_of_swords: {}", number_of_swords);
    // number_of_swords: 192

    println!("can_dual_wield: {}", can_dual_wield);
    // can_dual_wield: true

    println!("player_money: {}", player_money);
    // player_money: 43.2

    //_________________________________________________________________________

}
