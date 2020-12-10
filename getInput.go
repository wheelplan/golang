package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	userScan()
	//userScanln()
	//userBufio()
}

func userScan() {
	var name string
	fmt.Print("please input your name: ")

	fmt.Scan(&name)
	fmt.Printf("Hi, %s, Are you okay ?", name)
}

func userScanln() {
	var name string
	fmt.Print("please input your name: ")

	fmt.Scanln(&name)
	fmt.Printf("Hi, %s, Are you okay ?", name)
}

func userBufio() {
	fmt.Print("please input your name: ")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Hi, %s, Are you okay ?", name)
}
