package main

import (
	"fmt"
	"strconv"
)

func main()  {
	
	x := 2.4
	y := 2
	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	// Cuidado...
	fmt.Println("Teste " + string(97))

	// Int para string
	fmt.Println("Teste " + strconv.Itoa(123))

	// String para int
	num, _ := strconv.Atoi("123")
	fmt.Println(num - 122)

	// Booleano
	b, _ := strconv.ParseBool("true")
	if b {
		fmt.Println("Verdadeiro.")
	}

	/*
		Quando há métodos que retornam
		valor padrão ou valor de erro,
		eles podem ser atribuidos de
		respectivamente.
	*/

}