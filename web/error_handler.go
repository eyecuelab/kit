package web

import (
	"net/http"

	"github.com/eyecuelab/kit/goenv"
	"github.com/labstack/echo"
)

func ErrorHandler(err error, c echo.Context) {
	var (
		code = http.StatusInternalServerError
		msg  interface{}
	)

	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		msg = he.Message
	} else if !goenv.Prod {
		msg = err.Error()
	} else {
		// will result in 500 "Internal Server Error" on panic
		msg = http.StatusText(code)
	}
	if _, ok := msg.(string); ok {
		msg = echo.Map{"message": msg}
	}

	if !c.Response().Committed {
		if c.Request().Method == "HEAD" { // Issue #608
			if err := c.NoContent(code); err != nil {
				goto ERROR
			}
		} else {
			if err := c.JSON(code, msg); err != nil {
				goto ERROR
			}
		}
	}
ERROR:
	c.Logger().Error(err)
}
