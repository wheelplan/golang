package main

import (
	"fmt"
)

func main() {
	str := "22a115f11"
	fmt.Println(lengthOfLongestSubstring(str))
}

func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)
	var endPoint, res int

	for startPoint := 0; startPoint < len(s); startPoint++ {
		if startPoint != 0 {
			delete(m, s[startPoint-1])
		}

		for endPoint < len(s) && m[s[endPoint]] == 0 {
			m[s[endPoint]]++
			endPoint++
		}

		if res < endPoint-startPoint {
			res = endPoint - startPoint
		}
	}
	return res
}
