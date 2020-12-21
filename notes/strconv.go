package main

import (
	"fmt"
	"strconv"
)

func main() {
	a, _ := strconv.ParseInt("1100100", 2, 32)
	b, _ := strconv.ParseInt("144", 8, 32)
	c, _ := strconv.ParseInt("64", 16, 32)
	fmt.Println(a, b, c)

	fmt.Println("0b" + strconv.FormatInt(a, 2))
	fmt.Println("0" + strconv.FormatInt(b, 8))
	fmt.Println("0x" + strconv.FormatInt(c, 16))
}
