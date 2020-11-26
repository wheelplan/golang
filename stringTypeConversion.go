/*package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	s := "hi , are you okay"
	pp("s: %x\n", &s)

	bs := []byte(s)
	sb := string(bs)

	pp("string to []byte, bs: %x\n", &bs)
	pp("[]byte to string, sb: %x\n", &sb)

	rs := []rune(s)

	sr := string(rs)

	pp("string to []rune, rs: %x\n", &rs)
	pp("[]byte to string, sr: %x\n", &sr)

}

func pp(format string, ptr interface{}) {
	p := reflect.ValueOf(ptr).Pointer()
	h := (*uintptr)(unsafe.Pointer(p))
	fmt.Printf(format, *h)
}*/
//s: 78e6a8
//string to []byte, bs: c000010380
//[]byte to string, sb: c0000103a0
//string to []rune, rs: c000014140
//[]byte to string, sr: c0000103e0

//package main
//
//import (
//	"fmt"
//	"unsafe"
//)
//
//func main() {
//	bs := []byte("hi , are you okay")
//	s  := toString(bs)
//
//	fmt.Printf("bs: %x\n", &bs)
//	fmt.Printf("s : %x\n", &s)
//}
//
//func toString(bs []byte) string {
//	return *(*string)(unsafe.Pointer(&bs))
//}

package main

import "fmt"

func main() {
	var bs []byte
	bs = append(bs, "ABC"...) // append 函数将 string 追加到 []byte 内
	fmt.Println(bs)

	// 将 []byte 转换为 string key
	m := map[string]int{
		"abc": 123,
	}

	key := []byte("abc")
	x, ok := m[string(key)]
	fmt.Println(x, ok)
}
