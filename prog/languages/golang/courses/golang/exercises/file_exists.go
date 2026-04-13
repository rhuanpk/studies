//go:build file_exists

/*
	Create some program that:
		- check if a file exists on your local disk.
		- can you check if a file exists on an external disk?
			R: yes.

	Source: https://golangr.com/file-exists/
*/

package main

import (
	"bufio"
	"os"
	"strings"
)

func main() {
	read := bufio.NewReader(os.Stdin)
	print("enter with the path file: ")
	if file, err := read.ReadString('\n'); err != nil {
		println("input error!\nleaving the program!")
	} else {
		if _, err := os.Stat(strings.Replace(file, "\n", "", -1)); os.IsNotExist(err) {
			println("file not found!")
		} else {
			println("file found!")
		}
	}
}
