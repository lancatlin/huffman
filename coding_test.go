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
	var result []byte = Encode(m, data)
	fmt.Printf("data length: %d; result length: %d\n", len(file), len(result))
	var decode []byte = Decode(m, result)
	fmt.Printf("decode: %s\n", string(decode))
	fmt.Print("\n")
}
