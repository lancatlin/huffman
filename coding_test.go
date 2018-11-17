package main

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestEncode(t *testing.T) {
	file, err := ioutil.ReadFile("test.txt")
	if err != nil {
		t.Error("Read file Error")
	}
	data := file
	var m map[byte][]byte = Huffman(Count(data))
	result, last := Encode(m, data)
	fmt.Printf("data length: %d; result length: %d; last: %d\n", len(file), len(result), last)
	var decode []byte = Decode(m, result, last)
	if equal(decode, data) == false {
		t.Error("Not Equal")
	}
	fmt.Print("\n")
}
