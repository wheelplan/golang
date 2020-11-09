package main

import "fmt"

type User struct {
	name 	string
	age 	byte
}

type Manager struct {
	User
	title	string
}

func (u User) print() {
	fmt.Printf("%+v", u)
}

func (m Manager) print() {
	fmt.Printf("%+v", m)
}

type printer interface {
	print()
}

func main() {
	var u User
	u.name 	= "Yui"
	u.age	= 16

	var p printer = &u
	p.print()

	var m Manager
	m.age = 17
	m.title = "I Have A Dream !"

	var q printer = &m
	q.print()
}
