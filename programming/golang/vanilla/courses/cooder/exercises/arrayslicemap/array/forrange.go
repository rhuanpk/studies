package main

import "fmt"

func main() {
	
	// Dessa forma o compilador irá atribuir automágicamente o tamanho do array
	numbers := [...] int {1, 2, 3, 4, 5}

	for key, value := range numbers {
		fmt.Printf("%d) %d\n", key, value)
	}

	// Forma de suprimir a chave e trabalhar somente com o valor da chave
	for _, value := range numbers {
		fmt.Println(value)
	}

}
