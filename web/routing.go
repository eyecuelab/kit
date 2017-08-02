package web

import (
	"github.com/labstack/echo"
)

type MethodHandler struct {
	Method     string
	Handler    echo.HandlerFunc
	MiddleWare []echo.MiddlewareFunc
}

type Route struct {
	Path     string
	Handlers []*MethodHandler
}

type RouteConfig struct {
	Routes     []*Route
	MiddleWare []echo.MiddlewareFunc
}

func (route *Route) Handle(m string, hf echo.HandlerFunc, mw ...echo.MiddlewareFunc) *Route {
	handler := &MethodHandler{m, hf, mw}
	route.Handlers = append(route.Handlers, handler)
	return route
}

func (config *RouteConfig) AddRoute(path string) *Route {
	route := &Route{path, []*MethodHandler{}}
	config.Routes = append(config.Routes, route)
	return route
}

var Routing *RouteConfig

func init() {
	Routing = &RouteConfig{[]*Route{}, []echo.MiddlewareFunc{}}
}

func AddRoute(path string) *Route {
	return Routing.AddRoute(path)
}
