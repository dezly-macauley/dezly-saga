package main

import "fmt"

// NOTE: Creating a custom varible type
type Color int

// NOTE: This is how you create enums in Go

const (
    // iota is simply a convinient way of numbering items, and giving them the
    // same variable type as the previous line.
    // ColourBlue is the variable name, the variable type of ColourBlue is
    // Color (which is an integer)
    // iota gives ColourBlue the value of 0
    // and the next variable declared after ColourBlue will be given the value
    // of 1, and the next variable gets the value of 2 and so on...

    // ColorBlue = 0 
    ColorBlue Color = iota

    // ColorBlack = 1
    ColorBlack


    // ColorYellow = 2
    ColorYellow
    
    // ColorPink = 3
    ColorPink

    // NOTE: However this can become hard to manage as you have to keep track
    // of what number each color is associated with
    // So a better way to do this would be to implement a String inteface for
    // the custom variable type Color

)

// NOTE: How to create an inteface
// It works like a regular struct method but you don't have to call it on the
// instance. It works automatically. String() is a special function in Go

// In some code bases this may be simplified to something like
// func (c Color) newFunction() {}
// I just prefer to be explicit with my variable names when I'm programming
func (instanceOfColor Color) String() string {

    switch instanceOfColor {
    case ColorBlue:
        return "Blue"
    case ColorBlack:
        return "Black"
    case ColorYellow:
        return "Yellow"
    case ColorPink:
        return "Pink"
    default:
        panic("Invalid color")
    }

    // NOTE: String() is a special type of function in Go
    // This function will be avaible for any variable type that is 
    // an instance of the variable type Color
    // It will make variable name of any variable that is an instance of Color,
    // a string.
    // This will allow you to do things like:
    // fmt.Println("Your favourite color is:", ColorBlack)
    // Output will be: "Your favourite color is Black"

}

func main() {
    fmt. Println("The color is:", ColorBlue)
    // The color is: Blue
}
