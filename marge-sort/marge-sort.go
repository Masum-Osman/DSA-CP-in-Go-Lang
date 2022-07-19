package main

import "fmt"

func main() {
	unsorted := []int{10, 6, 2, 1, 5, 8, 3, 4, 7, 9}
	sorted := devide(unsorted)

	fmt.Println(sorted)
}

func devide(items []int) []int {
	fmt.Println("devider func-- ", items)

	if len(items) < 2 {
		fmt.Println("Inside IF?", len(items) < 2)
		return items
	}

	return conquer(devide(items[:len(items)/2]), devide(items[len(items)/2:]))
}

func conquer(a, b []int) []int {
	fmt.Println("conquerer func-- ", a, b)

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
