// NOTE: Tuples
// 1. Can hold values of different types
// 2. Can hold any number of values 


// NOTE: Tuples can be used when you want 
// a function to return multiple values

fn double_both_scores(score_one: i32, score_two: i32) -> (i32, i32) {
    let doubled_score_one = score_one * 2;
    let doubled_score_two = score_two * 2;
    
    // In Rust the last line of a function is automatically returned
    (doubled_score_one, doubled_score_two)
}
