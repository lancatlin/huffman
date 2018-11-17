package main

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestEncode(t *testing.T) {
	Testsliceto64(t)
	file, err := ioutil.ReadFile("LOGO.png")
	if err != nil {
		t.Error("Read file Error")
	}
	data := file
	var m map[byte][]byte = Huffman(Count(data))
	result, last := Encode(m, data)
	lf, lr := float64(len(file)), float64(len(result))
	fmt.Printf("data length: %0f; result length: %0f; last: %d\t %2f%%\n", lf, lr, last, (100 * lr / lf))
	var decode []byte = Decode(m, result, last)
	if !equal(decode, file[:len(file)-1], true) {
		fmt.Printf("%s\n", decode)
		fmt.Printf("decode '%c'\tdata '%c'\n", decode[len(decode)-1], data[len(data)-1])
		t.Error("Not Equal")
	}
	fmt.Print("\n")
}

func Testsliceto64(t *testing.T) {
	m := []byte{1, 0, 0, 1, 1, 1}
	x := SliceTo64(m)
	fmt.Printf("x: %064b\tm: %v\n", x, m)
}
