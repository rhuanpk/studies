package main

import (
	"fmt"
	"strings"
)

const str = `#abcdef,ghijklm,nopqrs,tuvwxyx#`

// go playground: https://go.dev/play/p/wHMQTdYOZUW
func main() {
	println("str :=", str)

	println("\n>>>>> cut functions <<<<<")

	cutBefore, cutAfter, cutFound := strings.Cut(str, ",")
	println("\nstrings.Cut(str, \",\"):")
	fmt.Printf("- before: %s\n- after: %s\n- found: %t\n", cutBefore, cutAfter, cutFound)

	cutPrefixAfter, cutPrefixFound := strings.CutPrefix(str, "#")
	println("\nstrings.CutPrefix(str, \"#\"):")
	fmt.Printf("- after: %s\n- found: %t\n", cutPrefixAfter, cutPrefixFound)

	cutSuffixBefore, cutSuffixFound := strings.CutSuffix(str, "#")
	println("\nstrings.CutSuffix(str, \"#\"):")
	fmt.Printf("- before: %s\n- found: %t\n", cutSuffixBefore, cutSuffixFound)

	println("\n>>>>> trim functions <<<<<")

	trim := strings.Trim(str, "#")
	println("\nstrings.Trim(str, \"#\"):")
	fmt.Printf("- return: %s\n", trim)

	trimPrefix := strings.TrimPrefix(str, "#")
	println("\nstrings.TrimPrefix(str, \"#\"):")
	fmt.Printf("- return: %s\n", trimPrefix)

	trimSuffix := strings.TrimSuffix(str, "#")
	println("\nstrings.TrimSuffix(str, \"#\"):")
	fmt.Printf("- return: %s\n", trimSuffix)

	trimLeft := strings.TrimLeft(str, "bc#a")
	println("\nstrings.TrimLeft(str, \"bc#a\"):")
	fmt.Printf("- return: %s\n", trimLeft)

	trimRight := strings.TrimRight(str, "z#xy")
	println("\nstrings.TrimRight(str, \"z#xy\"):")
	fmt.Printf("- return: %s\n", trimRight)
}
