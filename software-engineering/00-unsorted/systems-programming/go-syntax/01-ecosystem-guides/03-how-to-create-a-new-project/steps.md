Step 1: Create a project folder and cd into that folder
Step 2: Create main.go
Step 3: Run the command: go mod init name-of-your-project-folder
Step 4: Create any additional files
Step 5: Run the command: go buid -o what-you-want-your-program-to-be-called
Step 6: Run the binary executable: ./my_program
Step 7: Create a Makefile in the root directory of the project
`touch Makefile`

## What goes inside the Makefile

```make
build:
	@go build -o my_program

run: build 
	@./my_program

# NOTE: The indentation matters

# To use this file: make run
# First the build command is defined
# Then `run` will execute the `build` command that you defined
# and then it will run the line @./my_program
# The @ sign before each command is how you show Neovim that you do not want
# to see these commands typed in the terminal. 
# You only want to see the results that happen after they are run
```

## Declaring files with the "package keyword"
// If the file is in the root folder then it is simply `package main`
// NOTE: If the file is not in the root folder then the file will start with
// package name-of-folder
// Then in main.go import it by adding this line
// import "03-how-to-create-a-new-project/types"

## Remember that the case of the functions, variables and structs matters

```go

package util

// NOTE: Remember that the function name must start with a capital letter,
// in order for it to be accessed outside of this file

// To use this function outside of this file:
// 1. Add this line to the file:
// import "03-how-to-create-a-new-project/util"
// Then you can use GetNumber in that file like this
// x := util.GetNumber()

func GetNumber() int {
    return 36
}

func GetUserName() string {
    return "James"
}


```
