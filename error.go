package main

import (
	"fmt"
	"log"
)

func main() {
	z, err := div(3, 0)

	if err != nil {
		switch e := err.(type) {
		case DivError:
			fmt.Printf("%d / %d , %s\n", e.x, e.y, e)
		default:
			fmt.Println(e)
		}
		log.Fatalln(err)
	}

	fmt.Println(z)
}

type DivError struct {
	x, y int
}

func (DivError) Error() string {
	return "Division by zero !"
}

func div(x, y int) (int, error) {
	if y == 0 {
		return 0, DivError{x, y}
	}
	return x / y, nil
}
