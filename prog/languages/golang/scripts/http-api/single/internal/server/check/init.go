package check

import (
	"net/http"

	"dev/pkg/router"
)

func init() {
	router.AddScheme(
		router.Scheme{
			Method:  http.MethodGet,
			Routes:  []string{"/check"},
			Handler: check,
		},
	)
}
