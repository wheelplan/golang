package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	name := "13" +
		"14"

	fmt.Println(name, reflect.TypeOf(name))

	x, _ := strconv.Atoi(name)
	fmt.Println(x, reflect.TypeOf(x))

	fmt.Println("Hello, Github !")
	fmt.Println("Over Game")
	fmt.Println("I have a dream !")
}
