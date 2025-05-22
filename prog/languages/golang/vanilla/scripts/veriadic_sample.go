package main

import "fmt"

func function(args ...any) {
	if args != nil {
		fmt.Printf("args is not nil: %#v\n", args)
	}
	for index := range args {
		fmt.Println("iterate over index:", index)
	}
}

// go playground: https://go.dev/play/p/RkMEdXmzFuu
func main() {
	println("> function()")
	function()

	println("\n> var foo []any")
	var foo []any
	function(foo...)

	println("\n> bar := []any{}")
	bar := []any{}
	function(bar...)

	println("\n> function(nil)")
	function(nil)
}
