package handlers

import (
	"net/http"

	"github.com/google/jsonapi"
	"github.com/labstack/echo"
)

func ApiError(msg string, codes ...int) *echo.HTTPError {
	status := http.StatusBadRequest
	if len(codes) > 0 {
		status = codes[0]
	}

	return echo.NewHTTPError(status, msg)
}

func JsonApi(c echo.Context, status int, i interface{}) error {
	c.Response().Header().Set(echo.HeaderContentType, jsonapi.MediaType)
	c.Response().WriteHeader(status)

	return jsonapi.MarshalPayload(c.Response().Writer, i)
}

func BindAndValidate(c echo.Context, i interface{}) error {
	if err := c.Bind(i); err != nil {
		return err
	}
	if err := c.Validate(i); err != nil {
		return err
	}

	return nil
}
