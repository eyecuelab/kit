package web

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/eyecuelab/jsonapi"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestHealthz(t *testing.T) {
	t.Run("panic!", func(t *testing.T) {
		defer func() {
			recover()
		}()
		ctx := newMock(keyVal("oops", "1"), nil)
		Healthz(ctx)
		t.Error("should have panicked")
	})

	var want error = &jsonapi.ErrorObject{
		Title:  "Api Error",
		Status: fmt.Sprintf("%d", http.StatusBadRequest),
		Detail: "missing parameters: foo,bar",
	}
	got := Healthz(newMock(keyVal("apiErr", "1"), nil))
	assert.Equal(t, want, got)

	want = echo.NewHTTPError(http.StatusInternalServerError)
	assert.Equal(t, want, Healthz(newMock(keyVal("500", "1"), nil)))

	want = echo.NewHTTPError(http.StatusBadRequest, "Missing value.")
	assert.Equal(t, want, Healthz(newMock(keyVal("400", "1"), nil)))

	var ctx ApiContext = stringContext{newMock(nil, nil)}
	want = fmt.Errorf("%d:%s", http.StatusOK, "live")
	assert.Equal(t, want, Healthz(ctx))

}

type stringContext struct {
	ApiContext
}

func (ctx stringContext) String(code int, s string) error {
	return fmt.Errorf("%d:%s", code, s)
}
