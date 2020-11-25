package main

import (
	"errors"
	"fmt"
)

func main() {
	defer println("I Have A Dream" + " ! ")
	defer println("Are you okay" + " ? ")

	a, b := 10, 10
	c, err := div(a, b)

	fmt.Println(c, err)
}

func div(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}
