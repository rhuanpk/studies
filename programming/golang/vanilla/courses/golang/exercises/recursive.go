//go:build recursive

/*
	Create some program that:
		- when is a function recursive?
			R: a functions that call itself which have a recursive case and a stop point.
		- can a recursive function call non-recursive functions?
			R: yes.

	source: https://golangr.com/recursion/
*/

package main

import "fmt"

func factorial(multiplier uint) uint {
    if multiplier == 0 {
        return 1
    }
    return multiplier * factorial(multiplier - 1)
}

func main(){
    result := factorial(3)
    fmt.Println(result)
}
