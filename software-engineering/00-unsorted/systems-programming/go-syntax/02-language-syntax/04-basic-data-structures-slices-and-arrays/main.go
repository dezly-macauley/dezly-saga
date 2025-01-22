package main

import "fmt"

func main() {

    // NOTE: What is a slice?
    /* 
        In Go, a slice is a dynamically-sized, 
        flexible view into the elements of an array.
        Slices are more powerful and convenient than arrays,
        and they are used extensively in Go programming.

        In Go slices are not fixed. They can grow and shink.
    */

    //      index     0,  1,  2
    numbers := []int{70, 20, 40}
    fmt.Println(numbers)
    // This will return the data structure
    // [70 20 40]
    // That println(numbers) will not give you the result you want.
    // You must use fmt.Println(numbers)

    // NOTE: How to create an empty slice
    // numbers2 := []int{}

    println(len(numbers))
    // 3

    println("The first number is", numbers[0])
    // The first number is 70

    first_element_in_numbers := 0
    println("The first number is", numbers[first_element_in_numbers])
    // The first number is 70

    println("The last number is", numbers[2])
    // The last number is 40

    last_element_in_numbers := len(numbers) - 1
    println("The last number is", numbers[last_element_in_numbers])
    // The last number is 40

    // NOTE: How to use range to go through a slice

    for index, value := range numbers {
        fmt.Println("Index:", index, "Value:", value)
    }
    // Index: 0 Value: 70
    // Index: 1 Value: 20
    // Index: 2 Value: 40

    println("=======================================")

    // NOTE: If you only wcant the value, add an underscore, then value
    // This is because range expects "index" first, and "value"
    
    for _, value := range numbers {
        fmt.Println("Value:", value)
    }
    // Value 70
    // Value 20
    // Value 40

    println("=======================================")
    // NOTE: If you only want the index, do this
    // Just leave out "value"
    
    for index := range numbers {
        fmt.Println("Index:", index)
    }
    // Index: 0 
    // Index: 1 
    // Index: 2 

    // NOTE: How to add values to a slice

    /*
        When you append an element to a slice, the append function may create
        a new underlying array if the existing array does not have enough 
        capacity to accommodate the new elements. This is why the result of
        append must be assigned back to a variable; otherwise, the new slice 
        (with the potentially new underlying array) would be lost.
    */

    numbers = append(numbers, 100, 10, 15, 9)
    fmt.Println(numbers)
    // [70 20 40 100 10 15 9]

//=============================================================================

    // NOTE: Arrays
    // Arrays have the same syntax as a slice. The difference is that arrays
    // have a fixed size

    // This array can only store 3 integer values
    fightRoundScores := [3]int{}  

    fmt.Println(fightRoundScores)
    // [0 0 0]

    // NOTE: How to add values to an array
    fightRoundScores[0] = 15
    fightRoundScores[1] = 8
    fightRoundScores[2] = 20
    fmt.Println(fightRoundScores)
    // [15 8 20]

//=============================================================================

    // NOTE: How to make a slice from an array
    fmt.Println("=============================")

    // Index:         0, 1, 2, 3, 4 
    myArray := [5]int{6, 3, 8, 6, 5}
    fmt.Println(myArray)
    // [6 3 8 6 5]

    // This means start from index 1 to 4 (excluding 4)
    mySlice := myArray[1:4]
    fmt.Println(mySlice)
    // [3 6 8]

    // NOTE: When you make a slice from an array,
    // it is "Pass by Reference"
    // In simple terms the slice is a reference to the original array.
    // So changing a value in the slice will affect the original

    mySlice[0] = 21
    fmt.Println(mySlice)
    // [21, 6, 8]

    fmt.Println(myArray)
    // [6 21 8 6 5]
    
//=============================================================================
    
    // NOTE: This also works the other way around
    // So modifying the original array will affect the slice

    myArray[1] = 80
    
    fmt.Println(myArray)
    // [6 80 8 6 5]

    fmt.Println(mySlice)
    // [80 8 6]

    // NOTE: Also note that if you pass a slice into a function it might 
    // get modified
    // Also multiple slices can reference the same array

//=============================================================================

    // NOTE: How to create a slice without creating an array first

    epicSlice := make([]int, 5)
    fmt.Println(epicSlice)
    // [0 0 0 0 0]

    // NOTE: You can use the keyword "cap" to print out the capacity of a 
    // slice
    fmt.Println(cap(epicSlice))
    // 5

    // len() and cap() return 0 when a slice is nil

//=============================================================================

}
