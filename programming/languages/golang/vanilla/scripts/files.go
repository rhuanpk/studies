//go:build files

package main

import "os"

func main() {
	file, err := os.Create("/tmp/test.tmp")
	// file, err := os.Create("/tmp/test.tmp/")
	defer file.Close()
	if err != nil {
		println("Error on create file!")
		println("Error message:", err.Error())
		println("Leaving without create file!")
		os.Exit(1)
	} else {
		var array = []string{"hello", "world", "!"}
		for _, value := range array {
			// file.WriteString("key:" + key)
			file.WriteString("value: " + value + "\n")
		}
	}
}
