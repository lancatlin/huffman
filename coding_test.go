package main

import (
    "fmt"
    "testing"
    "io/ioutil"
)

func TestEncode(t *testing.T) {
    file, err := ioutil.ReadFile("test.txt")
    if err != nil {
        t.Error("Read file Error")
    }
    data := []rune(string(file))
    var m map[rune][]int = Huffman(Count(data))
    var result []byte = Encode(m, data)
    fmt.Printf("%X\n", result)
}
