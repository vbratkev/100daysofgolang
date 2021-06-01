// package and imports ommitted
package main

import "fmt"

func paintNeeded(width float64, height float64) float64 {
	area := width * height
	return area / 10.0
}

func main() {
	var amount, total float64
	amount = paintNeeded(5.2, 3.4)
	fmt.Printf("%0.2f liters needed\n", amount)
	total += amount
	amount = paintNeeded(3.6, 3.3)
	fmt.Printf("%0.2f liters needed\n", amount)
	total += amount
	fmt.Printf("Total: %0.2f liters\n", total)
}
