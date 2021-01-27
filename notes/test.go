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

	file.Chmod(0666 | os.ModeSticky)

	log.Println("Change after, file mode : ", getFileMode(file))

	//linkERR := os.Link(filename, "./doc/link.txt")
	//if linkERR != nil {
	//	log.Println(linkERR)
	//}

	symlink := os.Symlink(filename, "./doc/chmod-link.txt")
	if symlink != nil {
		fmt.Println(symlink)
	}
	linkPath, err := os.Readlink("./doc/chmod-link.txt")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(linkPath)

}

//

func getFileMode(file *os.File) os.FileMode {

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println(err)
	}
	return fileInfo.Mode()
}

//
