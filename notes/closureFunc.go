/*package main

import "fmt"

func main() {
	for _, f := range closure() {
		f()
	}
}

func closure() []func() {
	var s []func()

	for i := 0; i < 3; i++ {
		s = append(s, func() {
			fmt.Println(&i, i)
		})
	}
	return s
}*/
//0xc0000a2058 3
//0xc0000a2058 3
//0xc0000a2058 3

/*package main

func main() {
	for _, f := range closure() {
		f()
	}
}

func closure() []func() {
	var s []func()

	for i := 0; i < 3; i++ {
		x := i
		s = append(s, func() {
			fmt.Println(&x, x)
		})
	}
	return s
}*/
//0xc0000a2058 0
//0xc0000a2070 1
//0xc0000a2078 2

package main

import "fmt"

func main() {
	a, b := closure(100)
	a() // 100
	b() // 110	?
}

func closure(x int) (func(), func()) {
	return func() {
			fmt.Println(x)
			x += 10
		}, func() {
			fmt.Println(x)
		}
}
