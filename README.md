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

# Working with Multiple Packages in Go

As I dive deeper into Go, I'm learning that most programs need more than just one package. The Go standard library has tons of useful packages, and you'll typically need several of them in any real project.

## Ways to Import Multiple Packages

### Method 1: Individual Import Statements
You can import each package on its own line:

```go
import "fmt"
import "math"
import "strings"
```

### Method 2: Grouped Import (Recommended)
The cleaner approach is to group all your imports together using parentheses:

```go
import (
    "fmt"
    "math" 
    "strings"
)
```

This is the style most Go developers prefer because it's more organized and easier to read.

## Package Aliases - Making Life Easier

Sometimes package names are long or might conflict with other names in your code. Go lets you create shortcuts (aliases) for packages:

```go
import (
    f "fmt"
    m "math"
    "strings"
)
```

With aliases set up, instead of writing:
```go
fmt.Println("Hello")
math.Sqrt(16)
```

You can write:
```go
f.Println("Hello")
m.Sqrt(16)
```

This becomes really handy when working with packages that have long names or when you want to make your code more concise.

## Key Takeaways

- Most Go programs import multiple packages
- Group imports using parentheses for better organization  
- Use aliases to shorten long package names or avoid conflicts
- The standard library has everything you need for most common tasks

# Comments in Go - My Learning Notes

Now that I've got the hang of packages and imports, it's time to tackle comments. This is actually more important than I initially thought!

## The Golden Rule About Comments

I came across this great quote: *"Don't comment bad code — rewrite it."* This really hit home for me. The goal should always be to write code that's so clear it speaks for itself. Comments should add value, not explain messy code.

## When Should I Actually Use Comments?

I've learned that comments work best for:

- Explaining the "why" behind tricky logic (not just what it does)
- Marking sections that are particularly important or delicate
- Leaving myself reminders about what needs to be done next
- Temporarily turning off code without losing it completely
- Writing documentation that Go's tools can pick up later

## Two Types of Comments in Go

### Single Line Comments (`//`)

These are perfect for quick notes or disabling single lines:

```go
// This explains what's happening next
fmt.Println("Hello, World!")

fmt.Println("This runs") // This comment won't affect the code
// fmt.Println("This won't run because it's commented out")
```

The cool thing is you can put `//` anywhere on a line, and everything after it gets ignored.

### Multi-Line Comments (`/* */`)

When I need to comment out big chunks of code or write longer explanations:

```go
/*
This is useful when I need to explain something complex
or when I want to disable multiple lines at once.
fmt.Println("None of this code will execute")
fmt.Println("Everything between the stars is ignored")
*/
```

## My Strategy

I'm trying to focus on writing clean, readable code first. Then I add comments only where they genuinely help explain the reasoning or provide important context. Less is often more with comments!

# Finding Help and Resources for Go

As I'm learning Go, I've discovered there are tons of helpful resources out there. Here's what I've found most useful so far.

## Built-in Documentation (`go doc`)

This is probably the coolest feature I've discovered - Go comes with its own documentation tool! You can get info about any package or function right from your terminal.

### Getting Package Info
```bash
go doc fmt
```
This shows me everything about the fmt package - super handy when I forget what functions are available.

### Getting Function Details
```bash
go doc fmt.Println
```
This gives me specific details about how Println works, what parameters it takes, etc. Way faster than googling!

## Official Go Resources

The official Go website has everything you need:
- Complete language documentation
- Tutorials and getting started guides
- Standard library reference

These are always my first stop when I need authoritative information.

## Community Help

When I get stuck, these places have been lifesavers:
- **Stack Overflow** - Tons of Go questions and answers from experienced developers
- **Go community forums** - Great for asking questions and seeing what others are working on

## Pro Search Tip

Here's something I learned the hard way: when searching online, use **"Golang"** instead of just **"Go"**. "Go" is way too generic and you'll get results about everything except the programming language. "Golang" gets you exactly what you're looking for every time.

## My Learning Strategy

I try to use `go doc` first for quick reference, then check the official docs for deeper understanding, and hit up Stack Overflow when I'm really stuck on something specific.


# Go Build vs Go Run - My Learning Notes

One of the first things I wondered about when learning Go was: when should I use `go build` versus `go run`? Here's what I've figured out.

## The Big Difference

**`go run` does NOT create an executable file you can keep**
- It compiles your code behind the scenes (in memory or temporary files)
- Runs the program immediately
- Cleans up everything when it's done
- Your directory stays clean - no extra files created

**`go build` creates a permanent executable file**
- Makes a file like `hello.exe` that stays in your directory
- You can run this file anytime without recompiling
- Perfect for sharing or deploying your program

## When I Use `go build`

### For Distribution and Sharing
When I want to give my program to someone else or deploy it somewhere, I need an actual executable file. The person running it doesn't need Go installed on their machine.

### When Performance Matters
If I'm going to run the same program many times, it's faster to build once and run the executable multiple times. No recompilation overhead each time.

### For Production/Finished Programs
When my code is stable and ready for the real world, I build it into a proper executable.

### Just to Check if Code Compiles
Sometimes I want to see if my code has any compilation errors without actually running it.

## When I Use `go run`

### During Active Development
Perfect for quick testing while I'm writing code. I can make changes and test immediately without worrying about cleaning up executable files.

### For Scripts and Quick Programs
Great for small utilities or one-time programs that I don't need to keep around.

### While Learning and Experimenting
When I'm just playing around with Go concepts and trying out examples from tutorials.

## My Rule of Thumb

- **Development and learning**: `go run`
- **Production and sharing**: `go build`

## Quick Example

```bash
# This creates hello.exe in my directory (permanent file)
go build hello.go
./hello.exe

# This runs the program but leaves no files behind
go run hello.go
```

After `go run`, my directory looks exactly the same. After `go build`, I have a new executable file sitting there ready to use whenever I want.