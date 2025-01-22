// NOTE: If the file is not in the root folder then the file will start with
// package name-of-folder

package types

// The case of variable names in a package matters
// If it starts with an a Uppercase letter it is public access . 
// So the User struct and its fields Username and Age are public acces and 
// can be accessed outside of this file

// If it starts with lowercase it is private access. 
// So if it was `username` and `age` in lowercase, then those fields would be
// private access. 
// This means that you wouldn't be able to access them from another file
// outside of this like main.go

type User struct {
    Username string
    Age int
}
