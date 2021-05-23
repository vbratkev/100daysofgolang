package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

func main() {
	fmt.Println(math.Floor(2.75))
	fmt.Println(strings.Title("head first go"))
	fmt.Println(reflect.TypeOf(math.Floor(3.14)))
	fmt.Println(reflect.TypeOf(54))
	fmt.Println(reflect.TypeOf(3.6))
	fmt.Println(reflect.TypeOf(true))
	fmt.Println(reflect.TypeOf("What type is this"))
}
