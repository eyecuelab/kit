package web

import (
	"bytes"
	"fmt"
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/eyecuelab/kit/log"
	"github.com/google/jsonapi"
	"github.com/labstack/echo"
	"github.com/lib/pq"
	"github.com/eyecuelab/kit/brake"
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
		//FIXME: are we supposed to log this error twice?
		logErr(err, c)
		goto ERROR
	}

	if status < 500 {
		return
	}
ERROR: //FIXME - I feel like we should justify a GOTO
	logErr(err, c)
}

func logErr(err error, c echo.Context) {
	brake.Notify(err, c.Request())
	log.ErrorWrap(err, "Uncaught Error")
}

func toApiError(err error) (status int, apiErr *jsonapi.ErrorObject) {

	var (
		detail interface{}
		code   string
	)
	status = http.StatusInternalServerError
	switch err := err.(type) {
	case *jsonapi.ErrorObject:
		status, _ = strconv.Atoi(err.Status)
		return status, err

	case *echo.HTTPError:
		status, detail = err.Code, err.Message
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
