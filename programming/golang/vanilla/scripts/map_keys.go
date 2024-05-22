package main

import "fmt"

// go playground: https://go.dev/play/p/FneMYkqyPgJ
func main() {
	stringMap := map[any]string{
		"key1": "xpto",
	}
	intMap := map[any]int{
		"key1": 42,
	}
	floatMap := map[any]float32{
		"key1": 4.2,
	}
	boolMap := map[any]bool{
		"key1": true,
	}
	anyMap := map[any]any{
		"key1": struct{}{},
	}

	println("##### println #####")

	println("\n--- string ---")
	fmt.Println("key1:", stringMap["key1"])
	fmt.Println("key2:", stringMap["key2"])

	println("--- int ---")
	fmt.Println("key1:", intMap["key1"])
	fmt.Println("key2:", intMap["key2"])

	println("--- float ---")
	fmt.Println("key1:", floatMap["key1"])
	fmt.Println("key2:", floatMap["key2"])

	println("--- bool ---")
	fmt.Println("key1:", boolMap["key1"])
	fmt.Println("key2:", boolMap["key2"])

	println("--- any ---")
	fmt.Println("key1:", anyMap["key1"])
	fmt.Println("key2:", anyMap["key2"])

	println("\n##### if #####")

	println("\n--- bool ---")
	if boolMap["key1"] {
		fmt.Println("key1 exists")
	} else {
		fmt.Println("key1 not exists")
	}
	if boolMap["key2"] {
		fmt.Println("key2 exists")
	} else {
		fmt.Println("key2 not exists")
	}
}
