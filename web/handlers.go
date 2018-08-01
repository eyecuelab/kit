package web

import (
	"fmt"
	"net/http"

	"github.com/eyecuelab/jsonapi"
	"github.com/labstack/echo"
	"github.com/spf13/viper"
)

func Healthz(c ApiContext) error {
	badParam := func(param string) bool { return c.QueryParam(param) == "1" }
	switch {
	case badParam("oops"):
		panic("oops")
	case badParam("apiErr"):
		return &jsonapi.ErrorObject{
			Title:  "Api Error",
			Status: fmt.Sprintf("%d", http.StatusBadRequest),
			Detail: "missing parameters: foo,bar",
		}
	case badParam("500"):
		return echo.NewHTTPError(http.StatusInternalServerError)
	case badParam("400"):
		return echo.NewHTTPError(http.StatusBadRequest, "Missing value.")
	default:
		return c.String(http.StatusOK, "live")
	}
}

func Config(c ApiContext) error {
	return c.JSON(http.StatusOK, viper.AllSettings())
}
