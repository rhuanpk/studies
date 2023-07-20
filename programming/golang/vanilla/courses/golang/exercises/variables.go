//go:build variables

/*
	Create some program that:
		- calculate the year given the date of birth and age.
		- create a program that calculates the average weight of 5 people.

	Source: https://golangr.com/variables/
*/

package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {

	const error_exit_status int = 1

	read := bufio.NewReader(os.Stdin)
	age_slice := make([]int, 1, 1)
	max_interactions := 4

	for index := 0; index <= max_interactions; index++ {

		print("Age to input: ")
		aux, err := read.ReadString('\n')
		aux = strings.Replace(aux, "\n", "", -1)

		if err != nil {

			println("Input error!")
			println("Error message:", err.Error())
			os.Exit(error_exit_status)

		} else {

			converteded, err := strconv.Atoi(aux)

			if err != nil {

				println("Conversion error!")
				println("Error message:", err.Error())
				os.Exit(error_exit_status)

			} else {

				age_slice = append(age_slice, converteded)

			}

		}

	}

	var avarage int
	print("\n")

	for id_people, age := range age_slice {

		println("id_people:", id_people, "\nage:", age, "\n")
		avarage += age

	}

	slice_size := len(age_slice) - 1
	avarage /= slice_size
	println("slice_size:", slice_size, "\navarage:", avarage)

}
