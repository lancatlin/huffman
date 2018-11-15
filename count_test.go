package main

import "testing"
import "fmt"

func TestStringCount (t *testing.T) {
    data := "Hello My name is Justin. I'm fifteen today. Please give me a hand."
    char_map := Count(data)
    fmt.Println(char_map)
}
