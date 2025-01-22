// Every program that is in the root directory should be "package main"
package main

import "fmt"

import "03-how-to-create-a-new-project/types"

import "03-how-to-create-a-new-project/util"

// NOTE: Go run is for running a single file
// E.g. go run main.go
// E.g. go run util.go

func main() {

    // NOTE: How to use the imported package
    user := types.User {
            Username: util.GetUserName(),
            Age: util.GetNumber(),
    }

    // fmt.Println("the number is:", user)
    fmt.Printf("the user is %+v\n", user)
    // the user is {Username:James Age:36}

}

// NOTE: To build your program, you need to use:
// go build -o what-you-want-your-program-to-be-called
// E.g. go build -o my_program

// This will create a binary executable called `my_program`  
// The output will be:
// the number is: {James 36}

