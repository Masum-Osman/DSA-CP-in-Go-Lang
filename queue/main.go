package main

import "fmt"

type Queue []string

func (q *Queue) Push(data string) {
	*q = append(*q, data)
	fmt.Println(*q)
}

func (q *Queue) Pop() {
	if q.IsEmpty() {
		return
	}
	*q = (*q)[1:]
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

func (q *Queue) PrintQueue() {
	for i := 0; i < len(*q); i++ {
		fmt.Println((*q)[i])
	}
}

func main() {
	var queue Queue

	queue.Push("100")
	queue.Push("200")
	queue.Push("300")

	queue.Pop()

	queue.PrintQueue()
}
