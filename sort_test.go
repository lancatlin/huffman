package main

import (
	"math/rand"
	"testing"
)

func TestSort(t *testing.T) {
    const length int = 10000
	var list []Node = make([]Node, length)
	for i := 0; i < length; i++ {
		list[i].sum = rand.Intn(1000000)
	}
	qsort(list)
	for i := 1; i < length; i++ {
		if list[i-1].sum > list[i].sum {
			t.Error("Sort Error: The Value i is smaller then i-1")
		}
	}
}
