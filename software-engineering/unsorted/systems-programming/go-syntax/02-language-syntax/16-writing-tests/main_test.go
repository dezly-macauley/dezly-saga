package main

// NOTE: To use this file
// 1. Make sure that your terminal is in the root directory of the project 
// 2. `go mod init name-of-project-folder`
// 3. go test -v
// The -v means verbose. This will make it clear which tests Passed or Failed
// go test will test all the functions in a test file
// To run a specific file, use go test -run NameOfTestFunction -v

import (
    "testing"
    "fmt"
)

func TestOne(t *testing.T) {
    fmt.Println("hello from TestOne")
}

func TestTwo(t *testing.T) {
    fmt.Println("hello from TestTwo")
}

// NOTE: How to set a test to fail

func TestThree(t *testing.T) {
    t.Fail()
    fmt.Println("hello from TestThree")
}

func TestCalculateValues(t *testing.T) {
    var (
        expectedResult = 10
        x = 5
        y = 8
    )

    actualResult := calculateValues(x, y)

    if actualResult != expectedResult {
        // %d is for integers
        t.Errorf("Expected Result is %d but the Actual Result was %d",
            expectedResult, actualResult)
    }

}

