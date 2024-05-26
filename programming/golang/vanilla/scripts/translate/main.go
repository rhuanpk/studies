package main

import (
	"fmt"
	"log"

	"project/parse"
)

func main() {
	const (
		code = "BR"
		name = "Brasil"
	)

	c2n, err := parse.C2N(code)
	if err != nil {
		log.Println("error in parse code to name:", err)
	} else {
		fmt.Printf("Code to name: %s -> %s\n", code, c2n)
	}

	n2c, err := parse.N2C(name)
	if err != nil {
		log.Println("error in parse name to code:", err)
	} else {
		fmt.Printf("Name to code: %s -> %s\n", name, n2c)
	}
}
