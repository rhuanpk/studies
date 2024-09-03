package main

import "fmt"

func main()  {
	
	/*
	
		Valores zeros são quando não atribuímos nenhum valor
		inicial para a variável, porém, ainda sim fica implicito
		o valor padrão que a própria linguagem atribui.

	*/

	var a int
	var b float64
	var c bool
	var d string
	var e *int

	fmt.Printf("%v %v %v %q %v\n", a, b, c, d, e)

}