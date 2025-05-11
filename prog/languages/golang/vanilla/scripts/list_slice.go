package main

import (
	"fmt"
	"reflect"
)

func main() {

	var array [2]string = [...]string{"first", "second"}
	var slice_1 []string = []string{"first", "second"}
	var slice_2 []string = make([]string, 0)
	
	// array := [...]string{"first", "second"}
	// slice_1 := []string{"first", "second"}
	// slice_2 := make([]string, 0)

	// var all []any = []any{array, slice_1, slice_2}
	// all := []any{array, slice_1, slice_2}
	all := []any{
		array,
		slice_1,
		slice_2,
	}

	for key, object := range all {

		fmt.Println(
			"------------------------------\n" +
			"id:", key, "\n" +
			"typeof:", reflect.TypeOf(object),
		)

	}
	
	println("------------------------------")

}
