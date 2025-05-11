//go:build methods

/*
	Create some program that:
		- create a method that sums two numbers.
		- create a method that calls another method.

	Source: https://golangr.com/methods/
*/

package main

import "fmt"

func sum(number_1, number_2 float64) float64 {
	return number_1 + number_2
}

func printSum(str string, number_1, number_2 float64) {
	fmt.Println(str, sum(number_1, number_2))
}

func main() {

	number_1 := 12.0
	number_2 := 15.7

	printSum("result:", number_1, number_2)
	printSum("result:", 32.83, 12.17)

}
