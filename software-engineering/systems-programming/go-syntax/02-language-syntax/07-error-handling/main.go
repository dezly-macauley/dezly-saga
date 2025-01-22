package main

// NOTE: Error Handling
// You need to import the "errors" package from the standard library

import (
    "fmt"
    "errors"
)

type User struct {
    name, role, email string
    age int
}

// This is a method that will return a salary amount based on the "role"
// of the instance

// NOTE: Creating a custom method with a custom error messages

// The function will return two values when it is called
func (user_instance User) Salary() (int, error) {


    switch user_instance.role {
    case "Developer":
        // nil means nothing. No error
        return 100, nil

    case "Architect":
        return 200, nil

    case "":
        return 0, errors.New("Error: The person has not been assigned a role")

    default:
        return 0, errors.New(
            fmt.Sprintf("Error: '%s' is not a valid role", user_instance.role))
    }

} 

func main() {

    user1 := User {
        name: "Cassie",
        role: "Architect",
        email: "cassie@gmail.com",
        age: 40,
    }

    // NOTE: How to view the custom errors in the struct
    salary, error := user1.Salary()

    // If there was an error message, then print the error
    if error != nil {
        fmt.Println(error)
    } else {
        // otherwise, display the error
        fmt.Println("Salary:", salary)
    }

    // Error: 'Security' is not a valid role

}
