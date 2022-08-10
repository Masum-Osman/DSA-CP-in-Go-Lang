package lc231

func isPowerOfTwo(n int) bool {
	if n == 1 {
		return true
	}
	if n%2 != 0 || n <= 0 {
		return false
	}

	return isPowerOfTwo(n / 2)

}

func isPowerOfThree(n int) bool {
	if n == 1 {
		return true
	}
	if n%3 != 0 || n <= 0 {
		return false
	} else {
		return isPowerOfThree(n / 3)
	}
}

func isPowerOfFour(n int) bool {
	if n == 1 {
		return true
	}
	if n%4 != 0 || n <= 0 {
		return false
	} else {
		return isPowerOfFour(n / 4)
	}
}
