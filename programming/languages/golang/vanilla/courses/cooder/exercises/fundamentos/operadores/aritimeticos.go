package main

import (
	"fmt"
	"math"
	"time"
)

// Lógicos
func compras(trab1, trab2 bool) (bool, bool, bool) {
	
	comprarTv50 := trab1 && trab2 // END
	comprarTv32 := trab1 != trab2 // "XOR"
	comprarSorvete := trab1 || trab2 // OR

	return comprarTv50, comprarTv32, comprarSorvete

}

// "Ternário"
func naoTernario(xpto float64) string {
	if xpto >= 1 { return "true" } else { return"false" }
}

func main()  {
	
	// Operadores aritméticos simples.
	a := 3
	b := 2

	fmt.Println("Soma =>", a + b)
	fmt.Println("Subtração =>", a - b)
	fmt.Println("Divisão =>", a / b)
	fmt.Println("Multiplicação =>", a * b)
	fmt.Println("Módulo =>", a % b)

	// Operadores aritméticos bit a bit (bitwise).
	fmt.Println("And =>", a & b)
	fmt.Println("Or =>", a | b)
	fmt.Println("Xor =>", a ^ b)

	// Outras operações usando math
	c := 3.0
	d := 2.0

	fmt.Println("Maior =>", math.Max(float64(a), float64(b)))
	fmt.Println("Menor =>", math.Min(c, d))
	fmt.Println("Exponencialçao =>", math.Pow(c, d))

	// Atribuições
	a += b // Aditiva
	a -= b // Subtrativa
	a *= b // Multiplicativa
	a /= b // Divisiva
	a %= b // Modulativa?

	// Trocar o valor de duas variáveis
	a, b = b, a 

	// Relacionais
	fmt.Println("Strings:", "Banana" == "Banana")
	fmt.Println("!=", 3 != 2)
	fmt.Println("<", 3 < 2)
	fmt.Println(">", 3 > 2)
	fmt.Println("<=", 3 <= 2)
	fmt.Println(">=", 3 >= 2)

	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)

	fmt.Println("Datas:", d1 == d2)
	fmt.Println("Datas:", d1.Equal(d2))

	type Pessoa struct {
		Nome string
	}

	p1 := Pessoa{"João"}
	p2 := Pessoa{"João"}

	fmt.Println("Pessoas:", p1 == p2)

	// Lógicos
	tv50, tv32, sorvete := compras(true, false)
	fmt.Printf("Tv50: %t, Tv32: %t, Sorvete: %t, Saudável: %t", tv50, tv32, sorvete, !sorvete)

	// Unários (apenas postfix)
	i := 1
	j := 2

	i++
	j--

	fmt.Println(i, j)
	// fmt.Println(i == j--)

	// Ternário (not exist)
	fmt.Println(naoTernario(float64(i)))

}
