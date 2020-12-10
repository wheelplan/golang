package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	writeOne("./writeFileTest.go")
	//writeTwo("./writeFileTest.go")
	//writeThree("./writeFileTest.go")
	//writeFour("./writeFileTest.go")
}

// io.WriteString
func writeOne(s string) {

	content := "Are you okay 1.\n"

	f, err := os.OpenFile(s, os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("file open failed", err)
		return
	}

	defer f.Close()

	n, err := io.WriteString(f, content)
	if err != nil {
		fmt.Println("write file failed", err)
		return
	}

	fmt.Println("write byte", n)
}

// ioutil.WriteFile
func writeTwo(s string) {
	content := "Are you okay 2.\n"
	var buf = []byte(content)

	err := ioutil.WriteFile(s, buf, 0666)
	if err != nil {
		fmt.Println("write file failed", err)
		return
	}

	fmt.Println("write success .")
}

// File(Write,WriteString)
func writeThree(s string) {
	content := "Are you okay 3.\n"
	var buf = []byte(content)

	var f *os.File

	f, err := os.OpenFile(s, os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("file open failed", err)
		return
	}

	defer f.Close()

	n1, err1 := f.Write(buf)
	if err1 != nil {
		fmt.Println("write err", err1)
	}
	fmt.Println("write byte", n1)

	n2, err2 := f.WriteString(content)
	if err2 != nil {
		fmt.Println("write err", err2)
	}
	fmt.Println("write byte", n2)
}

// bufio.NewWriter
func writeFour(s string) {
	content := "Are you okay 4.\n"

	f, err := os.Create(s)
	if err != nil {
		fmt.Println("file create failed", err)
		return
	}

	defer f.Close()
	w := bufio.NewWriter(f)
	n, err := w.WriteString(content)
	fmt.Println("write byte", n)
	w.Flush()

}
