//go:build routines

/*
	Create some program that:
		- what is a goroutine?
			R: is the Golang multithread routines.
		- how can you turn a function into a goroutine?
			R: preffix the function call with 'go'.

	source: https://golangr.com/goroutines/
*/

package main

import "fmt"

func PrintMessage(msg string) {
   fmt.Println(msg)
}

func main() {
   go PrintMessage("go routine")
   PrintMessage("function")
   fmt.Scanln()
}
