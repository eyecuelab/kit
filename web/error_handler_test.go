package web

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

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
	//httpErr := echo.HTTPError{Code: 300, Message: "magic error", Inner: someErr}
	//assert.Equal(t, rendered, ErrorHandler(&httpErr, ctx))
}

func Test_logErr(t *testing.T) {
}

func Test_toApiError(t *testing.T) {
	t.Run("panic!", func(t *testing.T) {
		defer func() {
			recover()
		}()
		toApiError(nil)
		t.Error("should have panicked")
	})

	errObj := jsonapi.ErrorObject{Status: "200"}
	status, apiErr := toApiError(&errObj)
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

}

func toStr(n int) string {
	return fmt.Sprintf("%d", n)
}
func Test_renderApiErrors(t *testing.T) {

}

func Test_valuesToApiError(t *testing.T) {

}
