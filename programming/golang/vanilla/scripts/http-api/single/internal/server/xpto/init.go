package xpto

import (
	"net/http"

	"dev/pkg/router"
)

func init() {
	router.AddScheme(
		router.Scheme{
			Method:  http.MethodGet,
			Route:   "/xpto/{param}",
			Handler: xpto,
		},
	)
}
