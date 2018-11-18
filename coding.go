package main

import (
	"bytes"
	"fmt"
	"time"
)

func SliceTo64(list []byte) (x uint64) {
	for i, v := range list {
		x += uint64(v) << uint64(63-i)
	}
	x += 1 << uint64(63-len(list))
	return
}

func Encode(m map[byte][]byte, data []byte) (charlist []byte) {
	var list []byte = make([]byte, 0)
	for _, v := range data {
		list = append(list, m[v]...)
	}
	charlist = listToByte(append(list, 1))
	return
}

func listToByte(list []byte) (data []byte) {
	for i := 0; i < len(list); i += 8 {
		var char byte = 0
		for j := 0; j < 8; j++ {
			// if over index
			if i+j >= len(list) {
				break
			} else {
				char += list[i+j] << byte(7-j)
			}
		}
		data = append(data, char)
	}
	return
}

func byteToList(data []byte) (list []byte) {
	list = make([]byte, 0, len(data)<<3)
	for _, v := range data {
		swap := make([]byte, 8)
		for j := 0; j < 8; j++ {
			swap[j] = (v >> byte(7-j)) & 1
		}
		list = append(list, swap...)
	}
	return
}

func Decode(m map[byte][]byte, data []byte) (result []byte) {
	s := time.Now()
	uintmap := make(map[uint64]byte)
	for key, value := range m {
		x := SliceTo64(value)
		uintmap[x] = key
		//fmt.Printf("%c\t%064b\n", key, x)
	}
	list := byteToList(data)
	head := 0
	end := bytes.LastIndexByte(list, 1)
	for i, _ := range list {
		if head == i {
			continue
		}
		x := SliceTo64(list[head:i])
		if v, ok := uintmap[x]; ok {
			result = append(result, v)
			head = i
		}
		if i == end {
			fmt.Println(list[i:])
			break
		}
	}
	//fmt.Printf("data len: %d\tlist len: %d\tresult len: %d\n", len(data), len(list)/8+1, len(result))
	fmt.Println(time.Since(s))
	return result
}
