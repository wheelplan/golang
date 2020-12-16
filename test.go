package main

import (
	"fmt"
	"strings"
)

func main() {

	pattern := "absba"
	s := "cat dog dog dog cat"

	fmt.Println(wordPattern(pattern, s))

}

func wordPattern(pattern string, s string) bool {
	m := make(map[byte]string)

	words := strings.Split(s, " ")
	if len(pattern) != len(words) {
		return false
	}

	for i := 0; i < len(words); i++ {

		if m[pattern[i]] == "" {
			m[pattern[i]] = words[i]
			for j := i - 1; j >= 0; j-- {
				if words[i] == words[j] {
					return false
				}
			}
		} else {
			if m[pattern[i]] != words[i] {
				return false
			}
		}
	}
	return true
}
