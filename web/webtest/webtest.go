package webtest

import (
	"net/http/httptest"
	"strings"

	"github.com/labstack/echo"
)

func Post(url string, json string) (ctx echo.Context, rec *httptest.ResponseRecorder) {
	req := httptest.NewRequest(echo.POST, url, strings.NewReader(json))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec = httptest.NewRecorder()
	ctx = echo.New().NewContext(req, rec)
	return ctx, rec
}
