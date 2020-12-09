package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	writeOne("./writeFileTest.go")
}

func checkFileExist(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func writeOne(s string) {

	content := "Are you okay .\n"

	if !checkFileExist(s) {
		_, err := os.Create(s)
		if err != nil {
			fmt.Println("file create failed", err)
			return
		}
	}

	f, err := os.OpenFile(s, os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("file open failed", err)
		return
	}

	n, err := io.WriteString(f, content)
	if err != nil {
		fmt.Println("write file failed", err)
		return
	}

	fmt.Println("write byte", n)
}
