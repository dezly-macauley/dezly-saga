// NOTE: What is a Go routine?
// It is a function that is being scheduled by the Go scheduler
// Go routines do not run in parallel

package main

import(
    "time"
    "fmt"
)

func main() {

    // NOTE: There are two types of channels
    // 1. unbuffered channel
    // 2. buffered channel

    // A channel in Go will always block if it is full
    // Think of a channel as a box, and inside that box you can place cookies
    // With an unbuffered channel, you have no box
    // When an unbuffered channel is full, the code below it will never
    // execute

    // With a buffered channel you have a box where you can fit in a number
    // of cookies

    // NOTE: How to make an unbuffered channel in Go

    // chan is a keyword in Go. It means channel
    resultChannel := make(chan string)

    // NOTE: How to make a buffered channel
    // This has a buffer size of 10
    // So it can fit 10 cookies

    // bufferedChannel := make(chan string, 10)

    // NOTE: This is a go routine
    go func() {
        result := <- resultChannel
        fmt.Println(result)
    }()

    // NOTE: How to write a value into the result channel

    // This means write things into that channel
    // "foo" is being written into that result channel
    resultChannel <- "foo"

    // the prefix `go` turns the function fetchResource into a Go routine
    // This will now become an asynchronous call. 
    // In simple terms, the rest of the program can run without waiting for it
    // if they are not dependant on its outcome
    // go fetchResource()

    
    // fmt.Println("boom")
    // fmt.Println("Wait for it...")
    // result := fetchResource()
    // fmt.Println("The result is:", result)

}

func fetchResource(n int) string {
    // This will delay the output of the next line by 2 seconds
    time.Sleep(time.Second * 2)
    return fmt.Sprintf("result %d", n)
}
