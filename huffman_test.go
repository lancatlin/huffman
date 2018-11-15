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
    var result map[rune][]int = Huffman(count)
    var l []Node = []Node{}
    for key, value := range count {
        l = append(l, Root{value, key})
    }
    QuickSort(l)
    for i, _ := range l {
        if i == len(l) - 1 {
            break
        }
        a, _ := l[i].(Root)
        b, _ := l[i+1].(Root)
        al := len(result[a.c])
        bl := len(result[b.c])
        if al < bl {
            t.Errorf("Huffman code Error: a: %s b: %s al: %d bl: %s", a, b, al, bl)
        }
    }
}
