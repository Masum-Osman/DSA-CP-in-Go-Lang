package main

import "fmt"

func main7() {
	var i int = 42;
	// j := 2702;

	p := &i;
	fmt.Println(p)
	fmt.Println(*p)

	*p = 11;
	fmt.Println(p)
	fmt.Println(*p)
	fmt.Println(i)
	fmt.Println(&i)
}