package main

import "fmt"

func main() {
	data := [3]string{"a", "b", "c"}

	for i := range data {
		fmt.Println(i, data[i]) // 0 1 2 , a b c
	}

	for _, s := range data {
		fmt.Println(s) // a b c
	}

	for range data {
	}

	// 无论普通 for 循环，还是 range 迭代，其定义的局部变量都会重复使用
	for i, s := range data {
		fmt.Println(&i, &s)
	}
	/*	0xc00000a0a8 0xc00003a250
		0xc00000a0a8 0xc00003a250
		0xc00000a0a8 0xc00003a250  */
}
