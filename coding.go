package main

import "fmt"

func Encode(m map[byte][]byte, data []byte) []byte {
	var list []byte = make([]byte, 0)
	for _, v := range data {
		list = append(list, m[v]...)
	}
	fmt.Printf("data len: %d; list len %d\n\n", len(data), len(list))
	var charlist []byte = make([]byte, 0, len(list)>>3+1)
	for i := 0; i < len(list); i += 8 {
		var char byte = 0
		for j := 0; j < 8; j++ {
			// if over index
			if i+j > len(list)-1 {
				break
			} else {
				char <<= 1
				char += list[i+j]
			}
		}
		charlist = append(charlist, char)
	}
	return charlist
}

func Decode(m map[byte][]byte, data []byte) (result []byte) {
	var list []byte = make([]byte, 0, len(data)<<3)
	for _, v := range data {
		swap := make([]byte, 8)
		for j := 0; j < 8; j++ {
			swap[7-j] = v & 1
			v >>= 1
		}
		//fmt.Printf("swap: %v\tv: %08b\n", swap, data[i])
		list = append(list, swap...)
	}
	head := 0
	for i, _ := range list {
		for key, value := range m {
			if equal(list[head:i], value) {
				result = append(result, key)
				head = i
				fmt.Printf("key: %c \tvalue: %v \n ", key, value)
			}
		}
	}
	fmt.Println()
	return
}

func equal(s1, s2 []byte) bool {
	if (s1 == nil) != (s1 == nil) {
		return false
	} else if len(s1) != len(s2) {
		return false
	}
	for i, _ := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}
