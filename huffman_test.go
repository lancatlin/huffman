package main

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestHuffman(t *testing.T) {
	data, err := ioutil.ReadFile("test.txt")
	if err != nil {
		t.Error("File read Error")
	}
	str_data := []rune(string(data))
	var count map[rune]int = Count(str_data)
	var result map[rune][]int = Huffman(count)
	var l []Node = []Node{}
	for key, value := range count {
		l = append(l, Root{value, key})
	}
	QuickSort(l)
	for i, _ := range l {
		if i == len(l)-1 {
			break
		}
		a, _ := l[i].(Root)
		b, _ := l[i+1].(Root)
		al := len(result[a.c])
		bl := len(result[b.c])
		if al-bl > 2 {
			t.Errorf("Huffman code Error: a: {%c: %d} b: {%c: %d} al: %d bl: %d\n", a.c, a.get(), b.c, b.get(), al, bl)
		}
		fmt.Printf("'%c'\t%d\t%d\t%v\n", a.c, a.get(), al, result[a.c])
	}
}
