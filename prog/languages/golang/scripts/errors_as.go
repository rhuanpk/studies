package main

import "errors"

type errStruct struct {
	message string
}

func (es errStruct) Error() string {
	return es.message
}

func function() error {
	return errStruct{"error message"}
}

// go playground: https://go.dev/play/p/1sZlZTldZAY
func main() {
	// // panics
	// var target errStruct
	// println(errors.As(function(), target))

	// not panic and binds
	var target errStruct
	println(errors.As(function(), &target))

	// // panics
	// var target *errStruct
	// println(errors.As(function(), target))

	// not panic but not binds
	var targetPtr *errStruct
	println(errors.As(function(), &targetPtr))

	// --------------------------------------------------

	// // panics
	// println(errors.As(function(), errStruct{}))

	// checks
	println(errors.As(function(), &errStruct{}))

	// checks
	println(errors.As(function(), new(errStruct)))
}
