package main

import "fmt"

/*type X int

func (x *X) inc() {
	*x++
}

func main() {
	var x X
	x.inc()
	println(x)
}*/

type users struct {
	name	string
	age		byte
}

func (u users) toString() string {
	return fmt.Sprintf("%+v", u)
}

type managers struct {
	users
	title	string
}

func main() {
	var m managers
	m.name 	= "Alan"
	m.age	= 15

	println(m.toString())
}