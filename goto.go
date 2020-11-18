package main

import "fmt"

func main() {
	// 使用 goto 需要定义标签，标签区分大小写，且未使用的标签会引发编译错误。
	for i := 0; i < 3; i++ {
		fmt.Println(i)

		if i > 0 {
			goto exit // 注意：goto 不能跳转到其他函数，或内层代码块内。
		}
	}
exit:
	fmt.Println("EXIT")
}
