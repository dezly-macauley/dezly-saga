package main

import "fmt"

func main() {

//=============================================================================
    
    // NOTE: For loop

    // Note that i++ is just a shorter way of writing i = i + 1
    // So i++ means increase i by 1
    
    // i starts at 0, and as long as i is less than 10, this for loop will
    // continue to increase i by 1
    // This will print out the numbers from 0 to 4

    for i := 0; i < 5; i++ {
        println(i)
    }
    // Output:
    /* 
        0
        1
        2
        3
        4
    */

    // NOTE: Personally I prefer to use <= in situations like this to make
    // things clearer

    // This will print out 0 to 4 and it is more readable
    for i:= 0; i <= 4; i++ {
        println(i)
    }

    // Output:
    /* 
        0
        1
        2
        3
        4
    */

    // NOTE: Using a for loop with a slice.
    // index          0,  1, 2, 3,  4  
    numbers := []int{18, 20, 5, 52, 14}
    lengthOfnumbers := len(numbers)
    fmt.Println("====================")
    fmt.Println("There are", lengthOfnumbers, "numbers")
    // There are 5 numbers
    
    // Where to start | i:= 0 | Perform this action from index 0
    // When to stop | i < len(numbers) | Keep going until i is greater that
    // the number of items in the slice
    // How to proceed | i++ increase i by 1 and continue
    for i:= 0; i < len(numbers); i++ {
        fmt.Println("Index", i, "is", numbers[i])
    }
    // Index 0 is 18
    // Index 1 is 20
    // Index 2 is 5
    // Index 3 is 52
    // Index 4 is 14
    fmt.Println("====================")

    // NOTE: How to break out of a loop
    names := []string{"Cindy", "Jake", "Waldo", "Kim", "Ken"}

    // The format when using rang is index, value. 
    // _, value is how you tell Go that you don't care about the index number,
    // you only want the string value at each index
    for _, value := range names {

        if value == "Waldo" {
            fmt.Println("Found Waldo")
            break
        } else {
            fmt.Println(value)
        }

    }
    // Cindy
    // Jake
    // Found Waldo

//=============================================================================

    // NOTE: If statement
    // Remember that in Go you use == to check a condition
    // For multiple conditions use:
    // && which means and 
    // || which means or 
    // < Less than
    // <= Less than or equal to
    // > Greater than
    // >= Greater than or equal to
    // != Not equal

    hours_worked := 0
    var worked_overtime bool 

    if hours_worked > 8 {
        worked_overtime = true
        println(worked_overtime)
    } else if hours_worked == 0 {
        worked_overtime = false
        println("You did not come to work")
    } else {
       println("You did not work overtime") 
    }

//=============================================================================

    // NOTE: Switch Statement

    var customer = "young adult"
    var entry_fee int

    // TODO: Find out how to get this function done

    switch customer {
        case "adult":
            entry_fee = 50
            println(entry_fee)
        case "child":
            entry_fee = 20
            println(entry_fee)
        case "teenager", "young adult":
            entry_fee = 30
            println(entry_fee)
        default:
            // This is where you tell the program what to do if none of the
            // other options check out
            println("Please enter a valid customer")
    }

//=============================================================================

    // NOTE: Shorthand if statement

    /*
        if intial_statement; condition {
        
        }

    */

    // NOTE: Here's the long way

    /*
        account1 := 5
        account2 := 19

        totalFunds := account1 + account2

        if totalFunds > 10 {
            fmt.Println("Total funds is greater than 10")
        }

    */

    // NOTE: Shorthand

    account1 := 5
    account2 := 19

    // Total Funds is declared, given a value and then assigned all on the
    // same line
    // However the variable totalFunds is only accessible within the if
    // block, so don't do this if you need to use totalFunds somewhere else
    // in your code
    if totalFunds := account1 + account2; totalFunds > 10 {
        fmt.Println("Total funds is greater than 10")
    }


//=============================================================================

}
