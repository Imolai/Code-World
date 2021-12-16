# Go general development way

Explanation via development of Hello, World library.

## 1. Initialize library directory

```bash
$ mkdir greetings
$ cd greetings
greetings$ go mod init example.com/greetings
```

## 2. Make library code

```go
// greetings.go
package greetings

import "fmt"

// Returns a greeting for name.
func Hello(name string) string {
    // Returns a greeting hello with the name.
    return fmt.Sprintf("Hello, %v. Nice to meet you!", name)
}
```

## 3. Initialize program directory

```bash
greetings$ cd ..
$ mkdir hello
$ cd hello
hello$ go mod init example.com/hello
```

Content of go.mod:

```go
module example.com/hello

go 1.17
```

## 4. Make program code

```go
// hello.go
package main

import (
    "fmt"
    "example.com/greetings"
)

func main() {
    message := greetings.Hello("Gopher")
    fmt.Println(message)
}
```

## 5. Run manual test

```bash
hello$ go run hello.go
# hello.go:5:5: no required module provides package example.com/greetings; to add it:
#         go get example.com/greetings
```

## 6. Specify the place of library

```bash
hello$ go mod edit -replace example.com/greetings=../greetings
```

Content of go.mod:

```go
module example.com/hello

go 1.17

replace example.com/greetings => ../greetings
```

## 7. Synchronize the program module's dependencies

```bash
hello$ go mod tidy
# go: found example.com/greetings in example.com/greetings v0.0.0-00010101000000-000000000000
```

Content of go.mod:

```go
module example.com/hello

go 1.17

replace example.com/greetings => ../greetings

require example.com/greetings v0.0.0-00010101000000-000000000000
```

## 8. Run manual test again

```bash
hello$ go run hello.go
# Hello, Gopher. Nice to meet you!
```

## 9. Make unit test

```bash
hello$ cd ../greetings/
greetings$ vi greetings_test.go
```

Code:

```go
// greetings_test.go
package greetings

import (
    "testing"
    "regexp"
)

func TestHelloName(t *testing.T) {
    name := "World"
    want := regexp.MustCompile(`\b`+name+`\b`)
    msg := Hello(name)
    if !want.MatchString(msg) {
        t.Fatalf(`Hello("%v") = %q, want match for %#q, nil`, name, msg, want)
    }
}
```

## 10. Run unit test

```bash
greetings$ go test
# PASS
# ok      example.com/greetings   0.002s
```

If wrong:

```bash
# --- FAIL: TestHelloName (0.00s)
#     greetings_test.go:13: Hello("World") = "Hello. Nice to meet you!", want match for `\bWorld\b`, nil
# FAIL
# exit status 1
# FAIL    example.com/greetings   0.002s
```

## 11. Build

```bash
greetings$ cd ../hello/
hello$ go build
hello$ go list -f '{{.Target}}'    # discover the install path
# /home/user/go/bin/hello
hello$ ls
# go.mod  hello*  hello.go
hello$ ./hello
# Hello, Gopher. Nice to meet you!
```

## 12. Install

Before installation, add the Go install directory to your system's shell path:

- adding to the `PATH` environment variable, like `$ export PATH=$PATH:/path/to/your/install/directory`, or
- setting the `GOBIN` variable:

```bash
hello$ go env GOBIN
# ""
hello$ go env -w GOBIN=/home/user/.local/bin
hello$ go install
hello$ ls ~/.local/bin/hello
# /home/user/.local/bin/hello*
hello$ cd
$ hello
# Hello, Gopher. Nice to meet you!
```



Further reading: [Tutorial: Create a Go module](https://go.dev/doc/tutorial/create-module)
