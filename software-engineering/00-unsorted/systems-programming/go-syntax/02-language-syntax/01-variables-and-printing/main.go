// NOTE: What are variables?

// Variables is a name given to a memory address, where a value can be stored.
// So instead of having to use memory addresses in the code, you can just
// name the memory address and when you refer to that, the program knows what
// you are refering to

//=============================================================================

package main

import "fmt"

//=============================================================================

// NOTE: How to declare a global variable
// This is a variable that is outside the "main" function

var name = "Foo"
var firstName string = "Foo"
var lastName string

/* 
    Your could also group all the global variables under the "var" keyword

    var (
        name = "Foo"
        firstName string = "Foo"
        lastName string
    )

*/

//=============================================================================

// NOTE: How to declare constants 
const VERSION int = 1
const KEY_LENGTH int = 10

/* 
    Your could also group all the global variables under the "const" keyword

    const (
        VERSION int = 1
        KEY_LENGTH = keylength int = 10
    )

*/

//=============================================================================

func main() {
    
    // NOTE: Local variables
    
    // These are variables that are declared within the "main" function

//=============================================================================

    // NOTE: Signed Integers 
    // (Numbers that have no decimal part)
    // A signed integer can be a negative or a postive number

    // The numbers represent bits. So int16 means 16 bit integer
    // int, int8, int16, int32, int64

    var intVar32 int32 = 1; 
    fmt.Println(intVar32)

    var intVar64 int64 = 1; 
    fmt.Println(intVar64)

    // TODO: Find out what is the default type if you don't specify 32 or 64
    var intVar int = 10
    fmt.Println(intVar)

    var negativeIntVar int = -6
    fmt.Println(negativeIntVar)

    // You can also use this := syntax to make Go infer the type.
    // If you declare a variable like this, it must be declared within 
    // the main function.
    totalCars := 3;  
    fmt.Println(totalCars)

    // A byte is 8 bits
    // A nibble is 4 bits

//=============================================================================

    // NOTE: Unsigned Integers

    // uint, uint8, uint16, uint32, uint64, uintptr

    var uint32Var uint32 = 78
    fmt.Println(uint32Var)

    var uint64Var uint64 = 90
    fmt.Println(uint64Var)

    // This is 1 byte
    var uint8Var uint8 = 0x2
    fmt.Println(uint8Var)

    // You can create a int without setting a value
    // Its default value will be set to 0
    var unitializedInt int 
    fmt.Println(unitializedInt)

    // TODO: What is the default value of an unitialized rune

//=============================================================================

    // NOTE: Byte

    // It is an alias for uint8
    var byteVar byte = 0x2
    fmt.Println(byteVar)

//=============================================================================

    // NOTE: Rune

    // Remember to set the value using single quotes and not double quotes 
    // Single quotes are for runes, and double quotes are for strings.
    // If you index a string you will get a rune
    // It is an alias for int32
    // A rune represents a Unicode point
    // So think of a run as one character in a string

    var runeVar rune = 'a'
    fmt.Println(runeVar)

//=============================================================================

    // NOTE: Floats (Numbers that have a decimal portion)
    // Floats are signed by default, so they can be negative or positive 
    // numbers
    // Floats are a different variable type than integers, 
    // so you can't add the two

    // float32 is 4 bytes
    var floatVar32 float32 = 0.1
    fmt.Println(floatVar32)

    // float64 is 8 bytes. In most cases you want to use float64
    var floatVar64 float64 = 0.1
    fmt.Println(floatVar64)

    // You can create a float without setting a value
    // Its default value will be set to 0
    var unitializedFloat float64
    fmt.Println(unitializedFloat)

//=============================================================================

    // NOTE: Strings

    var user_name = "Sasuke"
    fmt.Println(user_name)


    // You can create a string without setting a value
    // Its default value will be set to ""
    var unitializedString string 
    fmt.Println(unitializedString)

//=============================================================================

    // NOTE: Boolean

    var hasSuperPowers = true 
    fmt.Println(hasSuperPowers)
    
    var isOnline bool = false
    fmt.Println(isOnline)
    
    // TODO: What is the default value of an unitialized String

//=============================================================================

    // NOTE: Complex

    // complex64, complex128

    // These are used to represent "imaginary numbers"

//=============================================================================
    // NOTE: Printing Variables

    favouriteCountry := "Italy"
    yearsInCountry:= 10

    // NOTE: println()
    // This is what you use for testing. E.g. You want to quickly check what
    // a the value of a variable is
    
    // NOTE: fmt.Println()
    // This is what you use for production code
    // To use this you need to add the line to the top of your file,
    // outside of the "main" function.
    // import "fmt"

    // NOTE: fmt.Printf() 
    // This is like fmt.Println() but it has more advanced formatting
    
    // where you get manual control of how the output appears 
    // Take note that Printf does not automatically add a new line after
    // printing. You need to use `\n`
    // %v displays the value of the variable into the string output
    fmt.Printf("I have lived in %v for %v years\n",
        favouriteCountry, yearsInCountry)

    // Default value is 0
    var myInt int

    // Default value is 0.000000
    var myFloat float64

    // Default value is false
    var myBool bool

    // Default value is string
    var myString string

//=============================================================================

    // NOTE: Printf prints a formatted string to standard out

    // NOTE: How to print these using Printf
    fmt.Printf(
        "%v %f %v %q\n",
        myInt, myFloat, myBool, myString,
    )

    // NOTE: Format specifiers
    // %v uses this if you are not sure what the format should be
    // %s string
    // %d integers
    // %f floats
    // You can also specify the number of decimal places 
    // Eg. Two decimal places %.2f

    // NOTE: How to print the variable type of a variable
    // Use %T
    fmt.Printf("The variable type of myFloat is: %T\n", myFloat)
    // The variable type of myFloat is: float64

//=============================================================================

    // NOTE: fmt.Sprintf()
    // This returns the formatted string

//=============================================================================

    // NOTE: How to convert integers to floats in GO

    number1 := 78
    number2 := 2.5

    num1PlusNum2 := float64(number1) + number2

    fmt.Println(num1PlusNum2)

    // NOTE: If you convert a float64 to an int
    // There is no rounding up or down
    // The decimal part of the float is simply removed!

//=============================================================================

}
