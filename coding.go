package main

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
	head := 0
	for i, _ := range list {
		for key, value := range m {
			if equal(list[head:i], value) {
				result = append(result, key)
				head = i
			}
		}
	}
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
