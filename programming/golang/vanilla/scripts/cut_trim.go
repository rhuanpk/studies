package main

import (
	"fmt"
	"strings"
)

const str = `#abcdef,ghijklm,nopqrs,tuvwxyx#`

func main() {
	println("str :=", str)

	println("\n>>>>> cut functions <<<<<")

	cutBefore, cutAfter, cutFound := strings.Cut(str, ",")
	println("\nstrings.Cut(str, \",\"):")
	fmt.Printf("\t- before: %s\n\t- after: %s\n\t- found: %t\n", cutBefore, cutAfter, cutFound)

	cutPrefixAfter, cutPrefixFound := strings.CutPrefix(str, "#")
	println("\nstrings.CutPrefix(str, \"#\"):")
	fmt.Printf("\t- after: %s\n\t- found: %t\n", cutPrefixAfter, cutPrefixFound)

	cutSuffixBefore, cutSuffixFound := strings.CutSuffix(str, "#")
	println("\nstrings.CutSuffix(str, \"#\"):")
	fmt.Printf("\t- before: %s\n\t- found: %t\n", cutSuffixBefore, cutSuffixFound)

	println("\n>>>>> trim functions <<<<<")

	trim := strings.Trim(str, "#")
	println("\nstrings.Trim(str, \"#\"):")
	fmt.Printf("\t- return: %s\n", trim)

	trimPrefix := strings.TrimPrefix(str, "#")
	println("\nstrings.TrimPrefix(str, \"#\"):")
	fmt.Printf("\t- return: %s\n", trimPrefix)

	trimSuffix := strings.TrimSuffix(str, "#")
	println("\nstrings.TrimSuffix(str, \"#\"):")
	fmt.Printf("\t- return: %s\n", trimSuffix)
}
