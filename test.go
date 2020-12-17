package main

import "fmt"

func main() {
	prices := []int{1, 3, 2, 8, 4, 9}
	fee := 2
	fmt.Println(maxProfit(prices, fee))
}

func maxProfit(prices []int, fee int) int {
	m := -prices[0]
	n := 0

	for i := 1; i < len(prices); i++ {
		m, n = max(m, n-prices[i]), max(n, prices[i]+m-fee)
	}
	return n
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
