package main

import "fmt"

func main() {

	/*
		O método "make" cria um "array virtual" pois o slice
		é um dado que faz referência a um array.
	*/

	/*
		Dessa forma você infoma que o slice será do tipo de dado inteiro
		e que terá 10 posições, logo, implicitamente o "array virtual"
		que foi criado para ele terá o mesmo tamanho/capacidade.
	*/ 
	s := make([]int, 10)
	s[9] = 12
	fmt.Println(s)

	/*
		Dessa forma você infoma que o slice será do tipo de dado inteiro
		e que terá 10 posições, porém, explicitamos a capacidade do
		"array virtual" que é criado para que o slice faça referência dele.
	*/ 
	s = make([]int, 10, 20)
	// len() método retornará o tamanho do array.
	// cap() método retornará o tamanho do array a qual o slice se referencia.
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Println(s, len(s), cap(s))

	/*
		Quando o tamanho slice passa da capacidade do array a qual ele se
		referência, automáticamente ele dobra o tamanho do array.
	*/
	s = append(s, 1)
	fmt.Println(s, len(s), cap(s))

}
