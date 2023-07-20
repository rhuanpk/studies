package main

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

// go playground: https://go.dev/play/p/7VYsR6BpMEf
func main() {

	str := "strawberry, blueberry, raspberry"
	println("\n>>> string to be used as an example (\"str\" is the variable name): \"" + str + "\"!\n")

	println("****************************************************************************************************************************\n")

	fmt.Printf(`strings.Split(str, ", "): `+"%#v\n", strings.Split(str, ", "))

	println("\n----------------------------------------------------------------------------------------------------------------------------\n")

	fmt.Printf(`strings.SplitAfter(str, ", "): `+"%#v\n", strings.SplitAfter(str, ", "))

	println("\n----------------------------------------------------------------------------------------------------------------------------\n")

	fmt.Printf(`strings.SplitN(str, ", ", 2): `+"%#v\n", strings.SplitN(str, ", ", 2))

	println("\n----------------------------------------------------------------------------------------------------------------------------\n")

	fmt.Printf(`strings.SplitAfterN(str, ", ", 2): `+"%#v\n", strings.SplitAfterN(str, ", ", 2))

	println("\n----------------------------------------------------------------------------------------------------------------------------\n")

	fmt.Printf(`strings.Fields(str): `+"%#v\n", strings.Fields(str))

	println("\n----------------------------------------------------------------------------------------------------------------------------\n")

	fmt.Printf(
		`strings.FieldsFunc(str, func(r rune) bool { return !unicode.IsLetter(r) }): `+"%#v\n",
		strings.FieldsFunc(
			str,
			func(r rune) bool {
				return !unicode.IsLetter(r)
			},
		),
	)

	println("\n----------------------------------------------------------------------------------------------------------------------------\n")

	regex := regexp.MustCompile(",\\s*")
	fmt.Println(`regex := regexp.MustCompile(",\\s*")`)
	fmt.Printf(`regexp.Split(str, -1): `+"%#v\n", regex.Split(str, -1))

	println("\n----------------------------------------------------------------------------------------------------------------------------\n")

	before, after, found := strings.Cut(str, ", ")
	fmt.Printf(`strings.Cut(str, ", "):`+"\n\tbefore: %s\n\tafter: %s\n\tseparator found: %t\n", before, after, found)

	println("\n****************************************************************************************************************************")

}
