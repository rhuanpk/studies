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
	return &errStruct{message: message}
}

func returnWrappedStructError(message string) error {
	return fmt.Errorf("an struct error ocurred: %w", &errStruct{message: message})
}

// go playground: https://go.dev/play/p/5AKQ3fBbdoD
func main() {
	err := returnError()

	println("----- returnError() -----")

	print("err != nil: ")
	if err != nil {
		println(true)
	} else {
		println(false)
	}

	print("errors.Is(err, errAnyone): ")
	if errors.Is(err, errAnyone) {
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

	print("errors.Is(err, errAnyone): ")
	if errors.Is(err, errAnyone) {
		println(true)
	} else {
		println(false)
	}

	unwrapedError := errors.Unwrap(err)
	fmt.Printf("unwrapedError (%%v): %v\nunwrapedError (%%#v): %#v\n", unwrapedError, unwrapedError)

	println("***")

	print("unwrapedError != nil: ")
	if unwrapedError != nil {
		println(true)
	} else {
		println(false)
	}

	print("errors.Is(unwrapedError, errAnyone): ")
	if errors.Is(unwrapedError, errAnyone) {
		println(true)
	} else {
		println(false)
	}

	unwrap = errors.Unwrap(unwrapedError)
	fmt.Printf("unwrap (%%v): %v\nunwrap (%%#v): %#v\n", unwrap, unwrap)

	// --------------------------------------------------

	var structError *errStruct
	err = returnStructError("message for a whatever function")

	println("----- returnStructError() -----")

	print("err != nil: ")
	if err != nil {
		println(true)
	} else {
		println(false)
	}

	print("errors.As(err, &structError): ")
	if errors.As(err, &structError) {
		println(true)
	} else {
		println(false)
	}

	unwrap = errors.Unwrap(err)
	fmt.Printf("unwrap (%%v): %v\nunwrap (%%#v): %#v\n", unwrap, unwrap)

	// --------------------------------------------------

	err = returnWrappedStructError("message for a whatever function")

	println("----- returnWrappedStructError() -----")

	print("err != nil: ")
	if err != nil {
		println(true)
	} else {
		println(false)
	}

	print("errors.As(err, &structError): ")
	if errors.As(err, &structError) {
		println(true)
	} else {
		println(false)
	}

	unwrapedError = errors.Unwrap(err)
	fmt.Printf("unwrapedError (%%v): %v\nunwrapedError (%%#v): %#v\n", unwrapedError, unwrapedError)

	println("***")

	print("unwrapedError != nil: ")
	if unwrapedError != nil {
		println(true)
	} else {
		println(false)
	}

	print("errors.As(unwrapedError, &structError): ")
	if errors.As(unwrapedError, &structError) {
		println(true)
	} else {
		println(false)
	}

	unwrap = errors.Unwrap(unwrapedError)
	fmt.Printf("unwrap (%%v): %v\nunwrap (%%#v): %#v\n", unwrap, unwrap)
}
