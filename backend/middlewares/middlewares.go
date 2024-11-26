package middlewares

import "net/http"

// ApplyMiddlewares chains multiple middlewares
func ApplyMiddlewares(handler http.Handler, middlewares []func(http.Handler) http.Handler) http.Handler {
	for _, mw := range middlewares {
		handler = mw(handler)
	}
	return handler
}
