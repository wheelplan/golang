package main

import "fmt"

// 匿名函数
func main() {
	// 赋值给变量
	add := func(x, y int) int {
		return x + y
	}
	fmt.Println(add(1, 2))

	// 作为参数
	test(func() {
		fmt.Println("Hi !")
	})

	// 作为返回值
	adding := testing()
	fmt.Println((adding(10, 20)))

	// 普通函数和匿名函数都可作为结构体字段，或经通道传递
	testStruct()
	testChannel()
}

func test(f func()) {
	f()
}

func testing() func(int, int) int {
	return func(x int, y int) int {
		return x + y
	}
}

func testStruct() {
	type cal struct {
		mul func(x, y int) int
	}

	x := cal{
		mul: func(x, y int) int {
			return x * y
		},
	}

	fmt.Println(x.mul(3, 3))
}

func testChannel() {
	c := make(chan func(int, int) int, 2)

	c <- func(x int, y int) int {
		return x + y
	}

	fmt.Println((<-c)(11, 22))
}
