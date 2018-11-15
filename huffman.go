package main

import "fmt"

func Huffman (m map[rune]int) (result map[rune][]bool) {
    result = map[rune][]bool{}
    //var tree Tree
    var l []Root = []Root{}
    for key, value := range m {
        l = append(l, Root{Node{value}, key})
    }
    fmt.Println(l)
    return result
}
