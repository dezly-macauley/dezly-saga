package main

// import "time"
import (
	"fmt"
	"time"
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
    // messagePrinter("Hello, World!") 

    // NOTE: So in simple terms the variable messagePrinter will behave like
    // a regular function

    // The go keyword turns each of these function calls into a go routine
    // so that they are executed concurrently
    go messagePrinter("Hello World")
    go messagePrinter("Hello goroutine")

    // This line is just to delay the completion of the main function by 
    // 1 second so that the goroutines can complete 
    time.Sleep(time.Second)

}
