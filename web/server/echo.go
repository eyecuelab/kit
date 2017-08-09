package server

import (
	"fmt"
	"regexp"

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
)

func jwtSkipper(skipRegex string) func(echo.Context) bool {
	skip := regexp.MustCompile(skipRegex)
	return func(c echo.Context) bool {
		return skip.MatchString(c.Request().URL.Path)
	}
}

func NewEcho(port int) *echo.Echo {
	e := echo.New()
	e.Validator = &apiValidator{}
	e.Server.Addr = fmt.Sprintf(":%v", port)
	e.Use(web.ApiContextMiddleWare())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		Skipper:    jwtSkipper(viper.GetString("skipjwt")),
		SigningKey: []byte(viper.GetString("secret")),
	}))
	e.HTTPErrorHandler = web.ErrorHandler

	return e
}

func Start(port int) {
	Echo = NewEcho(port)
	web.InitRoutes(Echo)

	Echo.Logger.Fatal(gracehttp.Serve(Echo.Server))
}
