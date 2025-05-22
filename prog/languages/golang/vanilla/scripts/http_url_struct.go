package main

import (
	"log"
	"net/url"
)

func params(URL *url.URL) {
	//URL, _ := url.Parse(`http://localhost:9999/test?hello=world&xpto=foo&xpto=bar`)
	//println(URL.String())

	// get query values
	query := URL.Query()

	// set some params
	query.Set("hello", "world")
	query.Set("xpto", "foo")
	query.Add("xpto", "bar")
	URL.RawQuery = query.Encode()
	println(URL.String())

	// xpto param is overrided
	query.Set("xpto", "test")
	URL.RawQuery = query.Encode()
	println(URL.String())
}

func structured() {
	// set URL struct
	URL := url.URL{
		Scheme: "http",
		Host:   "localhost:9999",
		Path:   "/test",
	}

	println("structured url")
	params(&URL)
}

func parsed() {
	// get URL struct from already exist raw URL
	//URL, err := url.Parse(`http://localhost:9z9z/test`)
	URL, err := url.Parse(`http://localhost:9999/test`)
	if err != nil {
		log.Fatalln("error in parse url:", err)
	}

	println("parsed url")
	params(URL)
}

func main() {
	structured()
	parsed()
}
