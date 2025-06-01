To buld and check build

first make sure we're in the right directory
second run go mod init hello
could be hello or whatever name

then do go build
then check the buiuld: ./hello.exe


**Now to run a go file**
use the go run command
unlike go build, go run doesn't create an executable file
it builds and run the file when you run that command

example: for main.go do go run main.go
# Go Learning Progress

This repository tracks my progress while learning the Go programming language.

## Building and Running Go Programs

### Building a Go Program

1. Make sure you're in the right directory
2. Run `go mod init hello` (replace "hello" with your desired module name)
3. Run `go build` to create an executable
4. Check the build by running `./hello.exe`

### Running a Go File

Use the `go run` command. Unlike `go build`, `go run` doesn't create an executable file - it builds and runs the file in one command.

**Example:**
go run main.go


## Package and Imports

### Package Declaration
package main

This line is called a package declaration and tells the Go compiler which package this `.go` file belongs to. If the package declaration is `package main`, then this program will be compiled into an executable.

### Import Statement
import "fmt"

The `import` keyword allows us to use code from other packages. In this example, we're importing the `fmt` package to use functions like `Println`. Note how the package name is enclosed with double quotes.

## Progress Notes

This README will be updated as I continue learning Go and exploring its features.
**package and imports**
packagae main: This line is called a package declaration and it tells the Go compiler to which package this ‘.go’ file belongs. If this package declaration is ‘package main’, then this program will be compiled into an executable.

import "fmt":we have an import statement, import "fmt". The import keyword allows us to use code from other packages, in this case the Println function from the fmt package. Note how the package name is enclosed with double quotes ".