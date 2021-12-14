// Package mypkg provides possibility to see Go documentation.
package mypkg
// Package comments should be written directly before a package clause (package x) and begin with Package x ...

// A documentation comment should be a complete sentence about what the thing does and contains, that starts with the
// name of the thing being described and ends with a period. Documentation comments should precede packages as well as
// exported identifiers.

// A code comment should be a complete sentence about why something is done, helping others and future you to understand
// the reason behind a particular piece of code.

import (
    "fmt"
    "os"
)

// Version represents the version of the package.
const Version = 1.0
// Exported package-level variable should have comment.

// Help returns nothing, prints the usage as side effect and exits by exit code.
func Help(exitCode int) {
    // A function comment should be written directly before the function declaration.
    fmt.Println("My own package.")
    fmt.Printf("Version: %.1f\n", Version)
    os.Exit(exitCode)
}


