package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("./test.go")
	if err != nil {
		fmt.Println("Open file failed", err)
		return
	}

	defer f.Close()

	var tmp []byte
	var buf [128]byte

	for {
		n, err := f.Read(buf[:])
		if err != nil && err != io.EOF {
			fmt.Println("read file failed", err)
			return
		}

		if n == 0 {
			break
		}

		tmp = append(tmp, buf[:n]...)
	}
	fmt.Println(string(tmp))

}
