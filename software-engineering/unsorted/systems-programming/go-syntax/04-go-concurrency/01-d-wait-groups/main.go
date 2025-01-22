package main

// NOTE: How to use a goroutine without having to use "sleep" to delay the 
// "main" go routine (func main) 

import (
    "fmt"
    "sync"
)

func main() {

    // NOTE: Step 1: Create a variable to hold the WaitGroup
    // The purpose of wait.Wait() is to synchronize the execution of goroutines. 
    /// It prevents the main goroutine (func main) from exiting before 
    // all other goroutines have completed their tasks.

    var wait sync.WaitGroup

    // NOTE: Step 2: Before the goroutine starts, you let the wait group
    // know that it has to wait for one thing to finish
    wait.Add(1)

    // NOTE: Step 3: Start the goroutine that the WaitGroup should wait for

    // This is an annoymous function that has been made into a go routine
    go func() {
        fmt.Println("Hello World")
        
        // NOTE: Step 4: Signal to the WaitGroup that this goroutine 
        // has completed.
        // And decrements the wait counter by 1
        wait.Done()
    }()

    // The () after the curly braces is used to execute the annoymous function
    // The annoymous function in this example doesn't require any arguements
    // to work so the brackets are empty. If it did, then you would write
    // those values in the ()

    // NOTE: Step 5: 
    // wait.Wait() blocks the main goroutine (func main) from exiting until
    // the WaitGroup counter goes back to 0. 
    // Which will indicating all goroutines are done
    // To put it in simply terms for this example:
    // "func main() is not allowed to exit until go func is done"
    wait.Wait()

}
