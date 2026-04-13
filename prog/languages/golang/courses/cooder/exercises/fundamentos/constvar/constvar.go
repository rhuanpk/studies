package main

// Import também suporta estrutura de bloco e esquema de alias.
import (
    "fmt"
    m "math"
)

func main () {
	// Constante (tipo declarado de forma explicita).
	const pi float64 = 3.14

	// Variável (tipo (float64) inferido pelo compilador).
	var raio = 5.29

	// Criar variável já inicializando e logo depois utilizando a mesma.
	area := pi * m.Pow(raio, 2)
	fmt.Println("A área da circunferência é", area)

	// Constantes e Variáveis podem ser criadas dentro de blocos também:
	const (
		a = 1.0
		b = 2.0
	)
	var (
		c = 3
		d = 4
	)
	fmt.Println(a, b, c, d)

	// Declaração múltipla em única linha fazendo o recebimento de forma respectiva
	var e, f bool = true, false
	fmt.Println(e, f)

	g, h, i := 2, false, "epa!"
	fmt.Println(g, h, i)
}