package main

import "fmt"

// go playground: https://go.dev/play/p/dSoVtbivadx
func main() {
	array := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("array to example: %#v\n\n", array)

	println("array[:5] -> getting from first index to 4:")
	fmt.Println(array[:5])

	println("\narray[3:7] -> getting from index 3 to 6:")
	fmt.Println(array[3:7])

	println("\narray[8:] -> getting from index 8 to last:")
	fmt.Println(array[8:])

	println("\narray[1:] -> excludding first index:")
	fmt.Println(array[1:])

	println("\narray[:len(array)-1] -> excludding last index:")
	fmt.Println(array[:len(array)-1])

	println("\narray[3:8] -> trimming from first index to 3 and 7 to last:")
	fmt.Println(array[4:7])

	println("\nappend(array[:4], array[6:]...) -> cutting indexes 4 and 5:")
	fmt.Println(append(array[:4], array[6:]...))

	fmt.Println("\nsyntax is 'array[<start>:<end>]' since indexes are 'inclusive' and 'exclusive', respectively")
}
