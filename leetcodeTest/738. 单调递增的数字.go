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
