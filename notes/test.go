package main

import "fmt"

func main() {

	sli := []int{1, 2, 3, 4, 5, 6, 7}

	s := sli[1:5:7]

	fmt.Println("s: ", s)
	fmt.Println("length: ", len(s))
	fmt.Println("capacity: ", cap(s))
}
