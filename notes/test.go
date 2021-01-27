package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	//filename := "C:\\GoProgram\\src\\goNotes\\doc\\chmod.txt"
	//
	//file, err := os.Create(filename)
	//if err != nil {
	//	log.Fatal("create file err, ", err)
	//}
	//
	//defer file.Close()
	//
	//fileMode := getFileMode(file)
	//
	//fmt.Println("file mode: ", fileMode)
	//
	//file.Chmod(0666 | os.ModeSticky)
	//
	//log.Println("Change after, file mode : ", getFileMode(file))

	//linkERR := os.Link(filename, "./doc/link.txt")
	//if linkERR != nil {
	//	log.Println(linkERR)
	//}

	//symlink := os.Symlink(filename, "C:\\GoProgram\\src\\goNotes\\doc\\chmod-link.txt")
	//if symlink != nil {
	//	fmt.Println(symlink)
	//}
	//linkPath, err := os.Readlink("C:\\GoProgram\\src\\goNotes\\doc\\chmod-link.txt")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//fmt.Println(linkPath)

	//mkdirerr := os.Mkdir("./doc/mkdir", 0660)
	//if mkdirerr != nil {
	//	log.Println(mkdirerr)
	//}

	//fileInfo, err := os.Lstat("./doc/mkdir")
	//fmt.Println(fileInfo.Name(), fileInfo.Mode(), fileInfo.IsDir())

	//os.RemoveAll("./doc/mkdir")

	f, err := os.Open("./doc/mkdir")

	name, err := f.Readdir(-1)
	if err != nil {
		log.Println("readdir", err)
	}

	for _, v := range name {
		fmt.Println(v.Name(), v.IsDir(), v.Mode())
	}

	fmt.Println(name)
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
