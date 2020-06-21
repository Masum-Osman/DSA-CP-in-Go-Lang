package main

import "fmt"

func main2() {
	
	var input int;
	var limit int;
	var min int;
	var max int; 
	sum := 0;

	fmt.Scanf("%d", &input);
	fmt.Scanf("%d", &limit);

	if input < limit {
		min = input;
		max = limit;
	} else {
		min = limit;
		max = input
	}

	
	for i := min; i <= max; i++ {
		if i%13 != 0 {
			sum = sum+i;
		}
	}
	fmt.Println("Hi Man")
	fmt.Println(sum);
}