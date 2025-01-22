package main

import "fmt"

func main() {

//=============================================================================

    // NOTE: Maps
    // In other programming languages they are called,
    // Associative arrays, hash tables, or dictionaries

    // Maps are an unordered collection of key value items

    /*
        | Key            | Value           |
        | -------------- | --------------- |
        | first_name     | Dezly           |
        | last_name      | Macauley        |
        | career         | Programmer      |
    */


    // NOTE: How to declare a map
    // The values must all be of the same type

    // [The variable type of the key] the variable type of the  value
    user := map[string]string {
        "first_name": "Dezly",
        "last_name": "Macauley",
        "career": "Programmer",
    }

    // NOTE: How to create an empty map
    emptyMap := map[string]int{}
    // Another way you could write this is
    // emptyMap := make(map[string]int)

    emptyMap["attackPower"] = 80
    emptyMap["seductionSkills"] = 20
    fmt.Printf("Map data is %+v\n", emptyMap)
    // Map data is map[attackPower:80 seductionSkills:20]

    // NOTE: How to delete a key from a map
    delete(emptyMap, "seductionSkills")
    fmt.Printf("Map data is %+v\n", emptyMap)
    // Map data is map[attackPower:80]

    // NOTE: When you print out a map using Println, 
    // They keys will be printed out in alphabetical order
    // with the corresponding value

    fmt.Println(user)
    //map[career:Programmer first_name:Dezly last_name:Macauley]

    // NOTE: How to make sure that a key exists in a map

    // Here I am checking if the "age" key exists, in the map called "user"
    age, ok := user["age"]

    if ok == true {
        fmt.Println("age is", age)
    } else {
        fmt.Println("age is not a key in the map called user")
    }

    // NOTE: How to declare a map where the values are not of the
    // same type

    dbz_character := map[string] interface{} {
        "name": "Goku",
        "power_level": 9000,
    }

    fmt.Println(dbz_character)
    // map[name:Goku power_level:9000]

    fmt.Println(dbz_character["power_level"])
    // 9000

    // NOTE: How to add a new key

    dbz_character["special_transformation"] = "Super Saiyan"
    fmt.Println(dbz_character["special_transformation"])
    // Super Saiyan
    
    // NOTE: How to go through a map

    for key, value := range dbz_character {
        fmt.Println("key is", key, "and value is", value)
    }

    // key is name and value is Goku
    // key is power_level and value is 9000
    // key is special_transformation and value is Super Saiyan

//=============================================================================

}
