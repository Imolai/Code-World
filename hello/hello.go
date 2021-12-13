package main

import (
    "fmt"
    "example.com/greetings"
)

func main() {
    message := greetings.Hello("Gopher")
    fmt.Println(message)
}
