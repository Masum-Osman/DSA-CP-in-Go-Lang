package main

import "fmt"

type coin struct {
	data int
	next *coin
}

func main7()  {
	var inputData *coin

	inputData = insertList(inputData, 100)	
	inputData = insertList(inputData, 131)	

	displayList(inputData)
}

func insertList(next *coin, info int) *coin {

	input := &coin { data : info }

	if next == nil {
		return input
	} else {
		input.next = next
		return input
	}
}

func displayList(head *coin)  {
	for head != nil {
		fmt.Print(head.data,  " => ")
		head = head.next
	}
}