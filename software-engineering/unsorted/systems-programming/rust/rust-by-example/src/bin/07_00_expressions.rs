fn main() {

    //_________________________________________________________________________
    // SECTION: Statement

    // NOTE: This is a statement (specifically a `let statement`)
    // 1. statements end with a semicolon
    // 2. statements perform an action, like declaring a variable.
    // 3. statements do not return a value.

    let cool_number: i32 = 8;

    println!("cool_number: {}", cool_number);
    // cool_number: 8
    
    //_________________________________________________________________________
    // SECTION: A Statement that contains an expression

    // This is a statement that contains an expression.
    // The whole thing from `let` until the `;` is the expression statement.

    // {}; is used to create a block of code.

    // `cool_number + 2` is the part that is called an expression.
    // expressions must not end with a semicolon.

    // An expression is code that evaluates to a value, 
    // and then returns that value.

    // In Rust the last line inside a block is implicity returned if that
    // line is an expression

    // So cool_number + 2 is an expression that evaluates to 10,
    // and then that 10 is returned and assigned to the variable hot_number  

    let hot_number: i32 = {
        cool_number + 2
    };

    println!("hot_number: {}", hot_number);
    // hot_number: 10

    //_________________________________________________________________________
    // SECTION: A Statement with an explicit return statement

    // let ninja_stars: i32 = {
    //     return hot_number + cool_number;
    // };

    // This statement still contains an expression
    // The `hot_number + cool_number` part is the expression 
    // that evaluates to a value.

    // The only difference is that the expression has been put inside a return
    // statement and then that return statement is used to set the value
    // of ninja_stars

    // The syntax is `return expression;`
    // Since return is a statement, it ends with a semicolon.
    

    //_________________________________________________________________________
    
}
