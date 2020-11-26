package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	s := "I" +
		"Aa"

	fmt.Println(s) // IAa

	fmt.Println(s == "IAa") // true
	fmt.Println(s < "IAb")  // true
	fmt.Println(s[2])       // 97
	c := "灿" +
		"CAN"
	c1 := c[:3]
	c2 := c[3:4]
	c3 := c[4:5]
	c4 := c[5:]
	fmt.Println(c1, c2, c3, c4) // 灿 C A N
	fmt.Printf("%#v\n", (*reflect.StringHeader)(unsafe.Pointer(&c)))
	fmt.Printf("%#v\n", (*reflect.StringHeader)(unsafe.Pointer(&c1)))

	for i, s := range c { // rune
		fmt.Printf("%d:%v ", i, s) // 28799 67 65 78
	}

	fmt.Printf("\n%d \n", len(c)) // type
	for i := 0; i < len(c); i++ {
		fmt.Printf("%d:%v ", i, c[i])
	}
}

//IAa
//true
//true
//97
//灿 C A N
//&reflect.StringHeader{Data:0x46c901, Len:6}
//&reflect.StringHeader{Data:0x46c901, Len:3}
//0:28799 3:67 4:65 5:78
//6
//0:231 1:129 2:191 3:67 4:65 5:78
