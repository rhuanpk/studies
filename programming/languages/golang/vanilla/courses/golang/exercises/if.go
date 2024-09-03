//go:build if

/*
	Create some program that:
		- make a program that divides x by 2 if it’s greater than 0.
		- find out if if-statements can be used inside if-statements.

	Source: https://golangr.com/if/
*/

package main

import "fmt"

func main() {
	any_float := 3.0
	if any_float > 2 {
		any_float /= 2
	}
	fmt.Println(any_float)
}
