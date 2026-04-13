package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func logFunctionCalled(path string) {
	log.Println("\""+path+"\"", "was called")
}

func foo(w http.ResponseWriter, r *http.Request) {
	logFunctionCalled(r.URL.Path)

	w.Header().Add("foo", "foo")
	w.Header().Set("content-type", "application/json")

	var data map[string]any
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte("error on data deserialization"))
		log.Println("error on body decode:", err)
		return
	}
	log.Println("decoded body:", data)

	err = json.NewEncoder(w).Encode(map[string]any{"body": data, "message": "thanks"})
	if err != nil {
		http.Error(
			w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
		)
		log.Println("error on body encode:", err)
	}
}

/*
func bar(w http.ResponseWriter, r *http.Request) {
	logFunctionCalled(r.URL.Path)

	w.Header().Add("foo", "foo")
	w.Header().Set("content-type", "application/json")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "erro on read request body", http.StatusUnprocessableEntity)
		log.Println("error on read request body:", err)
		return
	}
	r.Body.Close()
	// log.Println("request body:", string(body))

	var data map[string]any
	json.NewDecoder(bytes.NewReader(body)).Decode(&data)
	// marshaled, err := json.MarshalIndent(data, "", "\t")
	marshaled, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("error on indent request body:", err)
		return
	}
	// log.Println("marshaled body:", string(marshaled))
	fmt.Fprintf(w, `{"body":%v,"foo":"foo"}`, string(marshaled))
}
*/

func root(w http.ResponseWriter, r *http.Request) {
	logFunctionCalled(r.URL.Path)
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte(http.StatusText(http.StatusNotImplemented)))
}

func main() {
	println(">>> Golang HTTP API With Standard Library!")
	http.HandleFunc("/", root)
	http.HandleFunc("/foo", foo)
	// http.HandleFunc("/bar", bar)
	log.Fatalln(http.ListenAndServe(":8888", nil))
}
