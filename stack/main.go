package main

import (
	"fmt"
)

type Stack []string

func (s *Stack) Pop() bool {
	// so, the pop works like slicing the last element of an arr.
	// that's how it becomes LIFO

	if s.IsEmpty() {
		return false
	}

	lastIndex := len(*s) - 1
	*s = (*s)[:lastIndex]

	return true
}

func (s *Stack) Push(data string) {
	*s = append(*s, data)
	fmt.Println(*s)
}

func (s *Stack) TopValue() string {
	// if len(*s) == 0 {
	// 	return
	// }
	index := len(*s) - 1
	return (*s)[index]
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func IsBalancedStack(input string) bool {
	var stack Stack

	for i := 0; i < len(input); i++ {
		ch := string(input[i])

		if ch == "(" || ch == "{" || ch == "[" {
			stack.Push(ch)
		} else {
			fmt.Println("ELSE = ", ch, stack.TopValue())
			if stack.IsEmpty() {
				return false
			}
			// stack.TopValue()
			if stack.TopValue() == "(" && ch == ")" {
				stack.Pop()
			}
			if stack.TopValue() == "{" && ch == "}" {
				stack.Pop()
			}
			if stack.TopValue() == "[" && ch == "]" {
				stack.Pop()
			}
		}
	}

	if stack.IsEmpty() {
		return true
	} else {
		return false
	}
}

func main() {
	// var stack Stack
	var input string
	fmt.Scanln(&input)

	fmt.Println(IsBalancedStack(input))

	// fmt.Println(stack.TopValue())

	// fmt.Println(stack)
	// stack.Pop()
	// stack.Pop()
	// fmt.Println(stack)

	// fmt.Println(stack.IsEmpty())
}
