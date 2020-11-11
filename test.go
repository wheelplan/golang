package main

import (
	"fmt"
	"reflect"
)

type str string

func test(c str) {
	fmt.Println(c)
}

func main() {
	const a = 0x200
	fmt.Println(reflect.TypeOf(a), a, 256*3)
}
