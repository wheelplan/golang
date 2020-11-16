package main

import "fmt"

func main() {
	/*	const a = 0x100
		fmt.Println(reflect.TypeOf(a), a, 16*16)

		const b = 0100
		fmt.Println(reflect.TypeOf(b), b, 8*8)

		const c = 100
		fmt.Println(reflect.TypeOf(c), c, 10*10)
		fmt.Printf("0b%b, %#o, %#x\n", c, b, a)

		fmt.Println(math.MinInt8, math.MaxInt8)*/

	type data struct {
		x int
		y bool
	}

	var z data = data{1, true}

	fmt.Println(z)
}
