<div id="top">
<!-- HEADER STYLE: CLASSIC -->
<div align="center">

# LEARNING_GO

<em>My personal journey learning the Go programming language</em>

<!-- BADGES -->
<img src="https://img.shields.io/github/last-commit/megane18/learning_go?style=flat&logo=git&logoColor=white&color=0080ff" alt="last-commit">
<img src="https://img.shields.io/github/languages/top/megane18/learning_go?style=flat&color=0080ff" alt="repo-top-language">
<img src="https://img.shields.io/github/languages/count/megane18/learning_go?style=flat&color=0080ff" alt="repo-language-count">

<em>Built with the tools and technologies:</em>

<img src="https://img.shields.io/badge/Markdown-000000.svg?style=flat&logo=Markdown&logoColor=white" alt="Markdown">
<img src="https://img.shields.io/badge/Go-00ADD8.svg?style=flat&logo=Go&logoColor=white" alt="Go">

```
    ╔══════════════════════════════════════╗
    ║        🚀 Learning Go Journey        ║
    ║   ┌─────────────────────────────────┐ ║
    ║   │  📚 Reading docs...             │ ║
    ║   │  💻 Writing code...             │ ║
    ║   │  🔧 Building projects...        │ ║
    ║   │  🎯 Mastering concepts...       │ ║
    ║   └─────────────────────────────────┘ ║
    ╚══════════════════════════════════════╝
```

</div>
<br>

---

## About This Repository

This repository tracks my progress while learning Go. It contains my notes, code examples, and key concepts I've discovered along the way.

## What I've Learned So Far

### Building and Running Go Programs

**Building a Go Program:**
1. Make sure you're in the right directory
2. Run `go mod init hello` (replace "hello" with your desired module name)
3. Run `go build` to create an executable
4. Check the build by running `./hello.exe`

**Running a Go File:**
Use the `go run` command. Unlike `go build`, `go run` doesn't create an executable file - it builds and runs the file in one command.

Example: `go run main.go`

### Package Declaration and Imports

**Package Declaration:**
```go
package main
```
This tells the Go compiler which package this `.go` file belongs to. If the package declaration is `package main`, then this program will be compiled into an executable.

**Import Statement:**
```go
import "fmt"
```
The import keyword allows us to use code from other packages. In this example, we're importing the `fmt` package to use functions like `Println`.

### Working with Multiple Packages

**Individual Import Statements:**
```go
import "fmt"
import "math"
import "strings"
```

**Grouped Import (Recommended):**
```go
import (
    "fmt"
    "math" 
    "strings"
)
```

**Package Aliases:**
```go
import (
    f "fmt"
    m "math"
    "strings"
)
```

### Comments in Go

**Single Line Comments:**
```go
// This explains what's happening next
fmt.Println("Hello, World!")
```

**Multi-Line Comments:**
```go
/*
This is useful for explaining complex logic
or disabling multiple lines at once.
*/
```

**The Golden Rule:** "Don't comment bad code — rewrite it." Focus on writing clean, readable code first, then add comments only where they genuinely help explain the reasoning.

### Go Build vs Go Run

**When to use `go build`:**
- For distribution and sharing
- When performance matters (running the same program multiple times)
- For production/finished programs
- To check if code compiles without running

**When to use `go run`:**
- During active development
- For scripts and quick programs
- While learning and experimenting

**My Rule of Thumb:**
- Development and learning: `go run`
- Production and sharing: `go build`

### Finding Help and Resources

**Built-in Documentation:**
```bash
go doc fmt                # Get package info
go doc fmt.Println        # Get function details
```

**Pro Search Tip:** Use "Golang" instead of just "Go" when searching online for more specific results.

**Useful Resources:**
- Official Go website and documentation
- Stack Overflow (search for "Golang")
- Go community forums

---

## Getting Started

### Prerequisites
- Go installed on your system

### Running the Examples
1. Clone this repository:
   ```bash
   git clone https://github.com/megane18/learning_go
   cd learning_go
   ```

2. Initialize Go modules (if needed):
   ```bash
   go mod init learning_go
   ```

3. Run any Go file:
   ```bash
   go run main.go
   ```

4. Or build an executable:
   ```bash
   go build
   ./learning_go.exe
   ```

---

## Progress Notes

This README and repository will be updated as I continue learning Go and exploring its features. Each section represents concepts I've learned and understood well enough to explain in my own words.

---

<div align="left"><a href="#top">⬆ Return</a></div>

---

*Happy coding! 🚀*