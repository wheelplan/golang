package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	filename := "./doc/chmod.txt"

	file, err := os.Create(filename)
	if err != nil {
		log.Fatal("create file err, ", err)
	}

	defer file.Close()

	fileMode := getFileMode(file)

	fmt.Println("file mode: ", fileMode)

	file.Chmod(0600 | os.ModeSticky)

	log.Println("Change after, file mode : ", getFileMode(file))

	linkERR := os.Rename(filename, "./doc/link.txt")
	if linkERR != nil {
		log.Println(linkERR)
	}

}

func getFileMode(file *os.File) os.FileMode {

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println(err)
	}
	return fileInfo.Mode()
}

//
