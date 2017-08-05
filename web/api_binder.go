package web

import (
	"regexp"
	"strings"

	"github.com/eyecuelab/kit/web/handlers"
	"github.com/google/jsonapi"
	"github.com/labstack/echo"
)

const MIMEJsonAPI = "application/vnd.api+json"

type ApiBinder struct{}

var notJsonApi = regexp.MustCompile("not a jsonapi")

func (cb *ApiBinder) Bind(i interface{}, c echo.Context) error {
	ctype := c.Request().Header.Get(echo.HeaderContentType)

	// s, err := ioutil.ReadAll(c.Request().Body)
	// if err != nil {
	// 	return err
	// }
	// log.Infof("body is: %s", s)

	if strings.HasPrefix(ctype, MIMEJsonAPI) {
		return jsonApiBind(c, i)
	}

	db := new(echo.DefaultBinder)
	if err := db.Bind(i, c); err != nil {
		return err
	}

	return nil
}

func jsonApiBind(c echo.Context, i interface{}) error {
	if err := jsonapi.UnmarshalPayload(c.Request().Body, i); err != nil {
		if match := notJsonApi.MatchString(err.Error()); match {
			return handlers.ApiError("Request Body is not valid JsonAPI")
		}
		return err
	}
	return nil
}
