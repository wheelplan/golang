package main

import "fmt"

func testa(s string, a ...int) {
	fmt.Printf("%T, %T, %v, %v\n", s, a, s, a)
}
func testb(b ...int) {
	//fmt.Printf("%T, %v\n", b, b)
	for i := range b {
		b[i] += 100
		//fmt.Println(b[i])
	}
	//fmt.Println(b)
}

func main() {
	//testa("alan", 99, 100, 119)
	//testb(101, 97, 120)
	a := []int{1, 2, 3}
	testb(a[:]...)
	fmt.Println(a)
}
