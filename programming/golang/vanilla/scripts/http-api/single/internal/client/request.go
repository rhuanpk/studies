package client

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"dev/internal/consts"
	"dev/pkg/logger"
)

// Request do the request.
func Request() {
	// set variables
	const param = "otpx"
	log := logger.NewLogger("client")

	//endpoint := consts.APIEndpoint.URL() + param + "?key1=value1&key2=value2"
	query := url.Values{}
	query.Add("key1", "value1")
	query.Add("key2", "value2")
	URL := url.URL{
		Scheme:   string(consts.ServerHost.Proto),                    // http
		Host:     consts.ServerHost.Host + consts.ServerPort.Parse(), // localhost:9999
		Path:     "xpto/" + param,                                    // /xpto/value
		RawQuery: query.Encode(),                                     // key1=value1&key2=value2
	}
	endpoint := URL.String() // http://localhost:9999/xpto/value?key1=value1&key2=value2

	// prepare payload
	payload, err := json.Marshal(
		map[string]any{
			"client": "side",
			"foo":    "bar",
			"xpto":   map[string]any{"hello": "world"},
			"array":  []any{"string", 0, true, 4.2},
		},
	)
	if err != nil {
		log.Println("err: in marshal payload:", err)
	}

	// prepare request
	request, err := http.NewRequest(http.MethodGet, endpoint, bytes.NewBuffer(payload))
	if err != nil {
		log.Println("err: in create request:", err)
		return
	}
	request.Header.Add("client", "header")
	request.Header.Set("content-type", "application/json")

	// do request
	client := new(http.Client)
	response, err := client.Do(request)
	if err != nil {
		log.Println("err: in doing request:", err)
		return
	}
	defer response.Body.Close()

	// get response body
	payload, err = io.ReadAll(response.Body)
	if err != nil {
		log.Println("err: in read response body:", err)
		return
	}

	// trace logs
	log.Short()
	log.Println(`"xpto" was requested`)
	log.Println("response status:")
	log.Println("\t" + response.Status)

	// reply logs
	log.Println("request url:")
	log.Println("\t" + endpoint)

	log.Println("response headers:")
	for key, value := range response.Header {
		log.Printf("\tkey: %v", key)
		log.Printf("\tvalue: %v", value)
	}

	log.Println("request path param:")
	log.Println("\t" + param)

	log.Println("request query params:")
	for key, value := range request.URL.Query() {
		log.Printf("\tkey: %v", key)
		log.Printf("\tvalue: %v", value)
	}

	log.Println("response body:")
	var indented bytes.Buffer
	err = json.Indent(&indented, payload, log.FullPrefix(), "\t")
	if err != nil {
		log.Print("\t" + string(payload))
		log.Println("warn: in indent response body:", err)
	} else {
		log.Print(indented.String())
	}
}
