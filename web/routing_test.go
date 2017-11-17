package web

import (
	"fmt"
	"reflect"
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

func TestAddRoute(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want *Route
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddRoute(tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddRoute() = %v, want %v", got, tt.want)
			}
		})
	}
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
	type args struct {
		f HandlerFunc
	}
	tests := []struct {
		name string
		args args
		want echo.HandlerFunc
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wrapApiRoute(tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("wrapApiRoute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInitRoutes(t *testing.T) {
	type args struct {
		e *echo.Echo
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InitRoutes(tt.args.e)
		})
	}
}
