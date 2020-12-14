/*package main

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
}*/

// 超集接口可隐式转换为子集，反过来则不行
/*package main

import "fmt"

type stringer interface {
	string() string
}

type tester interface {
	stringer
	test()
}

type data struct{}

func (data) string() string {
	return "A"
}

func (*data) test() {}

func pp(a stringer) {
	fmt.Println(a.string())
}

func main() {
	var d data
	var t1 tester = &d
	pp(t1)

	t1.test()
	fmt.Println(t1.string())

	var t2 stringer = t1
	fmt.Println(t2.string())

	// var t3 tester = t2   // stringer does not implement tester (missing test method)
}*/

// 匿名接口类型可直接用于变量定义或作为结构字段类型
package main

import "fmt"

type data struct{}

type node struct {
	data interface {
		string() string
	}
}

func main() {
	var t interface {
		string() string
	} = data{}

	n := node{
		data: t,
	}

	fmt.Println(n.data.string())
}

func (data) string() string {
	return "B"
}
