package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	名字 := "13" +
		"14"

	fmt.Println(名字, reflect.TypeOf(名字))

	x, _ := strconv.Atoi(名字)
	fmt.Println(x, reflect.TypeOf(x))

	fmt.Println("Hello, Github !")
	fmt.Println("Over Game")
	fmt.Println("I have a dream !")
}
