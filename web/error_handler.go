package web

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/eyecuelab/kit/goenv"
	"github.com/google/jsonapi"
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

	if !c.Response().Committed {
		if c.Request().Method == "HEAD" { // Issue #608
			if err := c.NoContent(code); err != nil {
				goto ERROR
			}
		} else if err, b := errorToBytes(&msg, code); err != nil {
			goto ERROR
		} else if err := c.Blob(code, echo.MIMEApplicationJSON, b.Bytes()); err != nil {
			goto ERROR
			// c.Blob(code, echo.MIMEApplicationJSON, b.Bytes())
		}
		// if err, b := errorToBytes(&msg, code); err != nil {
		// 	goto ERROR
		// } else {
		// 	c.Blob(code, echo.MIMEApplicationJSON, b.Bytes())
		// }
		// }
	}
ERROR:
	c.Logger().Error(err)
}

func errorToBytes(title *interface{}, code int) (error, *bytes.Buffer) {
	var b bytes.Buffer
	err := jsonapi.MarshalErrors(&b, []*jsonapi.ErrorObject{{
		Title:  fmt.Sprint(*title),
		Detail: "",
		Status: fmt.Sprintf("%d", code),
	}})
	return err, &b
}
