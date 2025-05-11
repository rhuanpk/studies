package main

import (
	"fmt"
	"time"
)

func main() {

	// Não tem while em go
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	// For clássico
	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d ", i)
	}

	// For clássico
	fmt.Println("\nMisturando...")
	for i := 0; i < 10; i++ {
		if i % 2 == 0 {
			fmt.Println("Par ")
		} else {
			fmt.Println("Impar ")
		}
	}

	// Laço infinito (like infinite while)
	for {
		fmt.Println("Para sempre...")
		time.Sleep(time.Second)
	}

}