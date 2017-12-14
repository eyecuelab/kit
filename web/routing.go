package web

import (
	"fmt"
	"strings"

	"github.com/labstack/echo"
)

type namePrefix string
type routePrefix string

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
	Routes      []*Route
	MiddleWare  []echo.MiddlewareFunc
	RPrefix     routePrefix
	NPrefix     namePrefix
}

func (config *RouteConfig) AddNameLevel(level string) string {
	lvl := namePrefix(level)
	if config.NPrefix == "" {
		config.NPrefix = lvl
	} else {
		config.NPrefix = config.NPrefix + "." + lvl
	}
	return string(config.NPrefix)
}

func (config *RouteConfig) AddRouteLevel(level string) string {
	lvl := "/" + routePrefix(level)
	config.RPrefix = config.RPrefix + lvl
	return string(config.RPrefix)
}

func (config *RouteConfig) RemoveLastNameLevel() string {
	p := string(config.NPrefix)
	lastDotIndex := strings.LastIndex(p, ".")
	p = p[:lastDotIndex]
	config.NPrefix = namePrefix(p)
	return p
}

func (p routePrefix) with(path string) string {
	fullPath := "/" + path
	if p == "" {
		return fullPath
	}
	return string(p) + fullPath
}


func (p routePrefix) RemoveLastLevel() *routePrefix {
	lastDotIndex := strings.LastIndex(string(p), "/")
	p = p[:lastDotIndex]
	return &p
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
		Path:     config.RPrefix.with(path),
		Handlers: []*MethodHandler{},
	}
	config.Routes = append(config.Routes, route)
	return route
}

func newRouteConfig() *RouteConfig {
	return &RouteConfig{[]*Route{}, []echo.MiddlewareFunc{}, "", ""}
}

var Routing *RouteConfig

func init() {
	Routing = &RouteConfig{[]*Route{}, []echo.MiddlewareFunc{}, "", ""}
}

func AddRoute(path string) *Route {
	return Routing.AddRoute(path)
}
func routeByName(cfg *RouteConfig, name string) (*Route, error) {
	for _, r := range cfg.Routes {
		if r.Name == name {
			return r, nil
		}
	}
	return nil, fmt.Errorf("Could not find route with name %v", name)
}

func RouteByName(name string) (*Route, error) {
	return routeByName(Routing, name)
}

type HandlerFunc func(ApiContext) error

func wrapApiRoute(f HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return f(c.(ApiContext))
	}
}

func InitRoutes(e *echo.Echo) {
	initRoutes(e, Routing)
}

func initRoutes(e *echo.Echo, cfg *RouteConfig) {
	for _, route := range cfg.Routes {
		initRoute(e, route)
	}
}

func initRoute(e *echo.Echo, route *Route) {
	for _, handler := range route.Handlers {
		er := e.Add(handler.Method, route.Path, handler.Handler, handler.MiddleWare...)
		er.Name = route.Name
		fmt.Println(route.Name+" " + route.Path)
		route.ERoute = er
	}
}