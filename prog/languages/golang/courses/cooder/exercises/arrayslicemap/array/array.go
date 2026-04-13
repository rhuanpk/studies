package main

import "fmt"

func main() {
	
	// Estrutura homogênea (mesmo tipo), estática (fixo) e indexada (índices numérico que começam por 0).
	var notas [3]float64
	fmt.Println(notas)

	notas[0], notas[1], notas[2] = 7.8, 4.3, 9.1
	// notas[3] = 10.0
	fmt.Println(notas)

	var total float64
	for count := 0; count < len(notas); count++ {
		total += notas[count]
	}

	media := total / float64(len(notas))
	fmt.Printf("Média: %.1f\n", media)

}
