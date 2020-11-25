/*package main

import (
	"log"
)

var  a int = 0

func main() {
	defer func() {
		for {
			if err := recover(); err != nil {
				log.Println(err)
			} else {
				log.Fatalln("fatal",recover())
			}
		}
	}()

	defer func() {
		panic("you are dead")
	}()

	panic("i am dead")
}*/

//2020/11/24 19:21:29 you are dead
//2020/11/24 19:21:29 fatal <nil>

/*package main

import "log"

func main() {
	defer catch()	//	recover 必须在延迟调用函数中执行才能正常工作
	defer catch()
	defer log.Println(recover())
	defer recover()

	panic("i am dead")
}

func catch() {
	log.Println("catch:", recover())
}*/
//2020/11/24 19:31:55 <nil>
//2020/11/24 19:31:55 catch: i am dead
//2020/11/24 19:31:55 catch: <nil>

package main

import "fmt"

func main() {
	test(5, 0)
}
func test(x, y int) {
	z := 0

	func() {
		defer func() {
			if recover() != nil {
				z = 0
			}
		}()
		z = x / y
	}()

	fmt.Printf("%d / %d = %d", x, y, z)
}

// 5 / 0 = 0

// 使用 runtime、debug.PrintStack 函数输出完整调用堆栈信息
// 一般不建议使用 panic，除非是不可恢复性、导致系统无法正常工作的错误
// 例如：文件系统没有操作权限，服务端口被占用，数据库未启动等情况
/*package main

import "runtime/debug"

func main() {
	defer func() {
		if err := recover(); err != nil {
			debug.PrintStack()
		}
	}()

	test()
}

func test() {
	panic("i am" +
		"dead")
}*/
