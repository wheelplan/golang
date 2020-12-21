package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue // 仅用于 for 循环，终止后续逻辑，立即进入下一轮循环
		}

		if i > 5 {
			break // 用于 switch、for、select 语句，终止整个语句块执行
		}

		fmt.Println(i) // 1 3 5
	}

	// 配合标签，break 和 continue 可在多层嵌套中指定目标层级
outer:
	for x := 0; x < 5; x++ {
		for y := 0; y < 10; y++ {
			if y > 2 {
				fmt.Println()
				continue outer
			}

			if x > 2 {
				break outer
			}

			fmt.Print(x, ":", y, " ")
			// 0:0 0:1 0:2
			// 1:0 1:1 1:2
			// 2:0 2:1 2:2
		}
	}
}
