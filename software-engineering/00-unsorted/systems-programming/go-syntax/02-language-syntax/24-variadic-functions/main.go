// NOTE: A variadic function is simply a function that can take any arbitary
// number of arguements.
// It does this by using `...` in the function signature
// It recieves the arguements as a slice

//=============================================================================

package main

import "fmt"

// nums is passed into the function as a slice of integers
    
    // NOTE: An example of the function being used:

    // totalCost := sum(10, 80, 5)  
    // The numbers 10, 80, 5 are passed into the sum function as a slice

    // Index:                      0, 1,  2
    // E.g. myInputSlice := []int{10, 80, 5} 
    
    // len(myInputSlice) = 3 
    // 3 because it contains three numbers

func sum(nums...int) int {

    total := 0

    // NOTE: i++ happens AFTER the loop has performed the actions inside the
    // curly braces
    for i := 0; i < len(nums); i++ {
        total += nums[i]
    }
    // End of first loop: total = 10, i = 1
    // End of second loop: total = 90, i = 2
    // End of third loop: total = 95, i = 3
    // It will not loop again because i is now 3 which is not less than the 
    // len(nums)
    // The answer of 95 is returned
    return total

}

func main() {

    // The numbers 10, 80, 5 are passed into the sum function as a slice
    // E.g. myInputSlice := []int{10, 80, 5} 
    totalCost := sum(10, 80, 5)  
    fmt.Println(totalCost)
    // 95

}
