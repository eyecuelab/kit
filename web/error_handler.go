package web

import (
	"bytes"
	"fmt"
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/eyecuelab/kit/goenv"
	"github.com/eyecuelab/kit/log"
	"github.com/google/jsonapi"
	"github.com/labstack/echo"
	"github.com/lib/pq"
)

var pq500s = map[string]bool{
	"undefined_function": true,
}

func ErrorHandler(err error, c echo.Context) {
	if c.Response().Committed {
		return
	}

	status, apiError := toApiError(err)

	if c.Request().Method == "HEAD" {
		if err := c.NoContent(status); err != nil {
			goto ERROR
		}
		return
	}

	if err := renderApiErrors(c, apiError); err != nil {
		logErr(err)
		goto ERROR
	}

	if status < 500 {
		return
	}
ERROR:
	logErr(err)
}

func logErr(err error) {
	log.ErrorWrap(err, "Uncaught Error")
}

func toApiError(err error) (int, *jsonapi.ErrorObject) {
	status := http.StatusInternalServerError

	if he, ok := err.(*jsonapi.ErrorObject); ok {
		status, _ = strconv.Atoi(he.Status)
		return status, he
	}

	var (
		detail interface{}
		code   string
	)
	if he, ok := err.(*echo.HTTPError); ok {
		status = he.Code
		detail = he.Message
	} else if he, ok := err.(*pq.Error); ok {
		detail = he.Message
		code = he.Code.Name()
		if !pq500s[code] {
			status = http.StatusBadRequest
		}
	} else if he, ok := err.(govalidator.Errors); ok {
		status = http.StatusBadRequest
		detail = he.Error()
	} else if !goenv.Prod {
		detail = err.Error()
	}

	return status, valuesToApiError(status, http.StatusText(status), &detail, code)
}

func renderApiErrors(c echo.Context, errors ...*jsonapi.ErrorObject) error {
	var b bytes.Buffer
	if err := jsonapi.MarshalErrors(&b, errors); err != nil {
		return err
	}

	if i, err := strconv.Atoi(errors[0].Status); err != nil {
		return err
	} else {
		return c.Blob(i, jsonapi.MediaType, b.Bytes())
	}
}

func valuesToApiError(status int, title string, detail *interface{}, code string) *jsonapi.ErrorObject {
	var d string
	if *detail != nil {
		d = fmt.Sprint(*detail)
	}
	return &jsonapi.ErrorObject{
		Status: fmt.Sprintf("%d", status),
		Title:  title,
		Detail: d,
		Code:   code,
	}
}
