# Documenting Go source code - godoc

Go has a builtin tool to produce source documentation based on the comments in the source code.
This tool is the [`godoc`](https://go.dev/cmd/godoc/), or named shortly **Doc**, it can be called via `go doc`.

Doc prints the documentation comments associated with the item identified by its arguments (a package, const, func, type, var, method, or struct field) followed by a one-line summary of each of the first-level items "under" that item (package-level declarations for a package, methods for a type, etc.).

Godoc is conceptually related to Python’s **Docstring** and Java’s **Javadoc** but its design is simpler. The comments read by `godoc` are not language constructs (as with Docstring) nor must they have their own machine-readable syntax (as with Javadoc). **Godoc comments are just good comments, the sort you would want to read even if godoc didn’t exist.**

To see a working example, let's create a simple own package and document it, as godoc waits for.

## Godoc example

```go
// Package mypkg provides possibility to see Go documentation.
package mypkg

import (
    "fmt"
    "os"
)

// Version represents the version of the package.
const Version = 1.0

// Help returns nothing, prints the usage as side effect and exits by exit code.
func Help(exitCode int) {
    fmt.Println("My own package.")
    fmt.Printf("Version: %.1f\n", Version)
    os.Exit(exitCode)
}
```

> Notes: Package comments should be written directly before a package clause (package x) and begin with Package x ... A documentation comment should be a complete sentence about what the thing does and contains, that starts with the name of the thing being described and ends with a period. Documentation comments should precede packages as well as exported identifiers. A code comment should be a complete sentence about why something is done, helping others and future you to understand the reason behind a particular piece of code. Exported package-level variable should have comment. A function comment should be written directly before the function declaration.

## Using our package

```go
package main

import (
    "fmt"
    "example.com/mypkg"
)

func main() {
    fmt.Println("Run package mypkg.")
    mypkg.Help(0)
}
```

Output:

```text
$ go run main.go
Run package mypkg.
My own package.
Version: 1.0
```

## Documenting the source

```text
$ go doc -all mypkg.go
package mypkg // import "example.com/mypkg"

Package mypkg provides possibility to see Go documentation.

CONSTANTS

const Version = 1.0
    Version represents the version of the package.


FUNCTIONS

func Help(exitCode int)
    Help returns nothing, prints the usage as side effect and exits by exit
    code.

```
