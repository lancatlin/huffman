package main

import (
	"math/rand"
	"testing"
    "fmt"
)

func TestSort(t *testing.T) {
    const length int = 100
	var list []Node = make([]Node, length)
	for i := 0; i < length; i++ {
		list[i].sum = rand.Intn(10000)
	}
	qsort(list)
    fmt.Println(list)
	for i := 1; i < length; i++ {
		if list[i-1].sum > list[i].sum {
			t.Error("Sort Error: The Value i is smaller then i-1")
		}
	}
}
