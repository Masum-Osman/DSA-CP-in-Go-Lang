package recursion

func sum(x int) int {
	return sum(x + 1)
}
