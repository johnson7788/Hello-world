package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Saturday?")
	fmt.Println(time.Saturday)
	today := time.Now().Weekday()
	fmt.Println(today)
	fmt.Println(today+1)
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}
