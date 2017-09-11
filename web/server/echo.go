package server

import (
	"fmt"
	"regexp"

	valid "github.com/asaskevich/govalidator"
	"github.com/eyecuelab/kit/web"
	"github.com/facebookgo/grace/gracehttp"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	emw "github.com/labstack/echo/middleware"
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
		Skipper:    AuthedSkipper(),
		SigningKey: []byte(viper.GetString("secret")),
	}))
	e.Use(mws...)

	return e
}

func Start(port int) {
	Echo = NewEcho(port)
	web.InitRoutes(Echo)

	Echo.Logger.Fatal(gracehttp.Serve(Echo.Server))
}

func AddMiddleWare(mw echo.MiddlewareFunc) {
	mws = append(mws, mw)
}

func AuthedSkipper() func(echo.Context) bool {
	skip := viper.GetString("skipjwt")
	if len(skip) == 0 {
		return emw.DefaultSkipper
	}

	re := regexp.MustCompile(skip)
	return func(c echo.Context) bool {
		return isOptionsRequest(c) || re.MatchString(c.Request().URL.Path)
	}
}

func isOptionsRequest(c echo.Context) bool {
	return c.Request().Method == echo.OPTIONS
}
