package main

import "fmt"

// go playground: https://go.dev/play/p/NwaOXupoc61

func main() {
	array := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("array to example: %#v\n", array)
	println("-> array[:5] <- getting from first index to 4:")
	fmt.Println(array[:5])
	println("-> array[3:7] <- getting from index 3 to 6:")
	fmt.Println(array[3:7])
	println("-> array[8:] <- getting from index 8 to last:")
	fmt.Println(array[8:])
	println("-> array[:len(array)-1] <- excludding last index:")
	fmt.Println(array[:len(array)-1])
}
