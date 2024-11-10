package main

import (
	"encoding/json"
	"log"
	"os"
)

const end = 2

func main() {
	// open/create files
	single, err := os.OpenFile("single.json", os.O_CREATE|os.O_WRONLY, 0664)
	if err != nil {
		log.Println("error in open file:", err)
	}
	multi, err := os.OpenFile("multi.json", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0664)
	if err != nil {
		log.Println("error in open file:", err)
	}

	// write single payload
	payload := map[string]any{"foo": "bar", "xpto": 42}

	bytes, err := json.MarshalIndent(payload, "", "\t")
	if err != nil {
		log.Println("error in json marshal indent single:", err)
	}

	if _, err = single.Write(bytes); err != nil {
		log.Println("error in write single file:", err)
	}

	// write multi times
	if _, err := multi.WriteString("[\n"); err != nil {
		log.Println("error in wirte multi file:", err)
	}
	for index := 0; index < end; index++ {
		payload := map[string]any{"foo": "bar", "index": index}

		bytes, err := json.MarshalIndent(payload, "\t", "\t\t")
		if err != nil {
			log.Println("error in json marshal indent multi:", err)
		}

		var suffixes []byte
		if index <= end-2 {
			suffixes = append(suffixes, ',')
		}
		suffixes = append(suffixes, '\n')

		if _, err = multi.Write(append(append([]byte{'\t'}, bytes...), suffixes...)); err != nil {
			log.Println("error in write multi file:", err)
		}
	}
	if _, err := multi.WriteString("]\n"); err != nil {
		log.Println("error in wirte multi file:", err)
	}
}
