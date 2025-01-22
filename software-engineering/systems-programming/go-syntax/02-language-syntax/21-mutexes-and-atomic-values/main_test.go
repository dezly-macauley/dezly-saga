package main

// NOTE: Run this using:
// go test -v -count=1 --race
// count=1 means it won't be cached
// --race checks for race conditions

import (
    "fmt"
	"sync"
	"testing"
)

func TestState(t *testing.T) {

    // If you don't specify a field in a struct it will be set to the default
    // So count: 0
    state := &State{}

    // TODO: I need more information on weight groups
    wg := sync.WaitGroup{}

    for i := 0; i < 10; i++ {
       
        // Every time the for loop runs, 1 will be added to the wait group
        wg.Add(1)
       
        // this function accepts i 
        go func(i int) {
            state.count = i + 1
            state.setState(i + 1)
            wg.Done()
        }(i)
        // i is passed into the annonymous go function

    }

    wg.Wait()

    fmt.Printf("%+v\n", state)
}
