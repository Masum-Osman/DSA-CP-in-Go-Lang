package main
import "fmt"

type node2 struct {
	data int
	next *node2
}

func main()  {
	var link *node2
	var testCase int;
	var data int;

	fmt.Println("Enter test case: ")
	fmt.Scanf("%d", &testCase);

	fmt.Print("testcase: ", testCase)


	for i := 0; i < testCase; i++ {
		fmt.Scanf("%d", &data);
		fmt.Print(data, " data ")

		link = insert2(link, data)
	}

	// link = insert2(link, 10)
	// link = insert2(link, 2);
	// link = insert2(link, 1);
	// link = insert2(link, 3);
	// link = insert2(link, 4);

	printList2(link)
}

func insert2(head *node2, data int) *node2  {
	n := &node2{data : data}

	if head == nil {
		return n
	} else {
		n.next = head

		return n
	}
}

func printList2(head *node2)  {
	for head != nil {
		fmt.Print(head.data, " => ")
		head = head.next
	}
}