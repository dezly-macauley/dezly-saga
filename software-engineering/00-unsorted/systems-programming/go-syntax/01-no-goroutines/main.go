package main

import (
	"log"
	"time"
    "fmt"
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

// Why *UserProfile
func handleGetUserProfile(id int) (*UserProfile, error) {

    comments, err := getComments(id)
    if err != nil {
        return nil, err
    }
    
    likes, err := getLikes(id)
    if err != nil {
        return nil, err
    }
    
    friends, err := getFriends(id)
    if err != nil {
        return nil, err
    }

    return &UserProfile{
        ID: id,
        Comments: comments,
        Likes:  likes,
        Friends: friends,
    }, nil
}

func getComments(id int) ([]string, error) {
    time.Sleep(time.Millisecond * 200)

    comments := []string{
        "Is little Goblin junior gonna cry?",
        "I'm going to put some dirt in your eye.",
        "I missed the part where that's my problem.",
    }

    return comments, nil
}

func getLikes(id int) (int, error) {
    time.Sleep(time.Millisecond * 200)
    return 33, nil
}

func getFriends(id int) ([]int, error) {
    time.Sleep(time.Millisecond * 100)
    return []int{11, 34, 854, 455}, nil
}
