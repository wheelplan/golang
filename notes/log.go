package main

import (
	"log"
)

func init() {
	log.SetPrefix("[ YUIYA ] ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)
}

func main() {
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
