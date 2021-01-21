package main

import (
	"log"
)

func init() {
	log.SetPrefix("[ YUIYA ] ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)
}

func main() {
	// 先使用 Print 记录日志，然后调用 panic() 函数抛出一个恐慌，
	// 这时候除非使用 recover() 函数，否则程序就会打印错误堆栈信息，然后程序终止。
	// log.Panic("panic.")

	// 记录日志后，程序退出
	log.Fatal("error.")

	log.Println("Are you okay.")
}

//const (
//	Ldate         = 1 << iota     // 日期示例： 2009/01/23
//	Ltime                         // 时间示例: 01:23:23
//	Lmicroseconds                 // 毫秒示例: 01:23:23.123123.
//	Llongfile                     // 绝对路径和行号: /a/b/c/d.go:23
//	Lshortfile                    // 文件和行号: d.go:23.
//	LUTC                          // 日期时间转为 0 时区
//	LstdFlags     = Ldate | Ltime // Go 提供的标准抬头信息
//)
