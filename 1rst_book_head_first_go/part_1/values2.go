package main

import (
	"fmt"
)

func main() {
	var originalCount int = 10
	var eatenCount int = 4

	fmt.Println("I started with", originalCount, "apples.")
	fmt.Println("Some jerk ate", eatenCount, "apples.")
	fmt.Println("There are", originalCount-eatenCount, "apples left.")

	quantity := 4
	length, width := 1.2, 2.4
	customerName := "Damon Cole"

	// quantity = 5 // error: no new variables on left side of :=

	fmt.Println(customerName)
	fmt.Println("has ordered", quantity, "sheets")
	fmt.Println("each with an area of")
	fmt.Println(length*width, "square meters")
}
