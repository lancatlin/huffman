package main

type Node interface {
	get() int
}

type Tree struct {
	sum  int
	root [2]*Node
}

type Root struct {
	sum int
	c   byte
}

func (r Root) get() int {
	return r.sum
}
func (t Tree) get() int {
	return t.sum
}
