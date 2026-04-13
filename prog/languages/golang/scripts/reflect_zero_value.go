package main

import (
	"fmt"
	"reflect"
)

// go playground: https://go.dev/play/p/Dsou9HwhvTb
func main() {

	var variable string
	fmt.Printf("variable = %#v\n", variable)
	if reflect.ValueOf(variable).IsZero() {
		println("true for IsZero")
	} else {
		println("false for IsZero")
	}

	variable = "hello"
	fmt.Printf("variable = %#v\n", variable)

	if reflect.ValueOf(variable).IsZero() {
		println("true for IsZero")
	} else {
		println("false for IsZero")
	}

}
