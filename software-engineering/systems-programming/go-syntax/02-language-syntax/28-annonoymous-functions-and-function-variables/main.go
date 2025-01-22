package main

// import "time"
import (
    "fmt"
)

func main() {

    // Declare messagePrinter as a variable of type func(string)
    var messagePrinter func(string)
    
    // Assign an anonymous function to messagePrinter
    messagePrinter = func(msg string) {
        fmt.Println(msg)
    }

    // Call the function via the variable
    // This will call the anonymous function and print "Hello, World!"
    messagePrinter("Hello, World!") 

}
