//go:build arrays

/*
	Create some program that:
		- create an array with the number 0 to 10.
		- create an array of strings with names.

	Source: https://golangr.com/arrays/
*/

package main

import "fmt"

func main() {

	var int_array [5]int = [...]int{0, 1, 2, 3, 4}
	fmt.Println(int_array)

	string_array := []string{"rhuan", "patriky"}
	fmt.Println(string_array)

}
