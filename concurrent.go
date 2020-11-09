package main

import (
	"fmt"
	"time"
)

func main() {
	go task("a")
	go task("b")

	time.Sleep(time.Second * 1)
}

func task(s string) {
	for i := 0; i < 5; i++ {
		fmt.Printf("%s: %d\n", s, i)
		time.Sleep(time.Microsecond * 10)
	}
}
