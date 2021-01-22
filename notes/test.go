package main

import (
	"fmt"
	"reflect"
	"strings"
	"unicode"
)

func main() {
	//file, err := os.Open("C:\\GoProgram\\src\\goNotes\\notes\\test.go")
	//if err != nil {
	//	panic(err)
	//}
	//defer file.Close()
	//writer := bufio.NewWriter(os.Stdout)
	//writer.ReadFrom(file)
	//writer.Flush()
	//
	////var str string  = "I am string"
	//var num int32	= 10086
	//
	//a := strconv.Itoa(int(num))
	//b := int(num)
	//fmt.Println(a, b, reflect.TypeOf(a), reflect.TypeOf(b))
	//
	//fmt.Println(a == strconv.Itoa(b))
	//
	//fmt.Println(strings.Contains(a, strconv.Itoa(b)))
	//
	//var r rune = '1'
	//var c rune = '0'
	//
	//fmt.Println(strings.ContainsRune(a, r & c))
	//fmt.Println(strings.ContainsAny("in failure", "x "))
	//fmt.Println(strings.ContainsAny(" ", ""))

	fmt.Println(strings.Count("Hi ! Google china !", ""))

	a := strings.Fields("hi !")
	fmt.Println(a, reflect.TypeOf(a), len(a))
	c := "      "
	b := strings.Fields(c)
	fmt.Println(reflect.TypeOf(b), b, len(b))
	a = strings.FieldsFunc("Are you okay ?", unicode.IsSpace)
	fmt.Println(reflect.TypeOf(a), len(a), a)

	a = strings.FieldsFunc("are you ok", field)
	fmt.Println(reflect.TypeOf(a), len(a), a)

	fmt.Println(strings.Split("abc", ""))
	fmt.Println(strings.Count("abc", "a"))

	a = strings.Split("a man a plan a canal panama", "a ")
	fmt.Println(len(a), a)

	i := strings.IndexRune("abcdfeg", 'g')
	fmt.Println(i)
	fmt.Println(unicode.IsPunct('。'))
}

func field(r rune) bool {
	if r != 'o' && r != ' ' {
		return false
	}
	return true
}
