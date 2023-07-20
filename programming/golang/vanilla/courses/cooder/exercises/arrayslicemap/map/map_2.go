package main

import "fmt"

func main() {
	
	funcionarios_salarios := map[string]float64{
		"José João": 1230.45,
		"Gabriela Silva": 6780.90,
		"Pedro Junior": 1590.32,
	}

	funcionarios_salarios["Rafael Filho"] = 1225.88
	delete(funcionarios_salarios, "inexistente")

	for name, wage := range funcionarios_salarios {
		fmt.Printf("Funcionário: %s - Salário: R$ %.2f\n", name, wage)
	}

}
