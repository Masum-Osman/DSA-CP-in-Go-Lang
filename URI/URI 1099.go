package main

import "fmt"

func main3() {
	
	var testCase int;
	var input int;
	var limit int;
	var min int;
	var max int; 


	fmt.Scanf("%d", &testCase);


	for i := 0; i < testCase; i++ {

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

		for i := min + 1; i < max; i++ {
			if i%2 != 0 {
				sum = sum+i;
			}
		}
		fmt.Println(sum);
	}
}