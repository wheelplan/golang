package main

import (
	"fmt"
)

func main() {
	x, y := 1, 2

	defer func(a int) {
		fmt.Println("x, y = ", a, y)
	}(x)

	x += 100
	y += 200
	fmt.Println(x, y)

	defer fmt.Println("a")

	fmt.Println(test())
}

//101 202
//defer: 100
//110
//a
//x, y =  1 202

func test() (z int) {
	defer func() {
		fmt.Println("defer:", z)
		z += 10
	}()

	return 100
}
