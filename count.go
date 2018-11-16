package main

func Count(data []byte) (m map[byte]int) {
	m = make(map[byte]int)
	for _, v := range data {
		if _, ok := m[v]; ok {
			m[v] += 1
		} else {
			m[v] = 1
		}
	}
	return
}
