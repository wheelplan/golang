package main

import "fmt"

type animal interface {
	move()
	eat()
}

type cat struct {
	name string
	feet int8
}

type chicken struct {
	feet int8
}

func main() {

}

