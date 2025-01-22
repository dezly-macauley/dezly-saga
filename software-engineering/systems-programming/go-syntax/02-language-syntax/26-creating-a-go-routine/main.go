package main

import (
	"fmt"
	"time"
)


// NOTE: When you use the `go` keyword, this creates a new goroutine
// In simple terms, the program will be able to immeaditely execute 
// the next line of code while processing the code in the go routine

// NOTE: When you use the `go` keyword, you are not able to capture any return
// values from the function call

func sendEmail(message string) {

    // The function will skip over this line of code
	go func() {
		time.Sleep(time.Millisecond * 250)
		fmt.Printf("Email received: '%s'\n", message)
	}()

    // While everything inside of `go func` is being processed this next line
    // will immeaditely execute
	fmt.Printf("Email sent: '%s'\n", message)
    // Email sent: 'Message 1'

    // Then the response from the `go` func will be displayed 
    // Email received: 'Message 1'

}

// Don't touch below this line

func test(message string) {
	sendEmail(message)
	time.Sleep(time.Millisecond * 500)
	fmt.Println("========================")
}

func main() {
	test("Message 1")
	test("Message 2")
	test("Message 3")
}

