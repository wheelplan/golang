package main

import "fmt"

func main() {
	var age byte = 'A'
	var len uint8 = 11
	var height rune = '\xa1'
	fmt.Printf("%T, %v\n", age, age)
	fmt.Printf("%T, %v\n", len, len)
	fmt.Printf("%T, %v\n", height, height)
}
