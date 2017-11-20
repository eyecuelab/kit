package web

import (
	"fmt"
	"testing"

	"github.com/eyecuelab/kit/pretty"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestRoute_Handle(t *testing.T) {
	const (
		m    = "f"
		path = "http://localhost.com"
		name = "foo"
	)

	noOpMiddleWare := func(hf echo.HandlerFunc) echo.HandlerFunc { return hf }
	alwaysNilMiddleWare := func(hf echo.HandlerFunc) echo.HandlerFunc { return func(echo.Context) error { return nil } }
	hf := func(ApiContext) error { return nil }
	wantHF := wrapApiRoute(hf)
	rt := Route{Path: path, Name: name}

	wantMiddleware := []echo.MiddlewareFunc{noOpMiddleWare, alwaysNilMiddleWare}
	wantHandler := MethodHandler{m, wantHF, wantMiddleware}
	wantHandlers := []*MethodHandler{&wantHandler}

	got := rt.Handle(m, hf, noOpMiddleWare, alwaysNilMiddleWare)
	want := Route{Path: path, Name: name, Handlers: wantHandlers}
	assert.Empty(t, pretty.Diff(want, *got), 0, fmt.Sprint(pretty.Diff(got, want)))
}

func TestRoute_SetName(t *testing.T) {
	rt := new(Route)
	rt.SetName("foo")
	assert.Equal(t, rt.Name, "foo")
}

func TestRouteConfig_AddRoute(t *testing.T) {
	const path = "http://somepath.com"
	cfg := new(RouteConfig)
	rt := cfg.AddRoute(path)
	assert.Equal(t, *rt, Route{Path: path, Handlers: []*MethodHandler{}})
	assert.Equal(t, cfg.Routes, []*Route{rt})
}

func Test_routeByName(t *testing.T) {
	const name = "foo"
	cfg := new(RouteConfig)
	wantRoute := Route{
		Name: name,
	}

	cfg.Routes = []*Route{&Route{}, &wantRoute, &Route{Name: "bar"}}
	got, err := routeByName(cfg, name)
	assert.Equal(t, got, &wantRoute)
	assert.NoError(t, err)

	cfg = &RouteConfig{Routes: []*Route{&Route{}, &Route{Name: "bar"}}}
	_, err = routeByName(cfg, name)
	assert.Error(t, err)

}

func Test_wrapApiRoute(t *testing.T) {

}

func Test_initRoutes(t *testing.T) {
	const (
		path   = "http://localpath.com/foo"
		name   = "foo"
		method = "m"
	)
	hf := func(ApiContext) error { return nil }
	e := echo.New()
	cfg := new(RouteConfig)
	rt := cfg.AddRoute(path).SetName(name).Handle(method, hf)
	initRoutes(e, cfg)

	wantRoutes := []*Route{rt}
	assert.Equal(t, cfg.Routes, wantRoutes)
	assert.Equal(t, cfg.Routes[0].ERoute.Name, rt.Name)

	wantEchoRoute := echo.Route{Method: method, Path: path, Name: name}
	assert.Len(t, e.Routes(), 1)
	assert.Equal(t, wantEchoRoute, *e.Routes()[0])

}
