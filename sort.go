package main

func QuickSort(data []Node) {
	if len(data) <= 1 {
		return
	}
	mid := data[0].get()
	head, tail := 0, len(data)-1
	for i := 1; i <= tail; {
		if data[i].get() > mid {
			data[i], data[tail] = data[tail], data[i]
			tail--
		} else {
			data[i], data[head] = data[head], data[i]
			head++
			i++
		}
	}
	QuickSort(data[:head])
	QuickSort(data[head+1:])
}
