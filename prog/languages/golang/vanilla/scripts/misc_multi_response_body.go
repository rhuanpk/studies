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
	// do the request
	response, err := http.Get(`https://httpbin.org/get`)
	if err != nil {
		log.Println("error in do request:", err)
		return
	}

	// check response status code
	if response.StatusCode >= 400 {
		log.Println("bad response status code:", response.StatusCode)
		return
	}

	// extract response body and reapply it back
	bytesBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println("error in read body:", err)
		return
	}
	response.Body = io.NopCloser(bytes.NewBuffer(bytesBody))

	// bind response body into var
	var data map[string]any
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		log.Println("error in json decode:", err)
		return
	}

	// log variables
	fmt.Println("response body:", string(bytesBody))
	fmt.Println("data variable:", data)
}
