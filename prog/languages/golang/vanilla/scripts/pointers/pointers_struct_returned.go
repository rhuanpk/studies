package main

import "fmt"

type xpto struct {
	field string
}

func newXptoNil() *xpto {
	return nil
}

func newXptoNew() *xpto {
	return new(xpto)
}

func newXptoComposite() *xpto {
	return &xpto{}
}

func newXptoAlloc() *xpto {
	return &xpto{"xpto"}
}

// go playground: https://go.dev/play/p/dtynd5sgt-F
func main() {
	xptoNil := newXptoNil()
	fmt.Printf("%#v: %v\n", xptoNil, xptoNil)
	// fmt.Printf("%#v: %v\n", xptoNil.field, xptoNil.field) // panic

	xptoNew := newXptoNew()
	fmt.Printf("\n%#v: %v\n", xptoNew, xptoNew)
	fmt.Printf("%#v: %v\n", xptoNew.field, xptoNew.field)

	xptoComposite := newXptoComposite()
	fmt.Printf("\n%#v: %v\n", xptoComposite, xptoComposite)
	fmt.Printf("%#v: %v\n", xptoComposite.field, xptoComposite.field)

	xptoAlloc := newXptoAlloc()
	fmt.Printf("\n%#v: %v\n", xptoAlloc, xptoAlloc)
	fmt.Printf("%#v: %v\n", xptoAlloc.field, xptoAlloc.field)
}
