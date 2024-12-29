package main

import "fmt"

// go playground: https://go.dev/play/p/JUDPuoc8sxX
func main() {
	// create a empty map
	stuff := make(map[[2]any]any)

	// inserting values
	stuff[[2]any{1, "a"}] = "first value"
	stuff[[2]any{2.5, true}] = 42
	stuff[[2]any{"key", 3}] = []string{"slice", "of", "strings"}

	// iterate over map
	println("stuff map:")
	for key, value := range stuff {
		fmt.Printf("key: %v, value: %v\n", key, value)
	}

	// check if key exists
	println("\ncheck if key exists:")
	key := [2]any{1, "a"}
	if value, ok := stuff[key]; ok {
		fmt.Printf("found value for the key %v: %v\n", key, value)
	}

	// removing element of map
	delete(stuff, [2]any{2.5, true})

	// iterate over map after removal
	println("\nmap after removal:")
	for key, value := range stuff {
		fmt.Printf("key: %v, value: %v\n", key, value)
	}

	// example: registering coordinates with metadata
	coordinates := map[[2]any]any{
		{10.5, 20.3}: "point A",

		{15.7, 30.1}: struct {
			nome string
			tipo string
		}{"point B", "reference"},

		{5.2, 8.9}: map[string]int{
			"population": 1000,
			"altitude":   50,
		},
	}

	// iterate over coordinates
	println("\ncoordinates:")
	for coordinate, metadata := range coordinates {
		fmt.Printf("cordinate %v: %v\n", coordinate, metadata)
	}
}
