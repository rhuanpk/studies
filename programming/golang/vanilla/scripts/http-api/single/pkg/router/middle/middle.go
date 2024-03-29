package middle

import "net/http"

// Middleware is the middleware primary type.
type Middleware func(http.HandlerFunc) http.HandlerFunc
