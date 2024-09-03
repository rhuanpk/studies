package main

import "fmt"

func main() {
	
	array1 := [3]int{1, 2, 3}
	
	var slice1 []int
	slice1 = append(slice1, 4, 5, 6)
	fmt.Println(array1, slice1)

	slice2 := make([]int, 2)
	/*
		Dessa forma a slice 2 como tem apenas dois posições,
		na hora de copiar todos os dados da slice 1
		que tem três posições, copia somente os dois primeiros
		posições da slice 1 para a slice 2
	*/
	copy(slice2, slice1)
	fmt.Println(slice2)

}
