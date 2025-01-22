package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {

    start := time.Now()
    ctx := context.WithValue(context.Background(), 
        "username", "dezly-macauley")

    userId, err := fetchUserID(ctx)

    // If there was an error
    if err != nil {
        log.Fatal(err)
    }

    // If the function call succeeded
    fmt.Printf("The response was: %v\nAnd it took %+v long\n", 
        userId, time.Since(start))

}

func fetchUserID(ctx context.Context) (string, error) {

    // NOTE: How to set a time limit on how long the program should wait for
    // an answer from the server

    // If the server takes more than 100 Milliseconds to fetch the user ID
    // Then the request will be cancelled. 
    ctx, cancel := context.WithTimeout(ctx, time.Millisecond * 100) 

    defer cancel()

    val := ctx.Value("username")
    fmt.Println("The value is:", val)
    // The value is: dezly-macauley

    // This creates the struct of what will go through the channel
    type result struct {
        userId string
        err error
    }
    
    // NOTE: How to make a channel of type "result"
    // A channel allows goroutines to communicate with the rest of the 
    // program

    resultch := make(chan result, 1)

    // NOTE: A gorountine is created. 
    // Any code inside this block will be executed in the background will 
    // the next part of the code runs

    go func() {

        // The external call is made
        // This can return a respond and an error status
        // However the time that this info is filled in depends on how long 
        // the third party takes to respond
        res, err := thirdpartyHTTPCall()

        // This will write a result into the result channel
        resultch <- result {
            userId: res,
            err: err, 
        }

    }()

    select {
    // If the ctx.Done() flag is triggered then that means that the result
    // took longer than 100 Milliseconds
    case <- ctx.Done():
        return "", ctx.Err()
    // If next scenario is that the response from the external call was 
    // quick enough, and so the program can read from the channel 
    case res := <-resultch:
        // This will return the userId and the error status
        return res.userId, res.err
    }

}


// NOTE: What is an HTTP roundtrip?
// A roundtrip that happens from your computer to the server that you are 
// going to call. And then the server is going to respond and the value will
// come back to you


// NOTE: This thirdpartyHTTPCall is an example of making an external request
// and you have no control over how long that thing will take to respond

func thirdpartyHTTPCall() (string, error) {

    // This will create a delay of 500 Milliseconds
    time.Sleep(time.Millisecond * 90)
    return "user id 1", nil
}
