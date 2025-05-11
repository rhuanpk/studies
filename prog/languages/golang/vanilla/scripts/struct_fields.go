package main

import (
	"fmt"
	"reflect"
)

func main() {

	xpto := &struct {
		foo int
		bar bool
	}{0, true}

	// print all field of a struct and your repectives types.
	for index, structField := range reflect.VisibleFields(reflect.TypeOf(*xpto)) {
		fmt.Println(index, structField.Name, structField.Type)
	}

}
