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

