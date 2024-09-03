
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func notaParaConceito(nota float64) {

	// Exemplo de if comum
	if nota <= 10 && nota >= 9 {
		fmt.Println("A")
	} else if nota < 9 && nota >= 8 {
		fmt.Println("B")
	} else if nota < 8 && nota >= 5 {
		fmt.Println("C")
	} else if nota < 5 && nota >= 3 {
		fmt.Println("D")
	} else {
		fmt.Println("E")
	}

}


func random() int {

	s := rand.NewSource(time.Now().UnixMicro())
	r := rand.New(s)
	return r.Intn(10)

}

func main() {

	notaParaConceito(9.9)
	notaParaConceito(7.3)
	notaParaConceito(4.9)
	notaParaConceito(2.5)
	notaParaConceito(8.0)

	// Exemplo de if com inicialização de variável no bloco
	if i := random(); i > 5 {
		fmt.Println("Número", i, "---> GANHOU!")
	} else {
		fmt.Println("Número", i, "---> PERDEU!")
	}

}
