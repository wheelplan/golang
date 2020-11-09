package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["age"] = 15

	x, ok := m["length"]
	fmt.Println(x, ok)

	y, ok := m["age"]
	fmt.Println(y, ok)

	delete(m, "age")

	fmt.Println(m)
}
