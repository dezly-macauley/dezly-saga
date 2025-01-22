package main

import "time"

func main() {
    
    // NOTE: Use the "go" keyword to create a go routine
    // So what does this do exactly?
    // In simple terms, it tells the program "Start processing this line of
    // code but don't wait for it to finish. Immediately start executing the
    // next line of code"

    go helloWorld()

    /*
        This line pauses the execution 
        of the main goroutine (the main function) for one second.
        This allows the helloWorld goroutine enough time to execute 
        and print "Hello world!" before the main function completes.
        
        Without this line the program would complete and "Hello world" would
        not be printed
    */
    time.Sleep(time.Second)

}

func helloWorld() {
    println("Hello world!")
}
