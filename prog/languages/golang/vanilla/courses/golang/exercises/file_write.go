//go:build file_write

/*
	Create some program that:
		- write a list of cities to a new file.

	Source: https://golangr.com/write-file/
*/

package main

import (
	"os"
	"strconv"

	std "github.com/rhuan-pk/pkutils/go/standardutils"
)

func main() {

	file, err := os.Create("/tmp/tmp/go_file.tmp")
	if err != nil {
		std.ErrorMessage("Error on file creating!", err)
	}
	defer file.Close()

	cities := []string{"cianorte", "maringa", "japura", "ribeirao_preto"}
	for key, value := range cities {
		if _, err := file.WriteString("[" + strconv.Itoa(key) + "] " + value + "\n"); err != nil {
			std.ErrorMessage("Error on file writing!", err)
		}
	}

	// for {
	// 	print("Input (\"quit\" to quit): ")
	// 	if input := std.StringReader(); "quit\n" == strings.ToLower(input) {
	// 		println("Nothing more to write, thx!")
	// 		break
	// 	} else {
	// 		if _, err := file.WriteString(input); err != nil {
	// 			std.ErrorMessage("Error on file writing!", err)
	// 		}
	// 	}
	// }

}
