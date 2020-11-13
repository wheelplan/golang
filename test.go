package main

import (
	"fmt"
)

func main() {
	/*	const a = 0x100
		fmt.Println(reflect.TypeOf(a), a, 16*16)

		const b = 0100
		fmt.Println(reflect.TypeOf(b), b, 8*8)

		const c = 100
		fmt.Println(reflect.TypeOf(c), c, 10*10)
		fmt.Printf("0b%b, %#o, %#x\n", c, b, a)

		fmt.Println(math.MinInt8, math.MaxInt8)*/

	var b int = 1
	var c int64 = int64(b)
	fmt.Println(c)

	x := 100
	p := (*int)(&x)
	fmt.Println(*p)

	type flags byte
	const (
		read flags = 3 << iota
		write
		exec
	)
	f := 2 | 5
	fmt.Printf("%b && %b && %b\n", read, write, exec)
	fmt.Printf("%b && %v\n", f, f)
}
