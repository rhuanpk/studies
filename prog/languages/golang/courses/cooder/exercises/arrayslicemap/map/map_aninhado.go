package main

import "fmt"

func main() {

	funcionarios_por_letra := map[string]map[string]float64{
		"G": {
			"Gabriel Silva": 15456.78,
			"Gustavo Junior": 82738.12,
		},
		"H": {
			"Habadon Naste": 41234.13,
		},
		"R": {
			"Rhuan Patriky": 92834.82,
			"Roberto Gastro": 18234.12,
		},
	}

	delete(funcionarios_por_letra, "G")

	for letra, funcionarios := range funcionarios_por_letra {
		fmt.Println(letra, funcionarios)
	}

}
