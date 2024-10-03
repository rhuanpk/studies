package main

import (
	"errors"
	"fmt"
)

var (
	errAnyone = errors.New("anyone error")
	errAnytwo = errors.New("anytwo error")
)

type errStructone struct {
	message string
}

func (so *errStructone) Error() string {
	return "so.message: " + so.message
}

type errStructtwo struct {
	message string
}

func (st *errStructtwo) Error() string {
	return "st.message: " + st.message
}

func returnError() error {
	return errAnyone
}

func returnWrappedError() error {
	return fmt.Errorf("error occurred: %w", errAnyone)
}

func returnStructError(message string) error {
	return &errStructone{message}
}

func returnWrappedStructError(message string) error {
	return fmt.Errorf("error ocurred: %w", &errStructone{message})
}

func checkError(err error, errType string) {
	print(errType, " != nil: ")
	if err != nil {
		println(true)

		println("println("+errType+".Error()):", err.Error())
		fmt.Println("fmt.Println("+errType+"):", err)

		print("errors.Is(" + errType + ", errAnyone): ")
		if errors.Is(err, errAnyone) {
			println(true)
		} else {
			println(false)
		}

		print("errors.Is(" + errType + ", errAnytwo): ")
		if errors.Is(err, errAnytwo) {
			println(true)
		} else {
			println(false)
		}

		var structoneErr *errStructone
		print("errors.As(" + errType + ", &structoneErr): ")
		if errors.As(err, &structoneErr) {
			println(true)
			print("type assertion: ")
			if _, ok := err.(*errStructone); ok {
				println(true)
			} else {
				println(false)
			}
		} else {
			println(false)
		}

		var structtwoErr *errStructtwo
		print("errors.As(" + errType + ", &structtwoErr): ")
		if errors.As(err, &structtwoErr) {
			println(true)
			print("type assertion: ")
			if _, ok := err.(*errStructone); ok {
				println(true)
			} else {
				println(false)
			}
		} else {
			println(false)
		}
	} else {
		println(false)
	}
}

func logError(err error) {
	checkError(err, "err")
	unwrap := errors.Unwrap(err)
	fmt.Printf("unwrap (%%v): %v\nunwrap (%%#v): %#v\n", unwrap, unwrap)
	if unwrap != nil {
		println("***")
		checkError(unwrap, "unwrap")
	}
}

// go playground: https://go.dev/play/p/MKyUJ1qY6pB
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

	println("\n----- returnWrappedError() -----")

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

	var structoneError *errStructone
	err = returnStructError("structone error")

	println("\n----- returnStructError() -----")

	print("err != nil: ")
	if err != nil {
		println(true)
	} else {
		println(false)
	}

	print("errors.As(err, &structoneError): ")
	if errors.As(err, &structoneError) {
		println(true)
	} else {
		println(false)
	}

	unwrap = errors.Unwrap(err)
	fmt.Printf("unwrap (%%v): %v\nunwrap (%%#v): %#v\n", unwrap, unwrap)

	// --------------------------------------------------

	err = returnWrappedStructError("structone error")

	println("\n----- returnWrappedStructError() -----")

	print("err != nil: ")
	if err != nil {
		println(true)
	} else {
		println(false)
	}

	print("errors.As(err, &structoneError): ")
	if errors.As(err, &structoneError) {
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
	if errors.As(unwrapedError, &structoneError) {
		println(true)
	} else {
		println(false)
	}

	unwrap = errors.Unwrap(unwrapedError)
	fmt.Printf("unwrap (%%v): %v\nunwrap (%%#v): %#v\n", unwrap, unwrap)

	println("\n##################################################")

	err = errors.New("new error")
	println("\n----- errors.New() -----")
	logError(err)

	// --------------------------------------------------

	err = errAnyone
	println("\n----- errAnyone -----")
	logError(err)

	// --------------------------------------------------

	err = fmt.Errorf("error occured: %w", errAnyone)
	println("\n----- fmt.Errorf(\"error occured: %w\", errAnyone) -----")
	logError(err)

	// --------------------------------------------------

	err = fmt.Errorf("%w", &errStructone{"structone error"})
	println("\n----- fmt.Errorf(\"%w\", &errStruct{}) -----")
	logError(err)

	// --------------------------------------------------

	err = fmt.Errorf("%w: %w", errAnytwo, errAnyone)
	println("\n----- fmt.Errorf(\"%w: %w\", errAnytwo, errAnyone) -----")
	logError(err)

	// --------------------------------------------------

	err = fmt.Errorf("%w: %w", &errStructtwo{"error structtwo"}, &errStructone{"error structone"})
	println("\n----- fmt.Errorf(\"%w: %w\", &errStructone{}, &errStructone{}) -----")
	logError(err)

	// --------------------------------------------------

	err = fmt.Errorf("%w: %w: %w: %w", errAnytwo, &errStructone{"structone error"}, errors.New("new error"), errAnyone)
	println("\n----- fmt.Errorf(\"%w: %w: %w: %w\", errAnytwo, &errStruct{}, errors.New(), errAnyone) -----")
	logError(err)
}
