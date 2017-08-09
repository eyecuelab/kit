package web

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"regexp"
	"strings"

	"github.com/eyecuelab/kit/maputil"
	"github.com/google/jsonapi"
	"github.com/labstack/echo"
)

const MIMEJsonAPI = "application/vnd.api+json"

var notJsonApi = regexp.MustCompile("not a jsonapi")

type ApiContext struct {
	echo.Context

	Payload *jsonapi.OnePayload
}

func (c *ApiContext) Attrs() map[string]interface{} {
	return c.Payload.Data.Attributes
}

func (c *ApiContext) AttrKeys() []string {
	return maputil.Keys(c.Attrs())
}

func (c *ApiContext) Bind(i interface{}) error {
	ctype := c.Request().Header.Get(echo.HeaderContentType)

	if strings.HasPrefix(ctype, MIMEJsonAPI) {
		return jsonApiBind(c, i)
	}

	db := new(echo.DefaultBinder)
	if err := db.Bind(i, c); err != nil {
		return err
	}

	return nil
}

func (c *ApiContext) BindAndValidate(i interface{}) error {
	if err := c.Bind(i); err != nil {
		return err
	}
	if err := c.Validate(i); err != nil {
		return err
	}
	return nil
}

func (c *ApiContext) JsonApi(i interface{}, status int) error {
	c.Response().Header().Set(echo.HeaderContentType, jsonapi.MediaType)
	c.Response().WriteHeader(status)

	return jsonapi.MarshalPayload(c.Response().Writer, i)
}

func (c *ApiContext) JsonApiOK(i interface{}) error {
	return c.JsonApi(i, http.StatusOK)
}

func jsonApiBind(c *ApiContext, i interface{}) error {
	var buf bytes.Buffer
	tee := io.TeeReader(c.Request().Body, &buf)

	if err := jsonapi.UnmarshalPayload(tee, i); err != nil {
		if notJsonApi.MatchString(err.Error()) {
			return c.ApiError("Request Body is not valid JsonAPI")
		}
		return err
	}

	c.Payload = new(jsonapi.OnePayload)
	if err := json.Unmarshal(buf.Bytes(), c.Payload); err != nil {
		return err
	}

	return nil
}

func (c *ApiContext) ApiError(msg string, codes ...int) *echo.HTTPError {
	status := http.StatusBadRequest
	if len(codes) > 0 {
		status = codes[0]
	}

	return echo.NewHTTPError(status, msg)
}

func ApiContextMiddleWare() func(echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ac := &ApiContext{c, nil}
			return next(ac)
		}
	}
}
