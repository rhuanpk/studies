package main

import (
	"bufio"
	"os"
)

// Playground: https://go.dev/play/p/jLf8t5Ng8Mt
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	print("Do you want know about nil (0) or empty (1)?\nA: ")
	scanner.Scan()
	if "0" == scanner.Text() {
		nilCompare()
	} else if "1" == scanner.Text() {
		emptyCompare()
	} else {
		println("Wrong answer :P")
	}
}

func nilCompare() {
	println("Only map declaration:")
	var mapa1 map[string]any
	if mapa1 == nil {
		println("nil")
	} else {
		println("filled")
	}

	println("\nPost initialization with make but empty:")
	var mapa2 map[string]any
	mapa2 = make(map[string]any)
	if mapa2 == nil {
		println("nil")
	} else {
		println("filled")
	}

	println("\nDirectly initialization with composite literal but empty:")
	mapa3 := map[string]any{}
	if mapa3 == nil {
		println("nil")
	} else {
		println("filled")
	}

	println("\nAtributted value after initialized with make after declarated alone:")
	var mapa4 map[string]any
	mapa4 = make(map[string]any)
	mapa4["hello"] = "world"
	if mapa4 == nil {
		println("nil")
	} else {
		println("filled")
	}

	println("\nInitialized with composite literal:")
	mapa5 := map[string]any{"hello": "world"}
	if mapa5 == nil {
		println("nil")
	} else {
		println("filled")
	}
}

func emptyCompare() {
	println("Only map declaration:")
	var mapa1 map[string]any
	if len(mapa1) == 0 {
		println("empty")
	} else {
		println("filled")
	}

	println("\nPost initialization with make but empty:")
	var mapa2 map[string]any
	mapa2 = make(map[string]any)
	if len(mapa2) == 0 {
		println("empty")
	} else {
		println("filled")
	}

	println("\nDirectly initialization with composite literal but empty:")
	mapa3 := map[string]any{}
	if len(mapa3) == 0 {
		println("empty")
	} else {
		println("filled")
	}

	println("\nAtributted value after initialized with make after declarated alone:")
	var mapa4 map[string]any
	mapa4 = make(map[string]any)
	mapa4["hello"] = "world"
	if len(mapa4) == 0 {
		println("empty")
	} else {
		println("filled")
	}

	println("\nInitialized with composite literal:")
	mapa5 := map[string]any{"hello": "world"}
	if len(mapa5) == 0 {
		println("empty")
	} else {
		println("filled")
	}
}
