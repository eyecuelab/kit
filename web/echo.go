package web

import (
	"fmt"

	valid "github.com/asaskevich/govalidator"
	"github.com/facebookgo/grace/gracehttp"
	"github.com/labstack/echo"
)

type apiValidator struct{}

func (v *apiValidator) Validate(i interface{}) error {
	_, err := valid.ValidateStruct(i)
	return err
}

var Server *echo.Echo

func NewEcho(port int) *echo.Echo {
	e := echo.New()
	e.Validator = &apiValidator{}
	e.Server.Addr = fmt.Sprintf(":%v", port)

	return e
}

func Start(port int) {
	Server = NewEcho(port)
	initRoutes()

	Server.Logger.Fatal(gracehttp.Serve(Server.Server))
}

func initRoutes() {
	for _, route := range Routing.Routes {
		for _, handler := range route.Handlers {
			Server.Add(handler.Method, route.Path, handler.Handler, handler.MiddleWare...)
		}
	}
}
