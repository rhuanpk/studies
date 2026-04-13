package main

import "errors"

type errOne struct {
	message string
}

func (pr errOne) Error() string {
	return pr.message
}

type errTwo struct {
	message string
}

func (pr errTwo) Error() string {
	return pr.message
}

func returnErrOne() error {
	return errOne{"one"}
}

func returnErrTwo() error {
	return errTwo{"two"}
}

func log(prefix string, err error) {
	print(prefix + ": err != nil: ")
	if err != nil {
		println(true)
	} else {
		println(false)
	}

	print(prefix + ": errors.As(err, new(errOne)): ")
	if errors.As(err, new(errOne)) {
		println(true)
	} else {
		println(false)
	}

	print(prefix + ": errors.As(err, new(errTwo)): ")
	if errors.As(err, new(errTwo)) {
		println(true)
	} else {
		println(false)
	}
}

// go playground: https://go.dev/play/p/qtbDZBGtYiL
func main() {
	errOne := returnErrOne()
	errTwo := returnErrTwo()

	err := errors.Join(errOne, errTwo)
	log("join", err)

	// it could also be made recursive wrap if we split this block in a function
	if ass, ok := err.(interface{ Unwrap() []error }); ok {
		for _, err := range ass.Unwrap() {
			log(err.Error(), err)
		}
	}
}
