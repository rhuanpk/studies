package router

import (
	"net/http"

	"dev/pkg/router/middle"
)

// Scheme struct to define endpoints e.g.:
//
//	scheme := Scheme{
//		Method:      http.MethodGet,
//		Route:       "/shallow/{param}/deep/{path}",
//		Handler:     func(w http.ResponseWriter, r *http.Request) {},
//		Middlewares: []middle.Middleware{},
//	}
type Scheme struct {
	Method      string
	Routes      []string
	Handler     http.HandlerFunc
	Middlewares []middle.Middleware
}

// Parse return the Method and Route for http.HandleFunc() to set.
// If index is less then 0 or greater then len(s) returns empty string.
// E.g.:
//
//	"METHOD /foo/{bar}"
func (s Scheme) Parse(index int) string {
	if index < 0 || index >= len(s.Routes) {
		return ""
	}
	return s.Method + " " + s.Routes[index]
}

// Parses iterates over s.Routes returning a slice of Parse() returns.
func (s Scheme) Parses() (parses []string) {
	for _, route := range s.Routes {
		parses = append(parses, s.Method+" "+route)
	}
	return
}

// Schemes is the global variable to add scheme routers.
var Schemes []Scheme

// AddScheme add a new router scheme to Schemes global variable.
func AddScheme(scheme Scheme) {
	Schemes = append(Schemes, scheme)
}
