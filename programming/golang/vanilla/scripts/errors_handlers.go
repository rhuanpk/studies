package main

import (
	"errors"
	"fmt"
)

var errAnyone = errors.New("any error message")

type errStruct struct {
	message string
}

func (es *errStruct) Error() string {
	return "error from errStruct: " + es.message
}

func returnError() error {
	return errAnyone
}

func returnWrappedError() error {
	return fmt.Errorf("an error occurred: %w", errAnyone)
}

func returnStructError(message string) error {
	return fmt.Errorf("an error ocurred: %w", &errStruct{message: message})
}

// go playground: https://go.dev/play/p/UERQmS21QY2
func main() {
	err := returnError()

	println("----- returnError() -----")

	print("err != nil: ")
	if err != nil {
		println(true)
	} else {
		println(false)
	}

	print("errors.Is(errAnyone, err): ")
	if errors.Is(errAnyone, err) {
		println(true)
	} else {
		println(false)
	}

	unwrap := errors.Unwrap(err)
	fmt.Printf("unwrap (%%v): %v\nunwrap (%%#v): %#v\n", unwrap, unwrap)

	// --------------------------------------------------

	err = returnWrappedError()

	println("----- returnWrappedError() -----")

	print("err != nil: ")
	if err != nil {
		println(true)
	} else {
		println(false)
	}

	print("errors.Is(errAnyone, err): ")
	if errors.Is(errAnyone, err) {
		println(true)
	} else {
		println(false)
	}

	unwrap = errors.Unwrap(err)
	fmt.Printf("unwrap (%%v): %v\nunwrap (%%#v): %#v\n", unwrap, unwrap)

	// --------------------------------------------------

	err = returnStructError("message for a whatever function")

	println("----- returnStructError() -----")

	print("err != nil: ")
	if err != nil {
		println(true)
	} else {
		println(false)
	}

	var structError *errStruct
	print("errors.As(err, &structError): ")
	if errors.As(err, &structError) {
		println(true)
	} else {
		println(false)
	}

	unwrap = errors.Unwrap(err)
	fmt.Printf("unwrap (%%v): %v\nunwrap (%%#v): %#v\n", unwrap, unwrap)
}
