package main

import (
	"fmt"
	"reflect"
)

// resetValue object must be the pointer of the struct (a adderessable value)
func resetValue(object any) {
	// get the reflected value of object
	reflected := reflect.ValueOf(object).Elem()

	// check if the object can set
	if reflected.CanSet() {
		// sets the zero value of the object type
		reflected.Set(reflect.Zero(reflected.Type()))
	}
}

type structure struct {
	pointeredFloatField *float32
	booleanField        bool
}

func main() {
	stringVariable := "a"
	floatVariable := float32(1.1)

	pointeredStringVariable := &stringVariable
	integerVariable := 1
	structVariable := structure{
		pointeredFloatField: &floatVariable,
		booleanField:        true,
	}
	println("----- initial value of variables -----")
	fmt.Printf("pointeredStringVariable: %#v\n", pointeredStringVariable)
	fmt.Printf("integerVariable: %#v\n", integerVariable)
	fmt.Printf("structVariable: %#v\n", structVariable)

	resetValue(&pointeredStringVariable)
	resetValue(&integerVariable)
	resetValue(&structVariable)
	println("\n----- value of variables after reset -----")
	fmt.Printf("pointeredStringVariable: %#v\n", pointeredStringVariable)
	fmt.Printf("integerVariable: %#v\n", integerVariable)
	fmt.Printf("structVariable: %#v\n", structVariable)
}
