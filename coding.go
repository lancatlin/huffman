package main

import "fmt"
import "time"

type Rootl struct {
	c    byte
	path []byte
}

func (r Rootl) get() int {
	return len(r.path)
}

func SliceTo64(list []byte) (x uint64) {
	for i, v := range list {
		x += uint64(uint64(v) << uint64(63-i))
	}
	x += 1 << uint64(63-len(list))
	return
}

func Encode(m map[byte][]byte, data []byte) (charlist []byte, last int) {
	var list []byte = make([]byte, 0)
	for _, v := range data {
		list = append(list, m[v]...)
	}
	last = len(list) % 8
	charlist = make([]byte, 0, (len(list)>>3)+1)
	for i := 0; i < len(list); i += 8 {
		var char byte = 0
		for j := 0; j < 8; j++ {
			// if over index
			if i+j >= len(list) {
				char <<= 1
			} else {
				char <<= 1
				char += list[i+j]
			}
		}
		charlist = append(charlist, char)
	}
	//fmt.Printf("last one: %b\n", charlist[len(charlist)-1])
	return
}

func Decode(m map[byte][]byte, data []byte, last int) (result []byte) {
	s := time.Now()
	var list []byte = make([]byte, 0, len(data)<<3)
	for i, v := range data {
		var swap []byte
		if i+1 == len(data) {
			swap = make([]byte, last)
		} else {
			swap = make([]byte, 8)
		}
		for j := 0; j < len(swap); j++ {
			swap[j] = (v >> byte(7-j)) & 1
		}
		list = append(list, swap...)
	}
	uintmap := make(map[uint64]byte)
	for key, value := range m {
		x := SliceTo64(value)
		uintmap[x] = key
	}
	head := 0
	for i, _ := range list {
		if head == i {
			continue
		}
		x := SliceTo64(list[head:i])
		if v, ok := uintmap[x]; ok {
			result = append(result, v)
			head = i
		}
	}
	fmt.Println(time.Since(s))
	return
}

func equal(s1 []byte, s2 []byte, flag bool) bool {
	if (s1 == nil) != (s1 == nil) {
		return false
	}
	if len(s1) != len(s2) {
		return false
	}
	for i, _ := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}
