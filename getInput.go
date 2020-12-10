package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//userScan()
	userBufio()
}

func userScan() {
	var s string
	fmt.Print("please input your name: ")

	fmt.Scan(&s)
	fmt.Printf("Hi, %s, Are you okay ?", s)
}

func userBufio() {
	fmt.Print("please input your name: ")
	reader := bufio.NewReader(os.Stdin)
	s, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Hi, %s, Are you okay ?", s)
}
