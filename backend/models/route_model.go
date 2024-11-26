package models

import "net/http"

// RouteGroup represents a collection of routes under a common base path
type RouteGroup struct {
	BasePath    string
	Middlewares []func(http.Handler) http.Handler
	Routes      []Route
	SubGroups   []RouteGroup
}

// Route represents an individual route with method, handler, and middlewares
type Route struct {
	Method      string
	Path        string
	Handler     func(http.ResponseWriter, *http.Request) error
	Middlewares []func(http.Handler) http.Handler
}
