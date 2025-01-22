package main

import "fmt"

func main() {
    
    // NOTE: A for loop

    // This will print 0 to 5 (including five)
    for i := 0; i <= 5; i++ {
        fmt.Println(i)
    }

    // NOTE: There is no such thing as a while loop in Go
    // If you want to create a data structure like this in Go, you would just
    // use a for loop.
    // This is because go allows you to ommit sections of a for loop syntax
    // So in Go a "while" loop is simply a for loop that is declared with only
    // a condition

    floorNumber := 0
    
    // This will print 0 to 5 (including five)
    for floorNumber <= 5 {
        fmt.Println("You are at floor number", floorNumber)
        floorNumber++
    }

}
