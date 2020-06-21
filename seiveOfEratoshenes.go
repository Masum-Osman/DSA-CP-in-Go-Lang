package main

import "fmt"

func main5()  {
	var n int;

	fmt.Printf("Please input value: ")
	fmt.Scanf("%d", &n);

	seiveOfEratoshenes(n)
}

func seiveOfEratoshenes(a int)  {

	s := make([]int, a) 

	_ = s;

	for i := 2; i < a; i++ {
		for j := i*i; j < a; j=j+i {
			s[j - 1] = 1
		} 
	}

	for i := 1; i < a; i++ {
		if s[i - 1] == 0 {
			fmt.Printf("%d ", i)
		}
	}
}