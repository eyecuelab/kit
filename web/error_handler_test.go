package web

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/asaskevich/govalidator"

	"github.com/google/jsonapi"
	"github.com/lib/pq"

	"github.com/eyecuelab/kit/errorlib"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

const (
	someErr errorlib.ErrorString = "some error"
)

type noContentContext struct {
	echo.Context
}

func (ctx noContentContext) NoContent(code int) error {
	return someErr
}

var _ echo.Context = noContentContext{}

func newContext(method string) echo.Context {
	req := httptest.NewRequest(method, "http://localhost.com/whatever", new(bytes.Buffer))
	rec := httptest.NewRecorder()
	return echo.New().NewContext(req, rec)
}

func TestErrorHandler(t *testing.T) {
	ctx := newContext("GET")
	ctx.Response().Committed = true
	assert.Equal(t, alreadyCommited, ErrorHandler(someErr, ctx))

	noContentCtx := noContentContext{newContext("HEAD")}
	assert.Equal(t, noContent, ErrorHandler(someErr, noContentCtx))

	ctx = newContext("HEAD")
	var err error = &echo.HTTPError{Code: 1222, Message: "magic error"}
	assert.Equal(t, methodIsHead, ErrorHandler(err, ctx))

	ctx = newContext("GET")
	assert.Equal(t, nilErr, ErrorHandler(nil, ctx))

	ctx = newContext("GET")
	httpErr := echo.HTTPError{Code: 500, Message: "magic error", Inner: someErr}
	assert.Equal(t, statusOver500, ErrorHandler(&httpErr, ctx))

	ctx = newContext("GET")
	httpErr = echo.HTTPError{Code: 200, Message: "OK", Inner: someErr}
	assert.Equal(t, ignoredErr, ErrorHandler(&httpErr, ctx))

	ctx = newContext("GET")
	err = &jsonapi.ErrorObject{Title: "hey", Status: "deliberately not an integer"}
	assert.Equal(t, problemRendering, ErrorHandler(err, ctx))

}

func Test_logErr(t *testing.T) {
}

func Test_toApiError(t *testing.T) {
	status, apiErr := toApiError(nil)
	assert.Nil(t, apiErr)
	assert.Equal(t, 200, status)

	errObj := jsonapi.ErrorObject{Status: "200"}
	status, apiErr = toApiError(&errObj)
	assert.Equal(t, 200, status)
	assert.Equal(t, errObj, *apiErr)

	//bad status
	errObj = jsonapi.ErrorObject{Status: "foobar"}
	wantApiErr := jsonapi.ErrorObject{Status: "foobar", Detail: " bad status: foobar"}
	gotStatus, gotApiErr := toApiError(&errObj)
	assert.Equal(t, http.StatusInternalServerError, gotStatus)
	assert.Equal(t, wantApiErr, *apiErr)

	//httperr
	wantApiErr = jsonapi.ErrorObject{
		Status: "404",
		Title:  http.StatusText(404),
		Detail: "not found",
	}

	wantStatus := 404
	httpErr := echo.HTTPError{Code: 404, Message: "not found", Inner: someErr}
	gotStatus, gotApiErr = toApiError(&httpErr)
	assert.Equal(t, wantStatus, gotStatus)
	assert.Equal(t, wantApiErr, *gotApiErr)

	//pqErr
	const (
		msg      = "ok"
		code     = "0100C"
		codeName = "dynamic_result_sets_returned"
	)
	pqErr := pq.Error{Message: msg, Code: "0100C"}
	wantStatus = http.StatusBadRequest
	wantApiErr = jsonapi.ErrorObject{
		Status: toStr(http.StatusBadRequest),
		Detail: msg,
		Code:   codeName,
		Title:  "Bad Request"}
	gotStatus, gotApiErr = toApiError(&pqErr)
	assert.Equal(t, wantStatus, gotStatus)
	assert.Equal(t, wantApiErr, *gotApiErr)

	//goValidator
	gvErr := govalidator.Errors{someErr}
	wantApiErr = jsonapi.ErrorObject{
		Title:  http.StatusText(http.StatusBadRequest),
		Detail: gvErr.Error(),
		Code:   "",
		Status: toStr(http.StatusBadRequest),
	}
	wantStatus = http.StatusBadRequest

	gotStatus, gotApiErr = toApiError(gvErr)
	assert.Equal(t, wantStatus, gotStatus)
	assert.Equal(t, wantApiErr, *gotApiErr)
}

func toStr(n int) string {
	return fmt.Sprintf("%d", n)
}
func Test_renderApiErrors(t *testing.T) {
	ctx := newContext("HEAD")
	assert.Error(t, renderApiErrors(ctx, nil))

	assert.NoError(t, renderApiErrors(ctx, &jsonapi.ErrorObject{ID: "ok", Status: "404", Code: "hey"}))
}

func Test_valuesToApiError(t *testing.T) {

}
