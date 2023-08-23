package main

import "fmt"

type obj struct {
	field1 string
	field2 *int
}

// type jbo struct {
// 	field3 float64
// 	field4 *bool
// }

// go playground: https://go.dev/play/p/0DuphfW3a4A 
func main() {
	var foo obj
	var bar *obj
	boo := obj{}
	baz := &obj{}

	var xpto1 []obj
	var xpto2 []*obj
	xpto3 := []obj{}
	xpto4 := []*obj{}

	array1 := make([]obj, 0)
	array2 := make([]*obj, 0)
	array3 := make([]obj, 1)
	array4 := make([]*obj, 1)

	fmt.Printf(
		"%#v\n---\n%#v\n---\n%#v\n---\n%#v\n**********\n%#v\n---\n%#v\n---\n%#v\n---\n%#v\n**********\n%#v\n---\n%#v\n---\n%#v\n---\n%#v\n",
		foo,
		bar,
		boo,
		baz,
		xpto1,
		xpto2,
		xpto3,
		xpto4,
		array1,
		array2,
		array3,
		array4,
	)
}
