package web

import (
	"fmt"

	valid "github.com/asaskevich/govalidator"
	"github.com/labstack/echo"
)

type apiValidator struct{}

func (v *apiValidator) Validate(i interface{}) error {
	_, err := valid.ValidateStruct(i)
	return err
}

func NewEcho(port int) *echo.Echo {
	e := echo.New()
	e.Validator = &apiValidator{}
	e.Server.Addr = fmt.Sprintf(":%v", port)

	return e
}
