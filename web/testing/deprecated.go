package testing

import (
	"net/http/httptest"
	"strings"

	"github.com/labstack/echo"
)

type EchoTest struct {
	Echo *echo.Echo
}

func (e *EchoTest) TestPost(url string, json *string) (echo.Context, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(echo.POST, url, strings.NewReader(*json))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.Echo.NewContext(req, rec)
	return ctx, rec
}

// func TestPost(url string, json *string) (echo.Context, *httptest.ResponseRecorder) {
// 	req := httptest.NewRequest(echo.POST, url, strings.NewReader(*json))
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	rec := httptest.NewRecorder()
// 	ctx := e.NewContext(req, rec)
// 	return ctx, rec
// }

// func TestGet(url string) (echo.Context, *httptest.ResponseRecorder) {
// 	req := httptest.NewRequest(echo.GET, url, nil)
// 	rec := httptest.NewRecorder()
// 	ctx := e.NewContext(req, rec)
// 	return ctx, rec
// }
