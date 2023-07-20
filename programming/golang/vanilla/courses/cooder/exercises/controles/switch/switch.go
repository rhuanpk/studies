package main

import (
	"fmt"
	"time"
)

func notaParaConceito(nota float64) string {

	// Switch declarado com expressão busca nos cases o valor correspondente
	switch int(nota) {
		case 10:
			fallthrough
		case 9:
			return "A"
		case 8, 7:
			return "B"
		case 6, 5:
			return "C"
		case 4, 3:
			return "D"
		case 2, 1, 0:
			return "E"
		default:
			return "Nota inválida!"
	}

}

func tipo(tipo interface{}) string {

	// Switch comparando tipos de dados
	switch tipo.(type) {
		case int:
			return "inteiro"
		case float32, float64:
			return "float"
		case string:
			return "string"
		case func():
			return "function"
		case bool:
			return "boolean"
		default:
			return "undefined"
	}

}

func main() {

	fmt.Println(notaParaConceito(9.8))
	fmt.Println(notaParaConceito(6.9))
	fmt.Println(notaParaConceito(2.1))
	fmt.Println(notaParaConceito(-3.7))
	fmt.Println(notaParaConceito(12.3))

	// Switch declarado com expressão busca nos cases valor boolean true
	tempo := time.Now()
	switch {
		case tempo.Hour() < 12:
			fmt.Println("Bom dia!")
		case tempo.Hour() < 18:
			fmt.Println("Boa tarde!")
		default:
			fmt.Println("Boa noite!")
	}

	fmt.Println(tipo(2.7))
	fmt.Println(tipo(4))
	fmt.Println(tipo(float32(3.4)))
	fmt.Println(tipo("hello world!"))
	fmt.Println(tipo(time.Now()))
	fmt.Println(tipo(false))
	fmt.Println(tipo(func () {}))

}
