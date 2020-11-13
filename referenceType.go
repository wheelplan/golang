package main

import "fmt"

func main() {
	s := mkslice()
	m := mkmap()

	fmt.Println(s[0], m["a"])
}

func mkslice() []int {
	s := make([]int, 0, 10)
	s = append(s, 100)
	return s
}

func mkmap() map[string]int {
	m := make(map[string]int, 10)
	m["a"] = 1
	return m
}
