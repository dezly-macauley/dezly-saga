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

mainloop:
    for {
       
        select {
        case <- instanceOfServer.quitChannel:
            fmt.Println("The server is shutting down")
           
            // This will break out of the mainloop
            // break on its own will not break out of the for loop.
            break mainloop
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

func (instanceOfServer *Server) quit() {

    // NOTE: These two lines should do the same thing

    // close(instanceOfServer.quitChannel)
    instanceOfServer.quitChannel <- struct{}{}
}

func main() {

    server := newServer()

    // NOTE: `go func` is a function that spins up another go routine

    go func() {
        time.Sleep(time.Second * 4)
        server.quit()
    } ()

    server.start()


}
