// NOTE: When do you use typed functions and when do you use an interface?
// Use an interface when you want to have state 
// With a typed function there is no way to store state

// TODO: I need a better understanding of this

package main

import "fmt"

// NOTE: This allows you to use string methods like strings.ToUpper()
// which will make a string uppercase
import "strings"

//=============================================================================

type TransformFunc func(s string) string

func Uppercase(s string) string {
    return strings.ToUpper(s)
}

func Prefixer(prefix string) TransformFunc {
    return func(s string) string {
        return prefix + s
    }
}

func transformStrings(s string, fn TransformFunc) string {
    return fn(s)
}

func main()  {
    fmt.Println(transformStrings("dezly macauley", Uppercase))  
    // DEZLY MACAULEY
    fmt.Println(transformStrings("dezly macauley", Prefixer("Captain ")))  
    // Mr dezly macauley
    // Captain dezly macauley

}
