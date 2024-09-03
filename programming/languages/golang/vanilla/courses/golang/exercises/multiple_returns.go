//go:build multiple_returns

/*
	Create some program that:
		- change the return values from 2,4 to “hello”,”world”. Does it still work?
			R: yes.
		- can a combination of strings and numbers be used?
			R: yes.

	Source: https://golangr.com/multiple-return/
*/

package main

import "strconv"

func returnHelloWorld() (int, string, string) {
	return 7, "hello", "world"
}

func main() {
	code, str_1, str_2 := returnHelloWorld()
	println("[" + strconv.Itoa(code) + "] " + str_1 + " " + str_2 + "!")
}
