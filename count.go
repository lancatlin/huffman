package main

func Count (data string) (m map[rune]int) {
    m = make(map[rune]int)
    for _, v := range []rune(data) {
        if _, ok := m[v]; ok {
            m[v] += 1
        }else{
            m[v] = 1
        }
    }
    return
}
