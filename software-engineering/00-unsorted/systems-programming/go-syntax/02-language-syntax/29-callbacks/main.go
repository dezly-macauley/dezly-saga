// NOTE: What is a callback?

// It is an anonymous function that will be executed within the context of
// another function

/*
    func toUpperSync(word string) string {

    }

    func toUpperSync(word string, f func(string)) string {

    }

    func toUpperSync(word string, f func(string)) string {
        f(strings.toUpper(word))
    }


*/

package main

import (
    "fmt"
    "strings"
)

func main() {

    toUpperSync("Pizza", 
        func(v string) {
            fmt.Printf("Callback: %s\n", v)
    })

    // NOTE: The output should be
    // Callback: PIZZA

}

// NOTE: This is a call back function
// It accepts two arguements
// "word" which is a string
// and "f" which is a function that accepts a string

func toUpperSync(word string, f func(string)) {
    f(strings.ToUpper(word))
}
