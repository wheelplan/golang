package main

import "fmt"

func main() {

	a, b := 1, 2
	swap1(a, b)       // 2 1
	fmt.Println(a, b) // 1 2

	swap2(&a, &b)     // 2 1
	fmt.Println(a, b) // 2 1

}

func swap1(x, y int) {

	x, y = y, x
	fmt.Println(x, y)

}

func swap2(x, y *int) {

	*x, *y = *y, *x
	fmt.Println(*x, *y)

}
