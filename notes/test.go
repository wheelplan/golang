package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	file, err := os.Create("./doc/chmod.txt")
	if err != nil {
		log.Fatal("create file err, ", err)
	}

	defer file.Close()

	fileMode := getFileMode(file)

	fmt.Println("file mode: ", fileMode)

	file.Chmod(fileMode | os.ModeExclusive)

	log.Println("Change after, file mode : ", getFileMode(file))

}

func getFileMode(file *os.File) os.FileMode {

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println(err)
	}
	return fileInfo.Mode()
}
