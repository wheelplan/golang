package main

import (
	"fmt"
	"os"
	"syscall"
	"time"
)

func main() {

	filename := "./doc/poem.txt"
	file, err := os.OpenFile(filename, os.O_RDWR, 0666)
	if err != nil {
		fmt.Println(err)
	}

	buf := []byte{'1', '2', '3', '5', ' ', 'a', 'l', 'a', 'n'}
	var tmp []byte

	file.Seek(0, os.SEEK_END)
	n, err := file.Write(buf[:])
	if err != nil {
		fmt.Println(err)
	}

	tmp = append(tmp, buf[:n]...)

	fmt.Println(string(tmp))

	fmt.Println("end")
	file.Seek(-4, os.SEEK_END)
	file.Truncate(1)
	a, err := os.Stat(filename)
	fmt.Println(a.Name())
	info := a.Sys().(*syscall.StartupInfo)
	fmt.Println(info)

	error := os.Chtimes(filename, time.Now(), time.Now())
	if err != nil {
		fmt.Println(error)
	}

}
