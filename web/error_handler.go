package web

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/eyecuelab/kit/brake"
	"github.com/eyecuelab/kit/functools"
	"github.com/eyecuelab/kit/log"
	"github.com/google/jsonapi"
	"github.com/labstack/echo"
	"github.com/lib/pq"
)

//testCode is for internal testing
type testCode byte

const (
	alreadyCommited testCode = iota
	methodIsHead
	noContent
	problemRendering
	nilErr
	normal
)

var pq500s = map[string]bool{
	"undefined_function": true,
}

var criticalKeywords = []string{
	"reflect",
	"dereference",
	"runtime",
}

//errorHandler is the internal implementation of ErrorHandler. It returns a testCode,
//which does not fufill the interface expected by echo. Thus, it is wrapped by the function ErrorHandler.
func errorHandler(err error, c echo.Context) testCode {
	if c.Response().Committed {
		return alreadyCommited
	}
	if err == nil {
		trackErr(errors.New("nil error sent into ErrorHandler"), c, 0)
		return nilErr
	}

	status, apiError := toApiError(err)

	if c.Request().Method == "HEAD" {
		if err := c.NoContent(status); err != nil {
			trackErr(err, c, status)
			return noContent
		}
		return methodIsHead
	}

	if errRendering := renderApiErrors(c, apiError); errRendering != nil {
		trackErr(errRendering, c, 0)
		trackErr(err, c, status)
		return problemRendering
	}

	trackErr(err, c, status)
	return normal
}

//ErrorHandler handles errors
func ErrorHandler(err error, c echo.Context) {
	errorHandler(err, c)
}

func toApiError(err error) (status int, apiErr *jsonapi.ErrorObject) {
	var (
		detail string
		code   string
	)
	status = http.StatusInternalServerError
	switch err := err.(type) {
	case nil:
		return 200, nil
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
		return status, errorObj(status, http.StatusText(status), detail, code)

	case *pq.Error:
		detail = err.Message
		code = err.Code.Name()
		if _, ok := pq500s[code]; !ok {
			status = http.StatusBadRequest
		}
		return status, errorObj(status, http.StatusText(status), detail, code)

	case govalidator.Errors:
		status, detail = http.StatusBadRequest, err.Error()
		return status, errorObj(status, http.StatusText(status), detail, code)

	default:
		detail = err.Error()
		return status, errorObj(status, http.StatusText(status), detail, code)
	}
}

func renderApiErrors(c echo.Context, errors ...*jsonapi.ErrorObject) (err error) {
	var b bytes.Buffer
	if emptyOrAllNil(errors) {
		return fmt.Errorf("no errors to render")
	}

	if err = jsonapi.MarshalErrors(&b, errors); err != nil {
		return err
	}
	var code int
	if code, err = strconv.Atoi(errors[0].Status); err != nil {
		return err
	}
	return c.Blob(code, jsonapi.MediaType, b.Bytes())
}

func emptyOrAllNil(errs []*jsonapi.ErrorObject) bool {
	for _, err := range errs {
		if err != nil {
			return false
		}
	}
	return true
}

func errorObj(status int, title, detail, code string) *jsonapi.ErrorObject {
	return &jsonapi.ErrorObject{
		Status: fmt.Sprintf("%d", status),
		Title:  title,
		Detail: detail,
		Code:   code,
	}
}

func trackErr(err error, c echo.Context, status int) {
	notifyErr(err, c, status)
	if status == 0 || status >= 500 {
		logErr(err)
	}
}

func logErr(err error) {
	log.ErrorWrap(err, "Uncaught Error")
}

func notifyErr(err error, c echo.Context, status int) {
	if status == http.StatusUnauthorized {
		return
	}

	sev := brake.SeverityError
	if isCritical(err) {
		sev = brake.SeverityCritical
	} else if status > 0 && status < 500 {
		sev = brake.SeverityWarn
	}

	brake.Notify(err, c.Request(), sev)
}

func isCritical(err error) bool {
	return functools.StringContainsAny(err.Error(), criticalKeywords...)
}
