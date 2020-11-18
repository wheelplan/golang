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

	data1 := [3]int{10, 20, 30}

	for i, x := range data1 {
		if i == 0 {
			data1[0] += 100
			data1[1] += 200
			data1[2] += 300
		}
		fmt.Printf("x: %d, data1: %d\n", x, data1[i])
		//x: 10, data1: 110
		//x: 20, data1: 220
		//x: 30, data1: 330
	}
	for i, x := range data1[:] {
		if i == 0 {
			data1[0] += 100
			data1[1] += 200
			data1[2] += 300
		}
		fmt.Printf("x: %d, data1: %d\n", x, data1[i])
		//x: 110, data1: 210
		//x: 420, data1: 420
		//x: 630, data1: 630
	}
}
