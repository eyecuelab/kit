package web

import (
	"bytes"
	"errors"
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

//these are for internal testing.
type errorHandlerCode byte

const (
	alreadyCommited errorHandlerCode = iota
	methodIsHead
	noContent
	rendered
	statusOver500
	handledError
	nilErr
)

var pq500s = map[string]bool{
	"undefined_function": true,
}

func ErrorHandler(err error, c echo.Context) errorHandlerCode {
	if c.Response().Committed {
		return alreadyCommited
	}
	if err == nil {
		logErr(errors.New("nil error sent into ErrorHandler"), c)
		return nilErr
	}

	status, apiError := toApiError(err)

	if c.Request().Method == "HEAD" {
		if err := c.NoContent(status); err != nil {
			logErr(err, c)
			return noContent
		}
		return methodIsHead
	}
	if renderedErr := renderApiErrors(c, apiError); renderedErr != nil {
		logErr(renderedErr, c) //are we supposed to log this error twice?
		logErr(err, c)
		return rendered
	}
	if status >= 500 {
		logErr(err, c)
		return statusOver500
	}
	return handledError
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
	case nil:
		panic("nil error sent into toApiError")

	case *jsonapi.ErrorObject:
		if status, convErr := strconv.Atoi(err.Status); convErr == nil {
			return status, err
		}
		err.Detail += fmt.Sprintf(" bad status: %s", err.Status)

		return status, err

	case *echo.HTTPError:
		status = err.Code
		if err.Message != nil {
			detail = fmt.Sprint(err.Message)
		}
		return status, valuesToApiError(status, http.StatusText(status), detail, code)

	case *pq.Error:
		detail = err.Message
		code = err.Code.Name()
		if _, ok := pq500s[code]; !ok {
			status = http.StatusBadRequest
		}
		return status, valuesToApiError(status, http.StatusText(status), detail, code)

	case govalidator.Errors:
		status, detail = http.StatusBadRequest, err.Error()
		return status, valuesToApiError(status, http.StatusText(status), detail, code)

	default:
		detail = err.Error()
		return status, valuesToApiError(status, http.StatusText(status), detail, code)
	}
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
