//go:build map

/*
	Create some program that:
		- what is a map?
			R: a Golang map is a unordered collection of key-value pairs.
		- is a map ordered?
			R: no.
		- what can you use a map for?
			R: to store key value pairs.

	Source: https://golangr.com/maps/
*/

package main

import "fmt"

func main() {

	// var any_map map[string]int

	// var any_map map[string]int = map[string]int{"first": 0}

	// var any_map map[string]int
	// any_map = map[string]int{"first": 0}

	// var any_map map[string]int = make(map[string]int)

	// any_map := map[string]int{"first": 0}

	// any_map := make(map[string]int)
	// any_map["first"] = 0

	// fmt.Println(any_map)

	// -------------------------------------------------------

	peoples := map[int]map[string]any{

		0: {
			"nome":  "rhuan",
			"idade": 22,
			"sexo":  "m",
		},

		1: {
			"nome":  "barbara",
			"idade": 19,
			"sexo":  "f",
		},
	}

	println("***** PEOPLES MAP *****\n-----------------------")

	for key := range peoples {

		fmt.Printf(
			"name: %s\nage: %d\nsex: %s\n",
			peoples[key]["nome"],
			peoples[key]["idade"],
			peoples[key]["sexo"],
		)

		println("-----------------------")

	}

}
