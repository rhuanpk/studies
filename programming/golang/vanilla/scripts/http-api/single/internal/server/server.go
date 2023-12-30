package server

import (
	"encoding/json"
	"net/http"
	"regexp"

	"backend/pkg/logger"
)

// Handler is the endpoint handler.
func Handler(response http.ResponseWriter, request *http.Request) {
	// set variables
	log := logger.NewLogger(logger.Server)
	endpoint := regexp.MustCompile(`^/[^/]+`).FindString(request.URL.Path)
	param := regexp.MustCompile(`[^/]+`).FindAllString(request.URL.Path, -1)

	// trace logs
	log.Short()
	log.Printf(
		`"%s" was called`,
		map[bool]string{
			true:  request.URL.Path,
			false: endpoint,
		}[endpoint == ""],
	)

	// log request url
	log.Println("request url:")
	log.Println("\t" + request.URL.String())

	// get request headers
	if len(request.Header) > 0 {
		log.Println("request headers:")
		for key, value := range request.Header {
			log.Printf("\tkey: %v", key)
			log.Printf("\tvalue: %v", value)
		}
	}

	// get path param
	if len(param) > 1 {
		log.Println("request path param:")
		log.Println("\t" + param[1])
	}

	// get query params
	if len(request.URL.Query()) > 0 {
		log.Println("request query params:")
		for key, value := range request.URL.Query() {
			log.Printf("\tkey: %v", key)
			log.Printf("\tvalue: %v", value)
		}
	}

	// get request body
	var data map[string]any
	err := json.NewDecoder(request.Body).Decode(&data)
	if err != nil {
		//var errEOF *json.InvalidUnmarshalError
		//if errors.As(err, &errEOF) {
		if err.Error() != "EOF" {
			http.Error(
				response,
				"error in data deserialization",
				http.StatusUnprocessableEntity,
			)
			log.Println("error in body decode:", err)
			return
		}
		log.Println("warning in body decode:")
		log.Println("\t" + err.Error())
	}

	// pretty print body
	indented, err := json.MarshalIndent(data, log.FullPrefix(), "\t")
	if err != nil {
		http.Error(
			response,
			"error in data indent",
			http.StatusUnprocessableEntity,
		)
		log.Println("error in body indent:", err)
		return
	}
	converted := string(indented)

	log.Println("request decoded body:")
	if converted == "null" {
		converted = "\t" + converted
	}
	log.Println(converted)

	// set reply headers
	response.Header().Add("server", "header")
	response.Header().Set("content-type", "application/json")

	// do reply
	err = json.NewEncoder(response).Encode(
		map[string]any{
			"body":   data,
			"server": "body",
		},
	)
	if err != nil {
		http.Error(
			response,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
		)
		log.Println("error in body encode:", err)
	}
}
