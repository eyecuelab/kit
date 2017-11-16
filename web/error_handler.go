package web

import (
	"bytes"
	"fmt"
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/eyecuelab/kit/brake"
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
			logErr(err, c)
		}
	} else if renderedErr := renderApiErrors(c, apiError); err != nil {
		logErr(renderedErr, c) //are we supposed to log this error twice?
		logErr(err, c)
	} else if status >= 500 {
		logErr(err, c)
	}
}

func logErr(err error, c echo.Context) {
	brake.Notify(err, c.Request())
	log.ErrorWrap(err, "Uncaught Error")
}

func toApiError(err error) (status int, apiErr *jsonapi.ErrorObject) {

	var (
		detail string
		code   string
	)
	status = http.StatusInternalServerError
	switch err := err.(type) {
	case *jsonapi.ErrorObject:
		status, _ = strconv.Atoi(err.Status)
		return status, err

	case *echo.HTTPError:
		status = err.Code
		if err.Message != nil {
			detail = fmt.Sprint(err.Message)
		}
	case *pq.Error:
		detail = err.Message
		code = err.Code.Name()
		if _, ok := pq500s[code]; !ok {
			status = http.StatusBadRequest
		}
	case govalidator.Errors:
		status, detail = http.StatusBadRequest, err.Error()
	default:
		detail = err.Error()

	}

	return status, valuesToApiError(status, http.StatusText(status), detail, code)
}

func renderApiErrors(c echo.Context, errors ...*jsonapi.ErrorObject) (err error) {
	var b bytes.Buffer
	if err = jsonapi.MarshalErrors(&b, errors); err != nil {
		return err
	}
	var code int
	if code, err = strconv.Atoi(errors[0].Status); err != nil {
		return err
	}
	return c.Blob(code, jsonapi.MediaType, b.Bytes())

}

func valuesToApiError(status int, title, detail, code string) *jsonapi.ErrorObject {
	return &jsonapi.ErrorObject{
		Status: fmt.Sprintf("%d", status),
		Title:  title,
		Detail: detail,
		Code:   code,
	}
}
