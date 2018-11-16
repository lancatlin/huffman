package main

import "fmt"

func Encode(m map[rune][]int, data []rune) []byte {
	var list []int = []int{}
	for _, v := range data {
		list = append(list, m[v]...)
	}
	fmt.Printf("data len: %d; list len %d\n\n", len(data), len(list))
	var charlist []byte = make([]byte, len(list)>>3+1)
	for i := 0; i < len(list); i += 8 {
		var char byte = 0
		for j := 0; j < 8; j++ {
			// if over index
			if i+j > len(list)-1 {
				break
			} else {
				char += byte(list[i+j])
				char = char << 1
			}
		}
		charlist = append(charlist, char)
	}
	return charlist
}
