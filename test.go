package main

import "fmt"

type tester interface {
	test()
	string() string
}

type data struct{}
type sdb string

func (d *data) test() {}

func (data) string() string {
	return "A"
}

func main() {

	var d data
	var t tester = &d

	t.test()
	fmt.Println(t.string())

	var t1, t2 interface{}
	fmt.Println(t1 == t2)

	t1, t2 = [2]int{1, 1}, [2]int{1, 1}
	fmt.Println(t1 == t2)
}
