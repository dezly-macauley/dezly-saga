package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type UserProfile struct {
    ID int
    Comments []string
    Likes int
    // An array of their friend's ids
    Friends []int
}

func main() {
   
    start := time.Now()

    userProfile, err :=  handleGetUserProfile(10)
    
    // If there was an error
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(userProfile)
    fmt.Printf("It took %+v long to fetch the user's profile\n",
        time.Since(start))

}

type Response struct {
    data any
    err error
}

// Why *UserProfile
func handleGetUserProfile(id int) (*UserProfile, error) {

    // NOTE: Creating a channel to manage the goroutines

    // respch means "Response channel"
    // It is a channel of the variable type "Response" (the reponse struct)
    // It is a buffered channel at 3 reponses

    var (
        respch = make(chan Response, 3)
        wg = &sync.WaitGroup{}
    )

    // NOTE: Three function are running concurrently

    go getComments(id, respch, wg)
    go getLikes(id, respch, wg)
    go getFriends(id, respch, wg)

    // Let the wait group know that three functions are running
    wg.Add(3)

    // This will block until the wait group counter is 0
    wg.Wait()

    // Then the channel will be closed
    close(respch)

    // NOTE: Range over the response channel
    // Warning when you range over a response channel you need to tell 
    // the program when to stop

    userProfile := &UserProfile{}
    for resp := range respch {

        if resp.err != nil {
            return nil, resp.err
        }

        switch msg := resp.data.(type) {
        case int:
            userProfile.Likes = msg
        case []int:
            userProfile.Friends = msg
        case []string:
            userProfile.Comments = msg
        }

    }

    return userProfile, nil

}

func getComments(id int, respch chan Response, wg *sync.WaitGroup) {
    time.Sleep(time.Millisecond * 200)

    comments := []string{
        "Is little Goblin junior gonna cry?",
        "I'm going to put some dirt in your eye.",
        "I missed the part where that's my problem.",
    }

    // Write data into the response channel
    respch <- Response {
        data: comments,
        err: nil,
    }

    // The function will let the wait group know that this specific function 
    // is has completed its task. And the wait group will be decremented 
    // by 1
    wg.Done()

}

func getLikes(id int, respch chan Response, wg *sync.WaitGroup) {
    time.Sleep(time.Millisecond * 200)
    respch <- Response {
        data: 11,
        err: nil,
    }

    wg.Done()

}

func getFriends(id int, respch chan Response, wg *sync.WaitGroup) {
    time.Sleep(time.Millisecond * 100)
    friendsIds := []int{11, 34, 854, 455}

    respch <- Response {
        data: friendsIds,
        err: nil,
    }

    wg.Done()
}
