package main

import (
	"fmt"
	"reflect"
)

func main() {
	myInt := 4
	fmt.Println(reflect.TypeOf(myInt))
	fmt.Println(reflect.TypeOf(float64(myInt)))

	var length float64 = 1.2
	var width int = 3
	length = float64(width)
	fmt.Println(length)

	length2 := 3.635
	width2 := 7
	width2 = int(length2)
	fmt.Println(width2)

	var price int = 100
	fmt.Println("Price is", price, "dollars.")

	var taxRate float64 = 0.08
	var tax float64 = float64(price) * taxRate
	fmt.Println("Tax is", tax, "dollars.")

	var total float64 = float64(price) + tax
	fmt.Println("Total cost is", total, "dollars.")

	var availableFunds int = 120
	fmt.Println(availableFunds, "dollars available.")
	fmt.Println("Within budget?", total <= float64(availableFunds))
}
