package main

import "fmt"

func main() {

	// Go não permite aritmética de ponteiros

	i := 1
	var p *int = nil; p = &i
	// p := &i
	*p++
	i++
	fmt.Println(p, *p, i, &i)

}
