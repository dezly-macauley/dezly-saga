package main

import (
    "fmt"
    "sync"
)

func main() {

    var wait sync.WaitGroup
    
    // NOTE: In this example I know how many goroutines should be running
    numberOfGoRountines := 4
    wait.Add(numberOfGoRountines)
   
    // numberOfGoRountines = 4
    // Index:           0, 1, 2, 3
    // Goroutines:      a, b, c, d
    // NOTE: Because goroutines are executed concurrently,
    // the messages won't be recieved in order as


    for i := 0; i < numberOfGoRountines; i++ {

        // For each loop a go function will be created to print out a message
        // when that message is done.
        go func(goRoutineID int) {
            fmt.Printf("ID:%d - New message recieved\n", goRoutineID) 
            wait.Done()
        }(i)

    }

    wait.Wait()

    // NOTE: This an example of the output
    // The go routines are not executed in order

    // ID:3 - New message recieved
    // ID:0 - New message recieved
    // ID:1 - New message recieved
    // ID:2 - New message recieved

}
