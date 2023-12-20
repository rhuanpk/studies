package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	url := "http://localhost:8888/foo"
	// body := []byte(`{"patient":0,"boo":"baz"}`)
	body, err := json.Marshal(
		map[string]any{
			"patient": 0,
			"boo":     "baz",
		},
	)
	if err != nil {
		log.Println("error in marshal body:", err)
	}

	// response, err := http.Get(url)
	// if err != nil {
	// 	log.Fatalln("error in request:", err)
	// 	return
	// }
	// defer response.Body.Close()
	request, err := http.NewRequest(http.MethodGet, url, bytes.NewBuffer(body))
	if err != nil {
		log.Fatalln("error on create request:", err)
		return
	}

	request.Header.Add("bar", "bar")
	request.Header.Set("content-type", "application/json")

	client := new(http.Client)
	response, err := client.Do(request)
	if err != nil {
		log.Fatalln("error in request:", err)
		return
	}
	defer response.Body.Close()

	body, err = io.ReadAll(response.Body)
	if err != nil {
		log.Fatalln("error on read response body:", body)
		return
	}

	println(">>> Golang Client HTTP With Standard Library!")
	fmt.Printf("response status: %s\nresponse body: %v\n", response.Status, string(body))
}
