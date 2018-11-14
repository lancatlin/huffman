package main

func qsort(data []Node) {
	if len(data) <= 1 {
		return
	}
	mid := data[0].sum
	head, tail := 0, len(data)-1
	for i := 1; i <= tail; {
		if data[i].sum > mid {
			data[i], data[tail] = data[tail], data[i]
			tail--
		} else {
			data[i], data[head] = data[head], data[i]
			head++
			i++
		}
	}
	qsort(data[:head])
	qsort(data[head+1:])
}
