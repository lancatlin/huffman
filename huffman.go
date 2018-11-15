package main

import "fmt"

func Huffman (m map[rune]int) (result map[rune][]int) {
    //var tree Tree
    var l []Node = []Node{}
    for key, value := range m {
        l = append(l, Root{value, key})
    }
    var tree Tree = update(l)
    result = make(map[rune][]int)
    read(tree, result, []int{})
    return result
}

func read (tree Tree, m map[rune][]int, a []int) {
    for i, v := range tree.root {
        add := append(a, i)
        if root, ok := (*v).(Root); ok {
            c := root.c
            m[c] = add
        }else{
            read((*v).(Tree), m, add)
        }
    }
}

func update (m []Node) Tree {
    QuickSort(m)
    var tree Tree = Tree{m[0].get() + m[1].get(), [2]*Node{&m[0], &m[1]}}
    if (len(m) == 2) {
        return tree
    }
    return update(append(m[2:], tree))
}
