package main

import (
	"fmt"
)

func main() {
	var quantity int = 4
	var length, width float64 = 1.2, 2.4
	var customerName string = "Vasya Pupkin"

	fmt.Println(customerName)
	fmt.Println("Has ordered", quantity, "sheets")
	fmt.Println(length*width, "square meters")
}
