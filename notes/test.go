//package main
//
//import (
//	"fmt"
//)
//
//func main() {
//	stack()
//	heap()
//
//
//}
//
//func stack() {
//	v := "stack"
//	var stack *string = &v
//	fmt.Println(stack)
//}
//
//func heap() {
//	h := "heap"
//	heap := new(string)
//	heap = &h
//	fmt.Println(heap)
//}
package main

import "fmt"

func main() {
	a := 1
	b := 2
	var c *int = &a
	*c = b
	fmt.Println(*c)

	d := new(int)
	*d = 100
	fmt.Println(*d)

	//var b map[string]int
	//b["沙河娜扎"] = 100
	//fmt.Println(b)
}
