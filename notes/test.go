package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
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

	//f, err := os.Open("./doc/mkdir")
	//
	//name, err := f.Readdir(-1)
	//if err != nil {
	//	log.Println("readdir", err)
	//}
	//
	//for _, v := range name {
	//	fmt.Println(v.Name(), v.IsDir(), v.Mode())
	//}
	//
	//fmt.Println(name)

	//os.RemoveAll("./doc/mkdir")
	type text struct{ content string }

	type Params struct {
		toparty string
		msgtype string
		agentid int
		text    text
		safe    int
	}

	params := &Params{
		toparty: "da",
		msgtype: "text",
		agentid: 12,
		text: text{
			content: "message",
		},
		safe: 0,
	}
	paramsJSON, _ := json.Marshal(params)
	fmt.Println(params, reflect.TypeOf(paramsJSON))

	maps := map[string][]string{
		"a": []string{"a", "a"},
		"c": []string{"c", "c"},
	}
	fmt.Println(maps["a"])

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
