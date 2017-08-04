package web

import (
	"fmt"

	valid "github.com/asaskevich/govalidator"
	"github.com/facebookgo/grace/gracehttp"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type apiValidator struct{}

func (v *apiValidator) Validate(i interface{}) error {
	_, err := valid.ValidateStruct(i)
	return err
}

var Echo *echo.Echo

func NewEcho(port int) *echo.Echo {
	e := echo.New()
	e.Validator = &apiValidator{}
	e.Server.Addr = fmt.Sprintf(":%v", port)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.HTTPErrorHandler = ErrorHandler
	e.Binder = &ApiBinder{}

	return e
}

func Start(port int) {
	Echo = NewEcho(port)
	initRoutes()

	Echo.Logger.Fatal(gracehttp.Serve(Echo.Server))
}

func initRoutes() {
	for _, route := range Routing.Routes {
		for _, handler := range route.Handlers {
			Echo.Add(handler.Method, route.Path, handler.Handler, handler.MiddleWare...)
		}
	}
}
