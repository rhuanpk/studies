//go:build pointers

/*
	Create some program that:
		- where are variables stored in the computer?
			R: memory.
		- what is a pointer?
			R: a pointer is a variable whose address is the direct memory address of another variable.
		- how can you declare a pointer?
			R: use the asterisk, `var variable_name *<var_type>`.

	Source: https://golangr.com/pointers/
*/

package main

func main() {

	letter_a := 5
	var letter_b *int = &letter_a

	println(
		"content of [letter_a]:", letter_a, "\n"+
			"content of [letter_b]:", *letter_b,
	)

	letter_a += 3

	println(
		"content of [letter_a]:", letter_a, "\n"+
			"content of [letter_b]:", *letter_b,
	)

	*letter_b -= 1

	println(
		"content of [letter_a]:", letter_a, "\n"+
			"content of [letter_b]:", *letter_b,
	)

}
