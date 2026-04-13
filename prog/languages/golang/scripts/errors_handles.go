package main

import (
	"errors"
	"fmt"
)

var (
	errAnyOne = errors.New("anyone error")
	errAnyTwo = errors.New("anytwo error")
)

// --------------------------------------------------

type errStructOne struct {
	message string
}

func (so errStructOne) Error() string {
	return "so.message: " + so.message
}

// --------------------------------------------------

type errStructTwo struct {
	message string
}

func (st errStructTwo) Error() string {
	return "st.message: " + st.message
}

// --------------------------------------------------

type errStructThree struct {
	message   string
	errAnyOne error
	errAnyTwo error
}

func (st errStructThree) Error() string {
	return "st.message: " + st.message
}

func (st errStructThree) Unwrap() error {
	return errors.Join(st.errAnyOne, st.errAnyTwo)
}

// --------------------------------------------------

func returnError() error {
	return errAnyOne
}

func returnWrapedError() error {
	return fmt.Errorf("error occurred: %w", errAnyOne)
}

func returnStructError(message string) error {
	return errStructOne{message}
}

func returnWrapedStructError(message string) error {
	return fmt.Errorf("error ocurred: %w", errStructOne{message})
}

func checkError(err error, errType string) {
	print(errType, " != nil: ")
	if err != nil {
		println(true)

		// println("println("+errType+".Error()):", err.Error())
		fmt.Println("fmt.Println("+errType+"):", err)

		print("errors.Is(" + errType + ", errAnyOne): ")
		if errors.Is(err, errAnyOne) {
			println(true)
		} else {
			println(false)
		}

		print("errors.Is(" + errType + ", errAnyTwo): ")
		if errors.Is(err, errAnyTwo) {
			println(true)
		} else {
			println(false)
		}

		print("errors.As(" + errType + ", &errStructOne{}): ")
		if errors.As(err, &errStructOne{}) {
			println(true)
			print("type assertion: ")
			if _, ok := err.(errStructOne); ok {
				println(true)
			} else {
				println(false)
			}
		} else {
			println(false)
		}

		print("errors.As(" + errType + ", &errStructTwo{}): ")
		if errors.As(err, &errStructTwo{}) {
			println(true)
			print("type assertion: ")
			if _, ok := err.(errStructTwo); ok {
				println(true)
			} else {
				println(false)
			}
		} else {
			println(false)
		}

		print("errors.As(" + errType + ", &errStructThree{}): ")
		if errors.As(err, &errStructThree{}) {
			println(true)
			print("type assertion: ")
			if _, ok := err.(errStructThree); ok {
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

// go playground: https://go.dev/play/p/76VYLovvO5V
func main() {
	err := returnError()

	println("----- err := returnError() -----")

	print("err != nil: ")
	if err != nil {
		println(true)
	} else {
		println(false)
	}

	print("errors.Is(err, errAnyOne): ")
	if errors.Is(err, errAnyOne) {
		println(true)
	} else {
		println(false)
	}

	unwrap := errors.Unwrap(err)
	fmt.Printf("unwrap (%%v): %v\nunwrap (%%#v): %#v\n", unwrap, unwrap)

	// --------------------------------------------------

	err = returnWrapedError()

	println("\n----- err = returnWrapedError() -----")

	print("err != nil: ")
	if err != nil {
		println(true)
	} else {
		println(false)
	}

	print("errors.Is(err, errAnyOne): ")
	if errors.Is(err, errAnyOne) {
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

	print("errors.Is(unwrapedError, errAnyOne): ")
	if errors.Is(unwrapedError, errAnyOne) {
		println(true)
	} else {
		println(false)
	}

	unwrap = errors.Unwrap(unwrapedError)
	fmt.Printf("unwrap (%%v): %v\nunwrap (%%#v): %#v\n", unwrap, unwrap)

	// --------------------------------------------------

	var structOneError errStructOne
	err = returnStructError("structone error")

	println("\n----- err = returnStructError() -----")

	print("err != nil: ")
	if err != nil {
		println(true)
	} else {
		println(false)
	}

	print("errors.As(err, &structOneError): ")
	if errors.As(err, &structOneError) {
		println(true)
	} else {
		println(false)
	}

	unwrap = errors.Unwrap(err)
	fmt.Printf("unwrap (%%v): %v\nunwrap (%%#v): %#v\n", unwrap, unwrap)

	// --------------------------------------------------

	err = returnWrapedStructError("structone error")

	println("\n----- err = returnWrapedStructError() -----")

	print("err != nil: ")
	if err != nil {
		println(true)
	} else {
		println(false)
	}

	print("errors.As(err, &structOneError): ")
	if errors.As(err, &structOneError) {
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
	if errors.As(unwrapedError, &structOneError) {
		println(true)
	} else {
		println(false)
	}

	unwrap = errors.Unwrap(unwrapedError)
	fmt.Printf("unwrap (%%v): %v\nunwrap (%%#v): %#v\n", unwrap, unwrap)

	println("\n##################################################")

	err = errors.New("new error")
	println("\n----- errors.New(\"new error\") -----")
	logError(err)

	// --------------------------------------------------

	err = errAnyOne
	println("\n----- errAnyOne -----")
	logError(err)

	// --------------------------------------------------

	err = errors.Join(errAnyOne, errAnyTwo)
	println("\n----- errors.Join(errAnyOne, errAnyTwo) -----")
	logError(err)

	// --------------------------------------------------

	err = errors.Join(errAnyOne, errStructOne{"structone error"})
	println("\n----- errors.Join(errAnyOne, errStructOne{}) -----")
	logError(err)

	// --------------------------------------------------

	err = errors.Join(errStructOne{"structone error"}, errStructThree{"structthree error", errAnyOne, errAnyTwo})
	println("\n----- errors.Join(errStructOne{}, errStructThree{}) -----")
	logError(err)

	// --------------------------------------------------

	err = fmt.Errorf("%w", errAnyOne)
	println("\n----- fmt.Errorf(\"%w\", errAnyOne) -----")
	logError(err)

	// --------------------------------------------------

	err = fmt.Errorf("%w", errStructOne{"structone error"})
	println("\n----- fmt.Errorf(\"%w\", errStructOne{}) -----")
	logError(err)

	// --------------------------------------------------

	err = fmt.Errorf("%w: %w", errAnyTwo, errAnyOne)
	println("\n----- fmt.Errorf(\"%w: %w\", errAnyTwo, errAnyOne) -----")
	logError(err)

	// --------------------------------------------------

	err = fmt.Errorf("%w: %w", errStructTwo{"error structTwo"}, errStructOne{"structone error"})
	println("\n----- fmt.Errorf(\"%w: %w\", errStructTwo{}, errStructOne{}) -----")
	logError(err)

	// --------------------------------------------------

	err = fmt.Errorf("%w: %w: %w: %w", errAnyTwo, errStructOne{"structone error"}, errors.New("new error"), errAnyOne)
	println("\n----- fmt.Errorf(\"%w: %w: %w: %w\", errAnyTwo, errStructOne{}, errors.New(), errAnyOne) -----")
	logError(err)

	// --------------------------------------------------

	err = fmt.Errorf("%w: %w: %w: %w", errAnyTwo, errStructThree{"structthree error", errAnyOne, errAnyTwo}, errors.New("new error"), errAnyOne)
	println("\n----- fmt.Errorf(\"%w: %w: %w: %w\", errAnyTwo, errStructThree{}, errors.New(), errAnyOne) -----")
	logError(err)
}
