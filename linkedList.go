package main

import "fmt"

type node struct {
	data int
	next *node
}

func main6()  {
	var link *node 

	link = insert(link, 1);
	link = insert(link, 2);
	link = insert(link, 1);
	link = insert(link, 3);
	link = insert(link, 4);

	printList(link)
}

func insert(head *node, data int) *node  {
	n := &node{ data : data}

	if head == nil {
		return n
	} else {
		n.next = head
		return n
	}
}

func printList(head *node)  {
	for head != nil {
		fmt.Print(head.data,  " => ")
		head = head.next
	}
}