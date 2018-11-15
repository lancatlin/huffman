package main

import (
    "testing"
    "io/ioutil"
    "fmt"
)

func TestHuffman (t *testing.T) {
    data, err := ioutil.ReadFile("test.txt")
    if err != nil {
        t.Error("File read Error")
    }
    str_data := string(data)
    var count map[rune]int = Count(str_data)
    var result map[rune][]bool = Huffman(count)
    fmt.Print("result: ")
    fmt.Println(result)
}
