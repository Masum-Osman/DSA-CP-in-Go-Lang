package main

import "fmt"

func main()  {
	i := 1

	for i < 11 {
		fmt.Println(i)
		i++
	}

	for j := 0; j < 12; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("Loop")
		break
	}

	for k := 0; k < 20; k++ {
		if k %2 == 0 {
			continue
		} else {
			fmt.Println(k)
		}
	}
}