cd to the directory where you want to create your project folder
`cd github/public/boot-dot-dev`

create your project folder

`mkdir name-of-project`

E.g mkdir hello-go

Enter that folder

`cd name-of-project`

Intialize the repository

`go mod init the-github-url-of-project`

E.g. go mod init github.com/dezly-macauley/hello-go

This will create a go.mod folder in the directory

go.mod
```
module github.com/dezly-macauley/hello-go

go 1.22.4
```

It will contain the path to the module, the version of go required
by the project, and any dependencies needed by the project

This is to simplify remote downloading of packages

This project does not require any dependencies to work 

Next create a new file in the root directory of your project called
`main.go`

Write this code in your file:

For an executable binary it is always called "package main"

```
package main

import "fmt"

func main() {
    fmt.Println("Hello world")
}
```

To run this code, use the following command:
`go run main.go`

However this command is only suitable for running a single file go program.

To create an executable binary, run the following command:
`go build`

By default the binary will have the same name as your project directory

If you want to executable binary to have a seperate name, you can do this:
`go build -o my-program`

This will create an executable binary. To run this binary, use the command:
`./main.go`

If you want to execute both of those commands in one line, you can do this:
`go buid && ./hello-go`

---

# Skip this Section!!!

# NOTE: Only use this command if you are trying to create a script on your 
# local machine. 

TODO: Find out how to uninstall this program.

# How to install the Go program that you have created from the binary executable

cd to your project directory 

Then run:
`go install`

NOTE: That this command will work without go build

This will install your program on a global level on the computer

Then to run your program, you would simply do this:
`hello-go`

---

## How to create a library package in Go
## NOTE: This is in a seperate folder outside your project folder

Think of this as creating a new project

`cd github/public/boot-dot-dev`

`mkdir my-strings`

`cd my-strings`

`go mod init github.com/dezly-macauley/my-strings`

Create the following file:

`my_strings.go`

```go
// NOTE: Instead of `package main`, you name it `package name_of_your_file`

// This is not called package main because it is not meant to be an 
// executable program

// NOTE: How to check if this code compiles
// Use `go build` this will build a binary executable but it can still check
// if your code will compile
// `go run` will not work because this is not a `package main` file

// When you use `go build` the compiled package is silently saved to the local
// build cache

// NOTE: That go does not accept hyphens in the name of a package so use an
// underscore

package my_strings

// NOTE: The Function name must be uppercase must be uppercase

// In Go lowercase means that the variable or function can only be used within
// this file
// Making the name uppercase means that this Reverse function can be imported
// and used outside of this file

func Reverse(s string) string {
    result := ""

    for _, v := range s {
        result = string(v) + result
    }

    return result
}

// NOTE: Also note that this is a library so there is no package main

```

## Run this command:
`go build`

Note: That this won't produce any output, but it will let you know if there is 
an error in your code.

---

# Now close this folder and go back to the directory of the hello-go project

`cd github/public/boot-dot-dev/hello-go`

Update the main.go file:

main.go
```go
package main

import (
    "fmt"
    "github.com/dezly-macauley/my-strings"
)

func main() {
    fmt.Println(
        my_strings.Reverse("Hello world"),
    )
}
```

You will get a warning about the import. Ignore this for now

Then update the go.mod file in that folder:

```go mod
// NOTE: Go's dependency management is heavily based on git and remote URL's

module github.com/dezly-macauley/hello-go

go 1.22.4

// NOTE: This line here wouldn't be be used in production code
// This is just a way to get the my_strings package to work locally without
// pulling it from github
// In simple terms, I am telling Go:
// "My program requires this dependency: github.com/dezly-macauley/my-strings v0.0.0"
// "but you don't have to download it from github.com"
// "go one directory outside of this folder on my computer ../ 
// "and then look inside the /my-strings folder"

replace github.com/dezly-macauley/my-strings v0.0.0 => ../my-strings/

require (
    github.com/dezly-macauley/my-strings v0.0.0
)
```

Now build your code again:

`go build && ./hello_go`

If you have done everything correctly then you should get the following
output:

dlrow olleH

---

Now close this directoy

Create a new project outside of this folder

`cd github/public/boot-dot-dev`

`mkdir date-test`

`cd date-test`

create a main.go file inside the folder with the following code:

```go
package main

import (
    "fmt"
    "time"

    tinytime "github.com/wagslane/go-tinytime"
)

func main() {
    tt := tinytime.New(1585750374)

    tt = tt.Add(time.Hour * 48)
    fmt.Println(tt)
}
```

# How install a third party dependency from GitHub:

`go get github.com/wagslane/go-tinytime`

This will add the dependency to your go.mod file
It will also create a go.sum file which will bring in any dependencies
that the dependency you downloaded, needs to work

compile and run the program again

`go build && ./date-test`

---

You should see this date printed:
2020-04-03T14:12:54Z

