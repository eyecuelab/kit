package web

import (
	"fmt"
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
	Name     string
	ERoute   *echo.Route
}

type RouteConfig struct {
	Routes     []*Route
	MiddleWare []echo.MiddlewareFunc
}

func (route *Route) Handle(m string, hf HandlerFunc, mw ...echo.MiddlewareFunc) *Route {
	handler := &MethodHandler{m, wrapApiRoute(hf), mw}
	route.Handlers = append(route.Handlers, handler)
	return route
}

func (route *Route) SetName(n string) *Route {
	route.Name = n
	return route
}

func (config *RouteConfig) AddRoute(path string) *Route {
	route := &Route{
		Path:     path,
		Handlers: []*MethodHandler{},
	}
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

func RouteByName(name string) (*Route, error) {
	for _, r := range Routing.Routes {
		if r.Name == name {
			return r, nil
		}
	}
	return nil, fmt.Errorf("Could not find route with name %v", name)
}

type HandlerFunc func(ApiContext) error

func wrapApiRoute(f HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ac := c.(ApiContext)
		return f(ac)
	}
}

func InitRoutes(e *echo.Echo) {
	for _, route := range Routing.Routes {
		for _, handler := range route.Handlers {
			er := e.Add(handler.Method, route.Path, handler.Handler, handler.MiddleWare...)
			er.Name = route.Name
			route.ERoute = er
		}
	}
}
