package main

import (
	"errors"
	"fmt"
	"reflect"
)

func main() {
	err := errors.New("errors.New: Here peace of shit error")
	fmt.Println(err.Error())
	fmt.Println(reflect.TypeOf(err))

	err2 := fmt.Errorf("Errorf: %0.2f peace of shit", -2.33333)
	fmt.Println(err2.Error())
	fmt.Println(reflect.TypeOf(err2))
	fmt.Println(err2)
}
