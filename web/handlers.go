package web

import (
	"fmt"
	"net/http"

	"github.com/google/jsonapi"
	"github.com/labstack/echo"
	"github.com/spf13/viper"
)

func Healthz(c *ApiContext) error {
	if c.QueryParam("oops") == "1" {
		panic("oops")
	}
	if c.QueryParam("apierr") == "1" {
		return &jsonapi.ErrorObject{
			Title:  "Api Error",
			Status: fmt.Sprintf("%d", http.StatusBadRequest),
			Detail: "missing parameters: foo,bar",
		}
	}
	if c.QueryParam("500") == "1" {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	if c.QueryParam("400") == "1" {
		return echo.NewHTTPError(http.StatusBadRequest, "Missing value.")
	}
	return c.String(http.StatusOK, "live")
}

func Config(c *ApiContext) error {
	return c.JSON(http.StatusOK, viper.AllSettings())
}
