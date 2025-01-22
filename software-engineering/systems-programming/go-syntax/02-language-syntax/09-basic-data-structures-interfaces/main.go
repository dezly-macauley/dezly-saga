package main

import "fmt"

// NOTE: What is an interface?
// It is a collection of function signatures.

// If a struct has the same methods, and those methods return the 
// same variable type that is returned by the function signature, then Go
// will implicitly associate that struct with the interface

type employee interface {
    getName() string
    getSalary() int
}

//=============================================================================

// NOTE: This is a struct called "contractor"
// It has a getName() method that returns a string
// It also has a getSalary() method that returns an int
// Therefore Go will implicitly assoicate "contractor" as a type of "employee"

// It doesn't matter how the methods of "contractor" work, as long as they
// match the function signatures in the employee interfaces

type contractor struct {
    name string
    hourlyPay int
    hoursPerYear int
}

func (instanceOfContractor contractor) getSalary() int {
   return instanceOfContractor.hourlyPay * instanceOfContractor.hoursPerYear 
}

func (instanceOfContractor contractor) getName() string {
    return instanceOfContractor.name
}

//=============================================================================

// NOTE: The same thing applies to the fullTime struct

type fullTime struct {
    name string
    salary int
}

func (instanceOfFullTime fullTime) getSalary() int {
   return instanceOfFullTime.salary 
}

func (instanceOfFullTime fullTime) getName() string {
   return instanceOfFullTime.name
}

//=============================================================================

// NOTE: Here is a function that makes use of interfaces
// This function will work with any struct that satisfies the 
// conditions of the interface "employee"

// What are the conditions of the interface "employee"?

/*
    type employee interface {
        getName() string
        getSalary() int
    }
*/

// Any struct that has the ALL of the methods of the employee interfaces,
// and return the same variable types as those methods 

// E.g. the struct fullTime has the method getName() which returns a string,
// and it also has the method getSalary() which returns a string
// So without you having to do anything else, Go will implicitly consider the
// struct fullTime to be a type of "employee""

// E.g. the struct contractor has the method getName() which returns a string,
// and it also has the method getSalary() which returns a string
// So without you having to do anything else, Go will implicitly consider the
// struct contractor to be a type of "employee" as well

// NOTE: It does not matter that the getSalary() methods for fullTime and 
// contractor calculate the salary in different ways. What matters is that,
// they have a method call getSalary() AND this method returns an int

func test(e employee) {
    fmt.Println(e.getName(), e.getSalary())
    fmt.Println("=============================")

}

//=============================================================================

func main() {

    test(fullTime {
        name: "Jack",
        salary: 50000,
    })

    test(contractor {
        name: "Bob",
        hourlyPay: 100,
        hoursPerYear: 73,
    })
    
    test(contractor {
        name: "Jill",
        hourlyPay: 872,
        hoursPerYear: 982,
    })

    // NOTE: Output:

    /*
        Jack 50000
        =============================
        Bob 7300
        =============================
        Jill 856304
        =============================
    */
}

//=============================================================================
