//go:build input

/*
	Create some program that:
		- make a program that lets the user input a name.
		- get a number from the console and check if it’s between 1 and 10.

	Source: https://golangr.com/keyboard-input/
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	read := bufio.NewReader(os.Stdin)

	fmt.Print("Name: ")
	name, _ := read.ReadString('\n')
	fmt.Println(
		"Wellcome",
		strings.Replace(name, "\n", "", -1)+
			"!",
	)

	fmt.Print("Number: ")
	input, _ := read.ReadString('\n')
	input = strings.Replace(input, "\n", "", -1)
	if number, err := strconv.Atoi(input); err != nil {
		fmt.Println("Error on converting!")
		fmt.Println("Leaving the program...")
	} else {
		if number >= 1 && number <= 10 {
			fmt.Println("Number between 1 and 10 (including)!")

		} else {
			fmt.Println("Number not between 1 and 10 (including)!")
		}
	}

}
