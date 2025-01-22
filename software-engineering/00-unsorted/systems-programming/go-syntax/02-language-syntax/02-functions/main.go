package main

import (
	"errors"
	"fmt"
)

// In Go functions are declared using the keyword "func"
// Remember in Go that the variable type comes after the variable name
// This function does not return any value, it simply prints a message to
// the command line
// Just remember that the order of variable and its type are different than
// other programming languages.

// NOTE: A function that accepts one arguement and does not return a value

func say_hello(name string) {
    fmt.Println("Hello", name)
}

// NOTE: A function that accepts multiple arguments and does not return
// a value

func greetings(firstname string, lastname string) {
    fmt.Println("Greetings", firstname, lastname)
}

// NOTE: The function that accepts multiple values, 
// and returns a single value

// This part is known as the function signature:
// func add_two_numbers(number1 int, number2 int) int

func add_two_numbers(number1 int, number2 int) int {
    return number1 + number2
}

// NOTE: If the arguments are of the same type then you can list all of the
// arguments seperated by a comma and then add the variable type
func add_three_numbers(number1, number2, number3 int) int {
    return number1 + number2 + number3 
}

// NOTE: A function that accepts multiple arguments
// and returns multiple arguments

// This function will return two integers
func triple_and_return_both(number1 int, number2 int) (int, int) {
    a := number1 * 3
    b := number2 * 3
    return a, b
}

//=============================================================================

    // NOTE: How to ignore the return values of a function
    
    // This function returns multiple values
    // Let's say I wanted to call this function but I only care about the 
    // playeRank

    func getPlayerInfo() (string, int) {
        playerName := "Dezly" 
        playerRank := 18

        return playerName, playerRank
    }


//=============================================================================

// NOTE: Named returns
// You can also name the returned values in the function signature to
// to make a function easier to document

func getAmmoDetails() (arrows int, bullets int, bombs int) {
    arrows = 17
    bullets = 900
    bombs = 5

    return arrows, bullets, bombs
}
    
//=============================================================================

// NOTE: Early returns and gaurd clauses

// In math, you can't divide a number by 0 
// So if you have a function that divides two numbers then you need to 
// account for the fact that someone using this function could attempt to
// divide by zero

// This function always returns two values regardless of whether it failed,
// or succeeded
// 1. int will be the answer
// 2. error is an error status. 
// If the function succeeds it returns 
func divideTwoNumbers(number1 int, number2 int) (answer int, errorStatus error) {

    if number2 == 0 {
        return 0, errors.New("Can't divide by zero")
    }

    // If the function succeeds (The user did not try to divide by zero),
    // then the function will return the answer, and "nil" as the errorStatus
    // "nil" in this context simply means no error
    return (number1 / number2), nil
}     
    
//=============================================================================

// NOTE: The "main" function is the entry point of a Go program
func main() {
    
    // The functions say_hello and greetings do not return a value,
    // they simplpy perform an action
    say_hello("Obito")
    greetings("Max", "Payne")

    // The functions add_two_numbers and add_three_numbers actually return
    // values. So these values can be displayed by the println function
    println(add_two_numbers(2, 4))
    println(add_three_numbers(2, 4, 6))

    // The first return value of the function will be assigned to the first
    // variable on the left and so on.
    // E.g. This function wil return the numbers 9 and 15
    // So first_number is declared and given the value 9
    // So second_number is declared and given the value 15

    first_number, second_number := triple_and_return_both(3, 5)
    println(first_number, second_number)

//=============================================================================

// NOTE: What is an annoymous function?

    // This is function that is defined without a name, 
    // and its return value is set to be assigned to a specific variable

    number_of_sweets := func (red_sweets, green_sweets int) int {
        return red_sweets + green_sweets 
    }

    println(number_of_sweets(4,5))

//=============================================================================

    // NOTE: What is a closure? 
    // This is simply an annoymous function that uses variables from its
    // surrounding scope

    blue_gummies := 10

    number_of_gummies := func (pink_gummies int) int {
        return pink_gummies + blue_gummies
    }

    println(number_of_gummies(4))

//=============================================================================

    // NOTE: How to ignore the return values of a function
    // This function returns multiple values, playerName string, playerRank int
    // Let's say I wanted to call this function but I only care about the 
    // playeRank
    // _, means that I am ignoring the first return value of the function

_, rank := getPlayerInfo()

println(rank)

//=============================================================================

    // NOTE: Named returns

    _, bulletsLeft, _ := getAmmoDetails()
    fmt.Println("You have", bulletsLeft, "bullets left")
    // You have 900 bullets left

    _, _, bombsLeft := getAmmoDetails()
    fmt.Println("You have", bombsLeft, "bombs left")
    // You have 5 bombs left

    arrowsLeft, _, _ := getAmmoDetails()
    fmt.Println("You have", arrowsLeft, "arrows left")
    // You have 17 arrows left

//=============================================================================

}
