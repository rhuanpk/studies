package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type object struct {
	String               string  `json:"string"`
	StringOmitted        string  `json:"stringOmitted,omitempty"`
	StringPointer        *string `json:"stringPointer"`
	StringPointerOmitted *string `json:"stringPointerOmitted,omitempty"`

	Int               int  `json:"int"`
	IntOmitted        int  `json:"intOmitted,omitempty"`
	IntPointer        *int `json:"intPointer"`
	IntPointerOmitted *int `json:"intPointerOmitted,omitempty"`

	Float64               float64  `json:"float64"`
	Float64Omitted        float64  `json:"float64Omitted,omitempty"`
	Float64Pointer        *float64 `json:"float64Pointer"`
	Float64PointerOmitted *float64 `json:"float64PointerOmitted,omitempty"`

	Bool               bool  `json:"bool"`
	BoolOmitted        bool  `json:"boolOmitted,omitempty"`
	BoolPointer        *bool `json:"boolPointer"`
	BoolPointerOmitted *bool `json:"boolPointerOmitted,omitempty"`

	Any               any  `json:"any"`
	AnyOmitted        any  `json:"anyOmitted,omitempty"`
	AnyPointer        *any `json:"anyPointer"`
	AnyPointerOmitted *any `json:"anyPointerOmitted,omitempty"`

	Struct               struct{}  `json:"struct"`
	StructOmitted        struct{}  `json:"structOmitted,omitempty"`
	StructPointer        *struct{} `json:"structPointer"`
	StructPointerOmitted *struct{} `json:"structPointerOmitted,omitempty"`
}

func main() {
	// endpoint: GET localhost:9999/
	// response:
	// {
	//     "string": "",
	//     "stringPointer": null,
	//     "int": 0,
	//     "intPointer": null,
	//     "float64": 0,
	//     "float64Pointer": null,
	//     "bool": false,
	//     "boolPointer": null,
	//     "any": null,
	//     "anyPointer": null,
	//     "struct": {},
	//     "structOmitted": {},
	//     "structPointer": null
	// }
	http.HandleFunc("GET /", func(res http.ResponseWriter, req *http.Request) {
		var object object

		payload, err := json.Marshal(object)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			res.Header().Set("content-type", "text/plain")
			res.Write([]byte(err.Error()))
			return
		}

		res.Header().Set("content-type", "application/json")
		if _, err := res.Write(payload); err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			res.Header().Set("content-type", "text/plain")
			res.Write([]byte(err.Error()))
			return
		}
	})

	// endpoint: POST localhost:9999/
	// body: {}
	// output:
	// {
	//         "string": "",
	//         "stringPointer": null,
	//         "int": 0,
	//         "intPointer": null,
	//         "float64": 0,
	//         "float64Pointer": null,
	//         "bool": false,
	//         "boolPointer": null,
	//         "any": null,
	//         "anyPointer": null,
	//         "struct": {},
	//         "structOmitted": {},
	//         "structPointer": null
	// }
	http.HandleFunc("POST /", func(res http.ResponseWriter, req *http.Request) {
		var object object

		if err := json.NewDecoder(req.Body).Decode(&object); err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			res.Header().Set("content-type", "text/plain")
			res.Write([]byte(err.Error()))
			return
		}
		defer req.Body.Close()

		payload, err := json.MarshalIndent(object, "", "\t")
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			res.Header().Set("content-type", "text/plain")
			res.Write([]byte(err.Error()))
			return
		}

		println(string(payload))
	})

	log.Fatalln("error in start http server:", http.ListenAndServe(":9999", nil))
}
