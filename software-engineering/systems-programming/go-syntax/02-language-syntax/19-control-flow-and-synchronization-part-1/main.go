package main

import (
    "fmt"
    "time"
)

type Server struct {

    // This will be used to notify the channel that I want to quit
    quitChannel chan struct{}

    messageChannel chan string

}

// This will create a new Sever
func newServer() *Server {

    return &Server {
        quitChannel: make(chan struct{}), 

        // This is a buffered channel
        messageChannel: make(chan string, 128),
    }

}

func (instanceOfServer *Server) start() {
    fmt.Println("server starting")
    instanceOfServer.loop()
}

func (instanceOfServer *Server) sendMessage(message string) {

    // This function makes it convienient to send a message to the message 
    // You would use this channel in the main function by doing the following
    // instanceOfServer.sendMessage("Write your message here")  
    instanceOfServer.messageChannel <- message
}

func (instanceOfServer *Server) loop() {
    for {
       
        select {
        case <- instanceOfServer.quitChannel:
            // do some stuff when we need to quit
        case message := <- instanceOfServer.messageChannel:
            // This will call the handleMessage function when a message is 
            // recieved
            instanceOfServer.handleMessage(message)
        default:

        }

    }
}

func (instanceOfServer *Server) handleMessage(message string) {
    fmt.Println("Message recieved:", message)
}

func main() {

    server := newServer()

    // The `go` keyword creates a go routine. 
    // The server will run in the background and the rest of the program 
    // can continue.
    // Without this the `go` keyword the rest of the program will have to 
    // wait until the server has completed
    go server.start()

    server.sendMessage("It's Tsunade again")

    // This line is just there so that you can see the output before the 
    // program exits
    // It delayes the ending of the program by 3 seconds
    time.Sleep(time.Second * 3)

}
