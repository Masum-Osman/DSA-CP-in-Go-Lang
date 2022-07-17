package main

import "fmt"

func main() {
	unsorted := []int{10, 6, 2, 1, 5, 8, 3, 4, 7, 9}
	sorted := mergeSort(unsorted)

	fmt.Println(sorted)
}

func mergeSort(items []int) []int {
	fmt.Println("counter-- ", items)

	if len(items) < 2 {
		fmt.Println("Inside IF?", len(items) < 2)
		return items
	}
	firstHalf := mergeSort(items[:len(items)/2])
	fmt.Println("firstHalf-- ", firstHalf)
	secondHalf := mergeSort(items[len(items)/2:])
	fmt.Println("secondHalf-- ", secondHalf)

	return mergeThemToOne(firstHalf, secondHalf)
}

func mergeThemToOne(a, b []int) []int {
	fmt.Println("Merger-- ", a, b)

	final := []int{}
	i := 0
	j := 0

	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}

	for ; i < len(a); i++ {
		final = append(final, a[i])
	}

	for ; j < len(b); j++ {
		final = append(final, b[j])
	}
	return final
}
