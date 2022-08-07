package main

import "fmt"

func callToSort() {
	unsorted := []int{10, 6, 2, 1, 5, 8, 3, 4, 7, 9}
	sorted := devideAndConquer(unsorted)

	fmt.Println(sorted)
}

func devideAndConquer(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	return sortAndMerge(devideAndConquer(arr[:len(arr)/2]), devideAndConquer(arr))
}

func sortAndMerge(a, b []int) []int {
	sorted := [10]

	return sorted
}
