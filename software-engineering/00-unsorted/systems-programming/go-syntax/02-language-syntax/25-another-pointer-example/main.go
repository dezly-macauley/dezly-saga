package main

import(
    "fmt"
    "strings"
)

// Because this function accepts a pointer to a string, it will modify the
// string directly, without needing to modify anything.
// So this is "Pass by reference"

// If the function signature was simply: removeProfanity(message string),
// the function would recieve a COPY of the message to modify.
// and the original message would not be changed
func removeProfanity(pointerToMessage *string) {

    // NOTE: The pointer is simply a memory address of where the value of that
    // strin is located
    // fmt.Println(pointerToMessage)    
    // 0xc000014070

    // NOTE: In order to modify the value that the pointer is pointing to
    // You need to deference it
    messageValue := *pointerToMessage
    // fmt.Println(messageValue)
    // You can go flip yourself! 

    messageValue = strings.ReplaceAll(messageValue, "flip", "****")
    
    // NOTE: Now the deferenced message must be updated
    *pointerToMessage = messageValue

}

func main() {
    
    message := "You can go flip yourself!"
    fmt.Println(message)

    removeProfanity(&message)
    fmt.Println(message)

}
