//go:build slice

/*
	Create some program that:
		- take the string ‘hello world’ and slice it in two.
		- can you take a slice of a slice?
			R: yes.

	Source: https://golangr.com/slices/
*/

package main

import "fmt"

func main() {

	first_slice := make([]string, 0)
	second_slice := make([]string, 0)
	var switch_flag bool

	for _, value := range "hello world!" {

		// println(
		// 	"\nkey:", key,
		// 	"\nvalue:", string(value),
		// 	"\nhexa:", value,
		// )

		str := string(value)

		if (str == " ") || (str == "!") {
			switch_flag = true
			continue
		}

		if !switch_flag {
			first_slice = append(first_slice, str)
		} else {
			second_slice = append(second_slice, str)
		}

	}

	fmt.Println(
		"first_slice:", first_slice, "\n" +
		"second_slice:", second_slice,
	)

}
