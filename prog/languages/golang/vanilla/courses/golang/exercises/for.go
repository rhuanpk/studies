//go:build for

/*
	Create some program that:
		- can for loops exist inside for loops?
			R: yes.
		- make a program that counts from 1 to 10.

	Source: https://golangr.com/for/
*/

package main

func main() {
	for index := 0; index <= 9; index++ {
		println("interaction:", index)
	}
}
