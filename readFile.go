package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	//readOne("C:\\GoProgram\\src\\tt\\goNotes\\readFile.go")
	//readTwo("C:\\GoProgram\\src\\tt\\goNotes\\readFile.go")
	//readThree("C:\\GoProgram\\src\\tt\\goNotes\\readFile.go")
	readFour("./readFile.go")
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

func readThree(s string) {
	f, err := os.Open(s)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	r := bufio.NewReader(f)

	var buf [128]byte
	var tmp []byte

	for {
		n, err := r.Read(buf[:])
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

func readFour(s string) {
	f, err := os.Open(s)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	file, err := ioutil.ReadAll(f)
	if err != nil {
		panic("Read file failed.")
	}

	fmt.Println(string(file))
}
