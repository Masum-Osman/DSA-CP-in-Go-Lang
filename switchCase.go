package main

import (
	"fmt"
	"time"
)


func main()  {

	var i int

	fmt.Print("Write number (1/2/3): ")

	fmt.Scanf("%d", &i)
	switch i {
	case 1:
		fmt.Println("one")
	
	case 2:
		fmt.Println("two")

	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Friday, time.Saturday:
		fmt.Println("It's the weekend")

	default 
		fmt.Println("It's the weekday")
	}
}