// 给定一个非负整数 N，找出小于或等于 N 的最大的整数，同时这个整数需要满足其各个位数上的数字是单调递增。
// 
// （当且仅当每个相邻位数上的数字 x 和 y 满足 x <= y 时，我们称这个整数是单调递增的。）
// 
// 示例 1:
// 
// 输入: N = 10
// 输出: 9
// 示例 2:
// 
// 输入: N = 1234
// 输出: 1234
// 示例 3:
// 
// 输入: N = 332
// 输出: 299
// 说明: N 是在 [0, 10^9] 范围内的一个整数。

package main

import (
	"fmt"
	"strconv"
)

func main() {
	res := monotoneIncreasingDigits(12313134)
	fmt.Println(res)
}

func monotoneIncreasingDigits(n int) int {
	i := 1
	s := []byte(strconv.Itoa(n))

	for i < len(s) && s[i] >= s[i-1] {
		i++
	}

	if i < len(s) {
		for i > 0 && s[i] < s[i-1] {
			s[i-1]--
			i--
		}

		for i++; i < len(s); i++ {
			s[i] = '9'
		}
	}

	res, _ := strconv.Atoi(string(s))

	return res
}
