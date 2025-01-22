// NOTE: Mutext (Mutual Exclusion)

// A mutext ensures that only one Go routine can modify a value at once
// This helps to prevent data races.
// A data race is when two different parts of the program try to write 
// (change the value) of something at the same time.
// Or one part of the program is trying the read (view the value) of something
// while another program is trying to write from the program.

// NOTE: When it comes to mutexes, there are a couple of keywords
// sync Mutext
// lock()
// Unlock()

//=============================================================================

package main

import (
    "sync"
    "fmt"
)

type SafeCounter struct {
    mu sync.Mutex 
    NumMap map[string]int
}

func (s *SafeCounter) Add(num int) {
    s.mu.Lock()
    defer s.mu.Unlock()
    s.NumMap["key"] = num
}

func main() {

    s := SafeCounter{NumMap: make(map[string]int)}
    var wg sync.WaitGroup

    for i := 0; i < 100; i++ {
        wg.Add(1)
        go func (i int) {
            defer wg.Done()
            s.Add(i)
        }(i)
    }
    wg.Wait()
    fmt.Println(s.NumMap["key"])
}
