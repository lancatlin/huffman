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

func TestSliceTo64(t *testing.T) {
	slice := []byte{0, 1, 1, 0, 0, 0, 0, 1, 1, 1, 0}
	x := SliceTo64(slice)
	output := UintToSlice(x)
	t.Log(slice, output)
	for i, v := range slice {
		if v != output[i] {
			t.Error("Not Equal")
		}
	}
}
