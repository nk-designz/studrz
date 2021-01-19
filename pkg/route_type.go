package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
)

// MethodFunc registers a handler to a path in echo
type MethodFunc func(string, echo.HandlerFunc, ...echo.MiddlewareFunc) *echo.Route

// Route for connecting path to handler
type Route struct {
	Path    string
	Method  MethodFunc
	Handler echo.HandlerFunc
}


// CreateRoute reutns a route
func CreateRoute(path string, method MethodFunc, handler echo.HandlerFunc) Route {
	return Route{
		Path:    path,
		Method:  method,
		Handler: handler,
	}
}

// BundleRoutes returns an array of routes from arguments
func BundleRoutes(rts ...Route) []Route {
	return rts
}

// RegisterRoutes registers array of routes in echo
func RegisterRoutes(rts []Route) {
	fmt.Println("Registered routes:")
	for i, r := range rts {
		r.Method(
			r.Path,
			r.Handler,
		)
		fmt.Print(
			fmt.Sprintf("\t%dth Route on %s\n", i, r.Path))
	}
}

