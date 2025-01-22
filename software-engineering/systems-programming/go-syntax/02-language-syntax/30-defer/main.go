package main

import "fmt"

func main() {

    // NOTE: The keyword "defer" means that this line of code will executed
    // when then "main" function has completed

	defer fmt.Println("Message A")

	fmt.Println("Message B")
}

// NOTE: 
// Message B
// Message A
