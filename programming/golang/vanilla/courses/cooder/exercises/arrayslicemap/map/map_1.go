package main

import "fmt"

func main() {
	
	// Mapas devem ser inicializados!
	// var aprovados map[int]string

	aprovados := make(map[int]string)

	aprovados[12345678978] = "Maria"
	aprovados[98765432100] = "Pedro"
	aprovados[95135745682] = "Ana"

	for cpf, name := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", name, cpf)
	}

	fmt.Println(aprovados[95135745682])
	delete(aprovados, 95135745682)
	fmt.Println(aprovados[95135745682])
	
}
