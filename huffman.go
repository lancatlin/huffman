package main

func Huffman(m map[rune]int) (result map[rune][]int) {
	var l []Node = []Node{}
	// 將各個字元資料組成 Root 陣列
	for key, value := range m {
		l = append(l, Root{value, key})
	}
	var tree Tree = update(l)
	result = make(map[rune][]int)
	read(tree, result, []int{})
	return result
}

func read(tree Tree, m map[rune][]int, a []int) {
	// 從組織後的 Tree 中讀取路徑
	for i, v := range tree.root {
		add := append(a, i)
		if root, ok := (*v).(Root); ok {
			c := root.c
			m[c] = add
		} else {
			read((*v).(Tree), m, add)
		}
	}
}

func update(m []Node) Tree {
	//核心演算法，將 Node 陣列反覆取出 get() 最小的兩個組合成樹，再放回新的陣列
	QuickSort(m)
	var tree Tree = Tree{m[0].get() + m[1].get(), [2]*Node{&m[0], &m[1]}}
	// 如果只剩兩個，代表到終點，回傳樹
	if len(m) == 2 {
		return tree
	}
	return update(append(m[2:], tree))
}
