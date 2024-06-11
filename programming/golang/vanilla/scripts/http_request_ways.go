package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
)

func main() {
	var (
		request  *http.Request
		response *http.Response
		URL      url.URL
		url      string
		data     map[string]any
		payload  []byte
		body     *bytes.Buffer
		err      error
	)

	// set url variable
	url = URL.String() // struct
	url = `url`        // raw

	// set body variable
	data = map[string]any{"field": "value"}           // mapped
	payload, err = json.Marshal(data)                 // mapped
	body = bytes.NewBuffer(payload)                   // mapped
	body = bytes.NewBufferString(`{"field":"value"}`) // raw

	// requst direct function
	response, err = http.Get(url)                        // get
	response, err = http.Post(url, "content/type", body) // post

	// request new client
	request, err = http.NewRequest(http.MethodGet, url, nil)   // get
	request, err = http.NewRequest(http.MethodPost, url, body) // post
	response, err = new(http.Client).Do(request)               // do

	// check error and close defer
	if err != nil {
		log.Fatalln(err)
	}
	defer response.Body.Close()

	// read response body
	payload, err = io.ReadAll(response.Body)           // json
	err = json.Unmarshal(payload, &data)               // json
	err = json.NewDecoder(response.Body).Decode(&data) // direct
}
