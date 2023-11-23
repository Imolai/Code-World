# Python is just the language for you

As the [python.org](https://docs.python.org/3/tutorial/appetite.html) says, if you’re a professional software developer, you may have to work with several C/C++/Java/Go libraries but find the usual write/compile/test/re-compile cycle is too slow.

Python is just the language for you.

![PythonForYou](PythonForYou.jpg)

Let's see an easy example, why and how.

## Solve the same task

> Task: Write an idiomatic code that reads a text file line by line, searches a regular expression pattern "... stage ('...')..."  and prints only the string between the parenthesises, indented by the whitespaces before the "stage" string. If the line does not contain stage, then skip, only the searched and replaced strings should be printed. 

The code in Go, it's a modern compiled language:

```go
package main

import (
    "bufio"
    "fmt"
    "os"
    "regexp"
)

func main() {
    // Open the file
    file, err := os.Open("filename.txt")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    // Regular expression to find the pattern
    re := regexp.MustCompile(`(\s*)stage\s*\('([^']*)'\)`)

    // Read file line by line
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()

        // Find matches
        matches := re.FindStringSubmatch(line)
        if matches != nil {
            // matches[1] contains the whitespace
            // matches[2] contains the string inside the parentheses
            fmt.Printf("%s%s\n", matches[1], matches[2])
        }
    }

    // Handle any errors that might occur during the reading of the file with the scanner
    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "reading standard input:", err)
    }
}
```

The code in Python for the same task:

```python
import re

# Regular expression to find the pattern
pattern = r"(\s*)stage\s*\('([^']*)'\)"

# Open and read the file line by line
with open('filename.txt', 'r') as file:
    for line in file:
        if match := re.search(pattern, line):
            # Extract whitespace and the string inside the parentheses
            whitespace, extracted_string = match.groups()
            print(f"{whitespace}{extracted_string}")
```

> Let's compare both and think.

## Go way

In the Go example, the code requires **more lines** and **explicit error handling**. Go's **strong, static typing** and explicit error checking are evident here, which are beneficial for large-scale or performance-critical applications.

## Python way

The Python code is **notably shorter** and **more concise**, thanks to (also strong) **dynamic typing** and the **high abstraction level**. This makes Python an excellent choice for rapid development and scripting where development speed and ease of reading are priorities.

## Conclusion

This comparison illustrates well, that Python's ease of use and concise syntax make it a popular choice for **rapid development** and **automation**, while Go's robustness and performance are well-suited for **systems programming** and **large-scale applications**.

When the time is money, then Python can help.

