package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestEncode(t *testing.T) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		t.Error("Read file Error")
	}
	data := file
	var m map[byte][]byte = Huffman(Count(data))
	result := Encode(m, data)
	lf, lr := float64(len(file)), float64(len(result))
	fmt.Printf("data length: %.0f; result length: %.f; \t%.2f%%\n", lf, lr, (100 * lr / lf))
	var decode []byte = Decode(m, result)
	//fmt.Printf("'%s'\n", decode)
	fmt.Printf("decode '%d'\tdata '%d'\n", len(decode), len(data))
	if !bytes.Equal(decode, file) {
		t.Error("Not Equal")
	}
	fmt.Print("\n")
}
