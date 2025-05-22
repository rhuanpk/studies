package main

import (
	"errors"
	"log"
	"net/http"
)

const url = `<url>`

func simple() {
	res, err := http.Get(url)
	if err != nil {
		log.Println("error in http get:", err)
		return
	}
	defer res.Body.Close()

	log.Println("response:", res)
}

func basic() {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Println("error in http new request:", err)
	}

	res, err := new(http.Client).Do(request)
	if err != nil {
		log.Println("error in client do:", err)
	}

	log.Println("response:", res)
}

func enhanced() {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			log.Println("redirect to:", req.URL)
			return map[bool]error{true: errors.New("many redirects")}[len(via) >= 2]
		},
	}

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Println("error in http new request:", err)
		return
	}

	request.Header.Set("user-agent", `<user-agent>`)

	res, err := client.Do(request)
	if err != nil {
		log.Println("error in client do:", err)
		return
	}
	defer res.Body.Close()

	log.Println("response:", res)
}

func main() {
	simple()
	basic()
	enhanced()
}
