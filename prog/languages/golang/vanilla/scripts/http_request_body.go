package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
)

type hello struct {
	Message string `json:"message"`
}

func init() {
	http.HandleFunc("GET /hello", func(res http.ResponseWriter, req *http.Request) {
		message := "hello"

		if req.URL.Query().Has("name") {
			message += " " + req.URL.Query().Get("name")
		}

		payload := map[string]any{
			"message": message,
		}

		bytes, err := json.Marshal(payload)
		if err != nil {
			log.Println("error in json marshal:", err)
			res.WriteHeader(http.StatusInternalServerError)
			return
		}

		if _, err := res.Write(bytes); err != nil {
			log.Println("error in response write:", err)
			res.WriteHeader(http.StatusInternalServerError)
			return
		}
	})

	go func() {
		log.Fatalln(
			"error in start http server:",
			http.ListenAndServe(":9999", nil),
		)
	}()
}

// go playground: https://go.dev/play/p/9Rgo-EPfT-1
func main() {
	var hello hello

	res, err := new(http.Client).Do(
		&http.Request{
			Method: http.MethodGet,
			Header: http.Header{
				"accept": []string{"application/json"},
			},
			URL: &url.URL{
				Scheme:   "http",
				Host:     "localhost:9999",
				Path:     "/hello",
				RawQuery: "name=gopher",
			},
		},
	)
	if err != nil {
		log.Println("error in http request:", err)
		return
	}
	if res.StatusCode >= http.StatusBadRequest {
		log.Println("error in http status:", res.Status)
		return
	}

	// // using: io.ReadAll() > json.Unmarshal
	// body, err := io.ReadAll(res.Body)
	// if err != nil {
	// 	log.Println("error in io readall:", err)
	// 	return
	// }
	// if err := json.Unmarshal(body, &hello); err != nil {
	// 	log.Println("error in json unmarshal:", err)
	// 	return
	// }

	// using: json.Decode()
	if err := json.NewDecoder(res.Body).Decode(&hello); err != nil {
		log.Println("error in json decode:", err)
		return
	}

	println(hello.Message)
}
