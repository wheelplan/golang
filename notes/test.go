package main

import (
	"fmt"
)

func main() {
	stack()
	heap()

}

func stack() {
	v := "stack"
	var stack *string = &v
	fmt.Println(stack)
}

func heap() {
	h := "heap"
	heap := new(string)
	heap = &h
	fmt.Println(heap)
}
