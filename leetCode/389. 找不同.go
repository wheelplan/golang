//给定两个字符串 s 和 t，它们只包含小写字母。
//
//字符串 t 由字符串 s 随机重排，然后在随机位置添加一个字母。
//
//请找出在 t 中被添加的字母。
//
//示例 1：
//输入：s = "abcd", t = "abcde"
//输出："e"
//解释：'e' 是那个被添加的字母。
//
//示例 2：
//输入：s = "", t = "y"
//输出："y"
//
//示例 3：
//输入：s = "a", t = "aa"
//输出："a"

package main

import "fmt"

func main() {
	s := "abcaav"
	t := "aaaabcv"
	fmt.Println(string(findTheDifference(s, t)))
}

func findTheDifference(s, t string) (diff byte) {

	for i := range s {
		diff ^= s[i] ^ t[i]
	}

	return diff ^ t[len(t)-1]

}

/*func findTheDifference(s string, t string) byte {
	var sum int

	for i := 0; i < len(s); i++ {
		sum += int(t[i]) - int(s[i])
	}

	sum += int(t[len(s)])
	return byte(sum)

}*/

/*func findTheDifference(s string, t string) byte {
	maps := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		maps[s[i]]++
	}
	for i := 0; ; i++ {
		maps[t[i]]--
		if maps[t[i]] < 0 {
			return t[i]
		}
	}
}*/
