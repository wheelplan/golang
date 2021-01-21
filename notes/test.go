package main

import (
	"fmt"
	"log"
	"os"
)

func init() {
	log.SetPrefix("[ YUIYA ] ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)
}

func main() {

	log.Println("Are you okay.")
}

func print(a ...interface{}) (n int, err error) {
	return fmt.Fprintln(os.Stdout, a...)
}

//func WriteFrom(writer io.Writer, a... interface{}) ([]byte, error) {
//	n, err := writer.Write()
//	if n > 0 {
//		return p[:n], nil
//	}
//}
