package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	//readOne("C:\\GoProgram\\src\\tt\\goNotes\\readFile.go")
	readTwo("C:\\GoProgram\\src\\tt\\goNotes\\readFile.go")

}

func readOne(s string) {
	f, err := os.Open(s)
	if err != nil {
		panic("Open file failed.")
	}

	defer f.Close()

	var buf [128]byte
	var tmp []byte

	for {
		n, err := f.Read(buf[:])
		if err != nil && err != io.EOF {
			panic("Read file failed.")
		}

		if n == 0 {
			break
		}

		tmp = append(tmp, buf[:n]...)
	}

	fmt.Println(string(tmp))

}

func readTwo(s string) {
	f, err := ioutil.ReadFile(s)
	if err != nil {
		panic("Open file failed.")
	}
	fmt.Println(string(f))
}
