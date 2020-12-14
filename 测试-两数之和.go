package main

import (
	"fmt"
)

//
func main() {
	cal()
}

func cal() {
	var nums []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var target int = 7

	for {
		if len(nums) == 1 {
			return
		}

		num := target - nums[0]
		for i := 1; i < len(nums); i++ {
			if nums[i] == num {
				fmt.Println(nums[0], nums[i])
				nums = append(nums[:i], nums[i+1:]...)
				break
			}
		}
		nums = append(nums[1:])
	}
}
