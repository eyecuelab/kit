package server

import (
	"fmt"

	valid "github.com/asaskevich/govalidator"
	"github.com/eyecuelab/kit/web"
	"github.com/facebookgo/grace/gracehttp"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/viper"
)

type apiValidator struct{}

func (v *apiValidator) Validate(i interface{}) error {
	_, err := valid.ValidateStruct(i)
	return err
}

var (
	Echo *echo.Echo
	mws  = []echo.MiddlewareFunc{}
	host string
)

func NewEcho(port int) *echo.Echo {
	e := echo.New()
	e.Validator = &apiValidator{}
	e.Server.Addr = fmt.Sprintf(":%v", port)
	e.HTTPErrorHandler = web.ErrorHandler

	e.Use(web.ApiContextMiddleWare())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		Skipper:    web.AuthedSkipper(),
		SigningKey: []byte(viper.GetString("secret")),
	}))
	e.Use(mws...)

	return e
}

func Start(port int, h string) {
	Echo = NewEcho(port)
	web.InitRoutes(Echo)
	host = h

	Echo.Logger.Fatal(gracehttp.Serve(Echo.Server))
}

func URI(routeName string, args ...interface{}) (string, error) {
	path := Echo.Reverse(routeName, args...)
	if path == "" {
		return "", fmt.Errorf("Cannot form URI, route name '%v' not found", routeName)
	}

	if host == "" {
		return "", fmt.Errorf("Cannot form URI, host not set. (use --host or env to set one)")
	}

	return host + path, nil
}

func AddValidator(name string, f valid.CustomTypeValidator) {
	valid.CustomTypeTagMap.Set("score", f)
}

func AddMiddleWare(mw echo.MiddlewareFunc) {
	mws = append(mws, mw)
}
