package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// 从普通文件读取
	file, _ := os.Open("C:\\GoProgram\\src\\goNotes\\notes\\test.go")
	data, err := ReadFrom(file, 12)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(data))

	// 从标准输入读取
	data, err = ReadFrom(os.Stdin, 12)
	fmt.Println(string(data))

	// 从字符串读取
	data, err = ReadFrom(strings.NewReader("from string"), 12)
	fmt.Println(string(data))
}

func ReadFrom(reader io.Reader, num int) ([]byte, error) {
	p := make([]byte, num)
	n, err := reader.Read(p)
	if n > 0 {
		return p[:n], nil
	}
	return p, err
}
