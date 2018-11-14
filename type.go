package main

type Node struct {
    sum int
}

type Tree struct {
    Node
    a *Node
    b *Node
}

type Root struct {
    Node
    c rune
}
