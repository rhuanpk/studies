package middles

import (
	"net/http"

	"dev/pkg/router"
	"dev/pkg/router/middle"
)

func init() {
	router.AddScheme(
		router.Scheme{
			Method:  http.MethodGet,
			Routes:  []string{"/middles"},
			Handler: middles,
			Middlewares: []middle.Middleware{
				middle.Foo,
				middle.Bar,
			},
		},
	)
}
