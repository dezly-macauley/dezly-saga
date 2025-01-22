fn main() {
    let mut num_peaches: u32 = 3;
    
    while num_peaches > 0 {
        num_peaches = num_peaches - 1;
        println!("There are {} peaches remaining", num_peaches);
    }

    // There are 2 peaches remaining
    // There are 1 peaches remaining
    // There are 0 peaches remaining

}
