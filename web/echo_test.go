package web

import (
	"net/http/httptest"
	"strings"

	"github.com/labstack/echo"
)
//
func TestPost(e *echo.Echo, url, string, json *string) (echo.Context, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(echo.POST, url, strings.NewReader(*json))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	return ctx, rec
}
