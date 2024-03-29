package test

import (
	"net/http"

	"dev/pkg/router"
	"dev/pkg/router/middle"
)

func init() {
	router.AddScheme(
		router.Scheme{
			Method:  http.MethodGet,
			Route:   "/test",
			Handler: test,
			Middlewares: []middle.Middleware{
				middle.Foo,
				middle.Bar,
			},
		},
	)
}
