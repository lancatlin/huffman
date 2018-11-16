package main

func Count (data []rune) (m map[rune]int) {
    m = make(map[rune]int)
    for _, v := range data {
        if _, ok := m[v]; ok {
            m[v] += 1
        }else{
            m[v] = 1
        }
    }
    return
}
