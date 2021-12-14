package main

import (
    "fmt"
    "example.com/mypkg"
)

func main() {
    fmt.Println("Run package mypkg.")
    mypkg.Help(0)
}
