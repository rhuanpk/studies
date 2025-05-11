//go:build variadic_functions

/*
	Create some program that:
		- create a variadic function that prints the names of students.

	source: https://golangr.com/variadic-functions/
*/

// Básicamente, o "operador veriádico" expande o array/slice, ou seja, caso você passe como argumento `any_array`, você está passando o objeto do array, ou seja, está passando `[primeiro segundo terceiro quarto]`, caso você passe como argumento `any_array...` você está passando o array com os valores expandidos, cada valor do array como um argumento separado, ou seja, está passando `primeiro, segundo, terceiro, quarto` (equivalente a `${@}` no bash).

package main

import (
	"fmt"
)

func printStudentsName(names ...string) {
	for key, name := range names {
		fmt.Printf("[%d] - %s\n", key, name)
	}
}

func main() {
	names := []string{"rhuan", "patriky", "barbara", "esther"}
	printStudentsName("roberto", "andre")
	println("")
	printStudentsName(names...)
}
