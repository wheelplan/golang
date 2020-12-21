package main

func main() {
	x := 100
	y := 200
	f := test(x, y)
	f()
}

func test(x, y int) func() {
	return func() {
		println(x + y)
	}
}
