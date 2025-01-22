// NOTE: Structure of a basic GO program

// Every Go program belongs to a package. 
// A package is a collection of functions and methods
// For an executable, the package has to be called "main"
package main

// This imports the "fmt" package from the standard library
// This is a allows you to use the Println function, 
// from the Go standard library
import "fmt"

// The main function is the entry point of every Go program
func main() {
    fmt.Println("Hello, World!")
}

//=============================================================================

// NOTE: How to run a go program

// To run this file, run this command:
// go run main.go
// This command will compile and run the program 
// without creating a binary executable.
// This is convient for just checking the output of something.

//=============================================================================

// NOTE: How to create a binary executable, and run it

// Go build main.go
// This will create a binary executable in same folder called "main".
// To run this binary executable, do this.
// ./main

//=============================================================================

