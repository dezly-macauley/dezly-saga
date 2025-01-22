package main

import(
    // "time"
    "fmt"
)

func main() {

    // NOTE: How to create a buffered channel

    // Messages will be placed into this message channel and then other 
    // parts of code can read the messages from this channel and process them
    // This is a buffered channel because it has a limit of 128 
    // So this is 128 strings
    // It's a good practice to use powers of 2
    messageChannel := make(chan string, 128)

//=============================================================================


    // NOTE: How to write to a channel
    
    messageChannel <- "Hey it's Tsunade"
    messageChannel <- "Want to hang out?"
    messageChannel <- "I'm free at 9 PM"

    // NOTE: Remember to close a channel when you are done writing from it
    // If you don't add this line your program may stop working
    // fatal error: all goroutines are asleep - deadlock!

    close(messageChannel)

//=============================================================================

    // NOTE: How to read from a channel mannually
    // This will read a new message each time 

    /* 


        message := <- messageChannel
        fmt.Println("Message recieved:", message)
        // Message recieved: Hey it's Tsunade
        
        message = <- messageChannel
        fmt.Println("Message recieved:", message)
        // Message recieved: Want to hang out?

        message = <- messageChannel
        fmt.Println("Message recieved:", message)
        // Message recieved: I'm free at 9 PM
        
        message = <- messageChannel
        fmt.Println("Message recieved:", message)
        // Message recieved: 
        
        // The last read will return an empty string because there are no 
        // new messages

    */
   
    // NOTE: How to read from a channel using a loop
    // This is a more efficient way to read from the channel
    // When using range, it will automatically stop reading from the channel
    // when it runs out of messages

    /*
        for message := range messageChannel {
            fmt.Println("Message recieved:", message)
        }
    */

    // NOTE: How to read from a channel using for

    // TODO: Need an explanation of this syntax

    for {
        message, ok := <- messageChannel
        if !ok {
            break
        } 
        fmt.Println("Message recieved:", message)
    }

    fmt.Println("There are no more messages to read")

}

// func fetchResource(n int) string {
//     // This will delay the output of the next line by 2 seconds
//     time.Sleep(time.Second * 2)
//     return fmt.Sprintf("result %d", n)
// }
