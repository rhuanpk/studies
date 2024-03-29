package router

import (
	"net/http"

	"dev/pkg/router/middle"
)

// Scheme struct to define endpoints, e.g.:
//
//	scheme := Scheme{
//		Method:      http.MethodGet,
//		Route:       "/shallow/{param}/deep/{path}",
//		Handler:     func(w http.ResponseWriter, r *http.Request) {},
//		Middlewares: []middle.Middleware{},
//	}
type Scheme struct {
	Method      string
	Route       string
	Handler     http.HandlerFunc
	Middlewares []middle.Middleware
}

// Parse return the Method and Route for http.HandleFunc() to set.
func (s Scheme) Parse() string {
	return s.Method + " " + s.Route
}

// Schemes is the global variable to add scheme routers.
var Schemes []Scheme

// AddScheme add a new router scheme to Schemes global variable.
func AddScheme(scheme Scheme) {
	Schemes = append(Schemes, scheme)
}
