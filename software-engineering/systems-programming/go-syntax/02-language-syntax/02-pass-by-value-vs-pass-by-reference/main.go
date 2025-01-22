package main

import "fmt"

func main() {

//=============================================================================

    // NOTE: What is a variable?
    // A variable is simply a named reference to a value at a specific address

    sweetsAvailable := 50
    
    // In this example you placing the value of 50 at a unique memory address
    // and then you are naming this memory address "healthPotions" to make
    // it convienient to accesss this value of 50

    fmt.Printf(
        "The value of sweetsAvailable is located at the memory addres: %p\n",
        &sweetsAvailable)
    // The value of sweetsAvailable is located at the memory addres: 0xc0000120c0

    fmt.Printf("The value of sweetsAvailable is: %d\n", sweetsAvailable)
    // The value of sweetsAvailable is: 50

//=============================================================================

    // NOTE: Even when you change the value of the variable, 
    // the memory address stays the same

    println("==========================================")
    sweetsAvailable = 80
    fmt.Printf(
        "The value of sweetsAvailable is located at the memory addres: %p\n",
        &sweetsAvailable)
    // The value of sweetsAvailable is located at the memory addres: 0xc0000120c0

    fmt.Printf("The value of sweetsAvailable is: %d\n", sweetsAvailable)
    // The value of sweetsAvailable is: 80

//=============================================================================

    // NOTE: Pass by value
    // Pass by value is when a variable gets its own COPY of a value
    // In Go variables are passed by value

    println("==========================================")

    sweetsAvailable = 100

    pizzaAvailable := sweetsAvailable 

    fmt.Printf(
        "The value of sweetsAvailable is located at the memory addres: %p\n",
        &sweetsAvailable)
    // The value of sweetsAvailable is located at the memory addres: 0xc0000120c0

    fmt.Printf("The value of sweetsAvailable is: %d\n", sweetsAvailable)
    // The value of sweetsAvailable is: 100
  
    // NOTE: The value of 100 from sweetsAvailable is COPIED to 
    // the pizzaAvailable. When pizzaAvailable gets this value, this 100
    // is stored in a different memory location.
    // So in simple terms pizzaAvailable gets it own copy of the value 100
    // that is NOT linked to the original value at sweetsAvailable

    fmt.Printf(
        "The value of pizzaAvailable is located at the memory addres: %p\n",
        &pizzaAvailable)
    // The value of pizzaAvailable is located at the memory addres: 0xc0000120c8

    fmt.Printf("The value of pizzaAvailable is: %d\n", pizzaAvailable)
    // The value of pizzaAvailable is: 100

//=============================================================================

    // NOTE: Because this was a "Pass by value" situation
    // Changing the value of pizzaAvailable will not affect sweetsAvailable
    // in any case

    pizzaAvailable = 8
    
    fmt.Printf(
        "The value of sweetsAvailable is located at the memory addres: %p\n",
        &sweetsAvailable)
    // The value of sweetsAvailable is located at the memory addres: 0xc0000120c0

    fmt.Printf("The value of sweetsAvailable is: %d\n", sweetsAvailable)
    // The value of sweetsAvailable is: 100

//=============================================================================

    // NOTE: Pass by reference
    // This is when a variable DOES NOT get it own unique copy of a value,
    // but it instead gets a reference to the memory address of a value.

    println("======================================")

    bookOne := "Aylo"

    // NOTE: How to create a pointer using the var keyword
    // bookTwo *string = &bookOne

    bookTwo := &bookOne


    fmt.Printf(
        "The value of bookOne is located at the memory addres: %p\n",
        &bookOne)
    // The value of bookOne is located at the memory addres: 0xc000014070

    fmt.Printf(
        "The value of bookTwo is located at the memory addres: %p\n",
        bookTwo)
    // The value of bookTwo is located at the memory addres: 0xc000014070

    fmt.Printf("The value of book One is: %s\n", bookOne)
    // The value of book One is: Aylo

    // NOTE: Because bookTwo is a pointer, you have to deference it with a *
    // to print its value
    fmt.Printf("The value of book Two is: %s\n", *bookTwo)
    // The value of book Two is: Aylo

//=============================================================================

    // NOTE: Because this is pass by reference. Both variables are pointing
    // to the same address in memory. So updating bookOne or bookTwo will
    // update the other

    println("===================================")
    bookOne = "The Maverick Amorist"

    fmt.Printf(
        "The value of bookOne is located at the memory addres: %p\n",
        &bookOne)
    // The value of bookOne is located at the memory addres: 0xc000014070

    fmt.Printf(
        "The value of bookTwo is located at the memory addres: %p\n",
        bookTwo)
    // The value of bookTwo is located at the memory addres: 0xc000014070

    fmt.Printf("The value of book One is: %s\n", bookOne)
    // The value of book One is: Aylo

    // NOTE: Because bookTwo is a pointer, you have to deference it with a *
    // to print its value
    fmt.Printf("The value of book Two is: %s\n", *bookTwo)
    // The value of book Two is: Aylo

 //=============================================================================

}
