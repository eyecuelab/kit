package web

import (
	"bytes"
	"fmt"
	"net/http"
	"strconv"

	"github.com/eyecuelab/kit/goenv"
	"github.com/eyecuelab/kit/log"
	"github.com/google/jsonapi"
	"github.com/labstack/echo"
)

func ErrorHandler(err error, c echo.Context) {
	if c.Response().Committed {
		return
	}

	code, apiError := toApiError(err)

	if c.Request().Method == "HEAD" { // Issue #608
		if err := c.NoContent(code); err != nil {
			goto ERROR
		}
		return
	}

	if err := renderApiErrors(c, apiError); err != nil {
		logErr(err)
		goto ERROR
	}

	if code < 500 {
		return
	}
ERROR:
	logErr(err)
}

func logErr(err error) {
	log.ErrorWrap(err, "Uncaught Error")
}

func toApiError(err error) (int, *jsonapi.ErrorObject) {
	code := http.StatusInternalServerError

	if he, ok := err.(*jsonapi.ErrorObject); ok {
		code, _ = strconv.Atoi(he.Status)
		return code, he
	}

	var detail interface{}

	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		detail = he.Message
	} else if !goenv.Prod {
		detail = err.Error()
	}

	return code, valuesToApiError(code, http.StatusText(code), &detail)
}

func renderApiErrors(c echo.Context, errors ...*jsonapi.ErrorObject) error {
	var b bytes.Buffer
	if err := jsonapi.MarshalErrors(&b, errors); err != nil {
		return err
	}

	if i, err := strconv.Atoi(errors[0].Status); err != nil {
		return err
	} else {
		return c.Blob(i, echo.MIMEApplicationJSON, b.Bytes())
	}
}

func valuesToApiError(status int, title string, detail *interface{}) *jsonapi.ErrorObject {
	var d string
	if *detail != nil {
		d = fmt.Sprint(*detail)
	}
	return &jsonapi.ErrorObject{
		Status: fmt.Sprintf("%d", status),
		Title:  title,
		Detail: d,
	}
}
