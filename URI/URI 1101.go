package main

import "fmt"

func main4() {
	
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

		for i := min; i <= max; i++ {
			sum = sum+i;
			fmt.Printf("%v ", i, "Sum=%v", sum);
		}
		
	}
}