package main

import "fmt"

func twoSum(nums []int, target int) []int {

	for i, _ := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func main() {

	var nums []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var target int = 7
	fmt.Println(twoSum(nums, target))

}
