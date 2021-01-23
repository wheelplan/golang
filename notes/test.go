package main

import (
	"fmt"
	"strings"
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

	//fmt.Println(strings.Count("Hi ! Google china !", ""))
	//
	//a := strings.Fields("hi !")
	//fmt.Println(a, reflect.TypeOf(a), len(a))
	//c := "      "
	//b := strings.Fields(c)
	//fmt.Println(reflect.TypeOf(b), b, len(b))
	//a = strings.FieldsFunc("Are you okay ?", unicode.IsSpace)
	//fmt.Println(reflect.TypeOf(a), len(a), a)
	//
	//a = strings.FieldsFunc("are you ok", field)
	//fmt.Println(reflect.TypeOf(a), len(a), a)
	//
	//fmt.Println(strings.Split("abc", ""))
	//fmt.Println(strings.Count("abc", "a"))
	//
	//a = strings.Split("a man a plan a canal panama", "a ")
	//fmt.Println(len(a), a)
	//
	//i := strings.IndexRune("abcdfeg", 'g')
	//fmt.Println(i)
	//fmt.Println(unicode.IsPunct('。'))
	//a := []string{"1", "2", "3", "4", "5"}
	//bb := Join(a, "-")
	//fmt.Println(bb)
	t := make([]byte, 7)
	t = []byte{'1', '2', '3'}
	fmt.Println(t[1:])
	test := copy(t[0:], "8")
	fmt.Println(t, test, '3', '8')
	fmt.Println("1" + "2")
	fmt.Println(strings.Repeat("na", 10))

	r := strings.NewReplacer("<", "^", "is", "XXX")
	fmt.Println(r.Replace("this is big!<"))

}

func field(r rune) bool {
	if r != 'o' && r != ' ' {
		return false
	}
	return true
}

/*func Join(str []string, sep string) string {
	// 特殊情况应该做处理
	if len(str) == 0 {
		return ""
	}
	if len(str) == 1 {
		return str[0]
	}
	buffer := bytes.NewBufferString(str[0])
	for _, s := range str[1:] {
		buffer.WriteString(sep)
		buffer.WriteString(s)
	}
	return buffer.String()
}*/
func Join(a []string, sep string) string {
	if len(a) == 0 {
		return ""
	}
	if len(a) == 1 {
		return a[0]
	}
	n := len(sep) * (len(a) - 1)
	for i := 0; i < len(a); i++ {
		n += len(a[i])
	}

	b := make([]byte, n)
	bp := copy(b, a[0])
	for _, s := range a[1:] {
		bp += copy(b[bp:], sep)
		bp += copy(b[bp:], s)
	}
	return string(b)
}
