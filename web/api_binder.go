package web

import (
	"regexp"

	"github.com/eyecuelab/kit/web/handlers"
	"github.com/google/jsonapi"
	"github.com/labstack/echo"
)

type ApiBinder struct{}

var notJsonApi = regexp.MustCompile("not a jsonapi")

func (cb *ApiBinder) Bind(i interface{}, c echo.Context) error {
	// ctype := req.Header.Get(echo.HeaderContentType)

	if err := jsonapi.UnmarshalPayload(c.Request().Body, i); err != nil {
		if match := notJsonApi.MatchString(err.Error()); match {
			return handlers.ApiError("Request Body is not valid JsonAPI")
		}
		return err
	}
	return nil
}
