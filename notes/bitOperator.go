package main

import "fmt"

func main() {

	const (
		read byte = 1 << iota
		write
		exec
		freeze
	)

	a := read | write | freeze // 1011
	b := read | freeze | exec  // 1101
	c := a &^ b                // 0010

	fmt.Printf("%04b &^ %04b = %04b\n", a, b, c)
}
