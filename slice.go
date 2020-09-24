package main

import "fmt"

func main()  {

	s := make([]string, 3)
	fmt.Println("empty: ",s)

	s[0] = "a"
    s[1] = "b"
    s[2] = "c"
    fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	l := s[0:1]
	fmt.Println("sl1:", l)

	l = s[:5]
    fmt.Println("sl2:", l)	

    l = s[2:]
	fmt.Println("sl3:", l)
	
}