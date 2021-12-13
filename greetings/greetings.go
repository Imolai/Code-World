package greetings

import "fmt"

// Returns a greeting for name.
func Hello(name string) string {
    // Returns a greeting hello with the name.
    return fmt.Sprintf("Hello, %v. Nice to meet you!", name)
}
