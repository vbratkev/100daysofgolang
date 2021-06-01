// package and imports ommitted
package main

import "fmt"

func main() {
	var width, height, area float64
	width = 4.2
	height = 3.0
	area = width * height
	fmt.Printf("%.2f liters needed\n", area/10.0)

	width = 5.2
	height = 3.5
	area = width * height
	fmt.Printf("%.2f liters needed\n", area/10.0)

	fmt.Printf("%12s | %s\n", "Product", "Cost in Cents")
	fmt.Println("-----------------------------")

	fmt.Printf("%12s | %2d\n", "Stamps", 40)
	fmt.Printf("%12s | %2d\n", "Paper Clips", 5)
	fmt.Printf("%12s | %2d\n", "Tape", 99)
}
