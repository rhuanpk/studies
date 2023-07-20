package main

import (
	"fmt"
	"reflect"
)

// go playground: https://go.dev/play/p/eKC5__Egz_Z
func main() {

	structOfTypes := struct {
		first   string
		second  int
		third   float64
		fourth  bool
		fifth   [0]any
		sixth   []any
		seventh map[string]bool
		eighth  func()
		ninth   struct{}
		tenth   interface{}
		// eleventh any
		twelfth *any
	}{}

	const (
		firstCase  = "Name() .....:"
		secondCase = "String() ...:"
		thirdCase  = "Kind() .....:"
	)

	for _, structField := range reflect.VisibleFields(reflect.TypeOf(structOfTypes)) {

		typeOfString := structField.Type.String()
		message := "Difference between reflect.TypeOf(" + typeOfString + "):"

		fmt.Printf(
			"\n%s\n%s %s\n%s %s\n%s %s\n\n",
			message,
			firstCase,
			structField.Type.Name(),
			secondCase,
			typeOfString,
			thirdCase,
			structField.Type.Kind(),
		)

	}

}
