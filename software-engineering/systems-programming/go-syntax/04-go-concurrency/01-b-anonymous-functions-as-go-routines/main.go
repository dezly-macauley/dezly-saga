package main

import "time"

func main() {

    // NOTE: This is how you create a go routine using an annoymous function

    // An annoymous function is simply a function that doesn't have a name.
    // Unlike a regular function that is declared outside of func main(),
    // and then the name is called in func main()...

    // With an annoymous function, it is declared and called 
    // within func main()
    // This annoymous functions accepts msg which is a string

    go func(msg string) {
        println(msg)
    }("Hello World")
    // THe () after the curly braces signifies the execution of the function
    // If you want to pass arguements into the annoymous function, then you
    // would simply put the arguements into the brackets at the end


    time.Sleep(time.Second)

}
