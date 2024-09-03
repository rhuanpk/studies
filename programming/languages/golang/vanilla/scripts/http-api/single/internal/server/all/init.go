package all

import (
	"net/http"

	"dev/pkg/router"
)

func init() {
	// Any other that not match with the below's:
	//
	// GET /<any>
	// GET /<any>/<any>
	// POST /xpto/<any>
	router.AddScheme(
		router.Scheme{
			Routes: []string{
				"/",
				"/{_}/{param}",
			},
			Handler: all,
		},
	)

	// Only this:
	//
	// POST /<any>
	router.AddScheme(
		router.Scheme{
			Method:  http.MethodPost,
			Routes:  []string{"/{param}"},
			Handler: all,
		},
	)

	// Only this:
	//
	// GET /xpto/<any>
	router.AddScheme(
		router.Scheme{
			Method:  http.MethodGet,
			Routes:  []string{"/xpto/{param}"},
			Handler: all,
		},
	)
}
