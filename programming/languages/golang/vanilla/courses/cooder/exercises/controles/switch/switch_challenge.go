package main

import "fmt"

func notaParaConceito(nota float64) string {

	switch {
		case nota <= 10 && nota >= 9:
			return "A"
		case nota < 9 && nota >= 8:
			return "B"
		case nota < 8 && nota >= 5:
			return "C"
		case nota < 5 && nota >= 3:
			return "D"
		case nota < 3 && nota >= 0:
			return "E"
		default:
			return "Nota inválida!"
	}

}

func main() {
	fmt.Println(notaParaConceito(-12.7))
}