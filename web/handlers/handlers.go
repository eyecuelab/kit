package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

func ApiError(msg string, codes ...int) *echo.HTTPError {
	status := http.StatusBadRequest
	if len(codes) > 0 {
		status = codes[0]
	}

	return echo.NewHTTPError(status, msg)
}
