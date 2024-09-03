//go:build file_read

/*
	Create some program that:
		- think of when you’d read a file ‘line by line’ vs ‘at once’?
			R: yes.
		- create a new file containing names and read it into an array

	Source: https://golangr.com/read-file/
*/

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	std "github.com/rhuan-pk/pkutils/go/standardutils"
)

func main() {

	print("Enter with the path file: <enter to default> ")
	file_path := std.StringReader()
	var relativeFilePath string

	if file_path != "\n" {
		file_path = strings.Replace(file_path, "\n", "", -1)
		// --------------------------------------------------------
		// regex := regexp.MustCompile("(/[^/]*)$")
		// relativeFilePath = regex.ReplaceAllString(file_path, "")
		// --------------------------------------------------------
		// dir, file := path.Split(file_path)
		dir, _ := path.Split(file_path)
		relativeFilePath = strings.TrimSuffix(dir, "/")
		// --------------------------------------------------------
	} else {
		var err error
		file_path = "go_test_file.txt"
		relativeFilePath, err = os.Getwd()
		if err != nil {
			std.ErrorMessage("Error on get current directory!", err)
		}
	}

	file, err := os.Create(file_path)
	defer file.Close()
	// defer os.Remove(file_path)
	if err != nil {
		std.ErrorMessage("Error on file creating!", err)
	} else {
		fmt.Printf("Success on file creating on \"%s\"!\n", relativeFilePath)
	}

	for {
		print("Enter with the data (type \"quit\" to quit): ")
		if input := std.StringReader(); strings.ToLower(input) == "quit\n" {
			println("Nothing more to write!")
			break
		} else {
			if _, err := file.WriteString(input); err != nil {
				std.ErrorMessage("Error on file writing!", err)
			}
		}
	}

	if file_content, err := ioutil.ReadFile(file_path); err != nil {
		std.ErrorMessage("Error on file reading!", err)
	} else {
		println("Printing the file at once:")
		print(string(file_content))
	}

	if file_lines, err := std.FileLineReader(file_path); err != "" {
		std.ErrorMessage("Error on file reading!", err)
	} else {
		println("Printing the file line by line:")
		for line, content := range file_lines {
			fmt.Printf("%d) %s\n", (line + 1), content)
		}
	}

}
