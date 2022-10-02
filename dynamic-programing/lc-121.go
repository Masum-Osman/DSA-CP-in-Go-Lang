package dynamicprograming

import "fmt"

func maxProfit(prices []int) int {
	var MAX int = 20000
	var min_in_arr int

	for i := 0; i < len(prices); i++ {
		if prices[i] < MAX {
			fmt.Println(min_in_arr, prices[i], MAX)
			min_in_arr = prices[i]
			MAX = prices[i]
		}
	}
	fmt.Println(min_in_arr)
	return min_in_arr
}
