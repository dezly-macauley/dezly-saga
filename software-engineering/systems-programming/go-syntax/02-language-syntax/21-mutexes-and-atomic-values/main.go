// NOTE: State
// You will either set the state, update the state, or delete the state

package main

import (
	"sync"
)

type State struct {

    // TODO: What the actual heck is a Mutex

    // the variable name mu is short for mutex
    mu sync.Mutex 
    count int
}

func (instanceOfState *State) setState(i int) {
    instanceOfState.mu.Lock()
    defer instanceOfState.mu.Unlock()

    instanceOfState.count = i
}

func main() {

}
