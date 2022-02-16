package web

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"errors"

	"github.com/eyecuelab/kit/flect"
	"github.com/eyecuelab/kit/maputil"
	"github.com/eyecuelab/kit/web/meta"
	"github.com/google/jsonapi"
	"github.com/labstack/echo"
)

var reNotJsonApi = regexp.MustCompile("not a jsonapi|EOF")

func notJsonApi(err error) bool {
	return reNotJsonApi.MatchString(err.Error())
}

type (
	ApiContext interface {
		echo.Context

		Payload() *jsonapi.OnePayload
		Attrs(permitted ...string) map[string]interface{}
		AttrKeys() []string
		RequireAttrs(...string) error
		BindAndValidate(interface{}) error
		BindMulti(interface{}) ([]interface{}, error)
		BindIdParam(*int, ...string) error
		JsonApi(interface{}, int) error
		JsonApiOK(interface{}, ...interface{}) error
		JsonApiOKPaged(interface{}, *meta.Pagination, ...interface{}) error
		ApiError(string, ...int) *echo.HTTPError
		JsonAPIError(string, int, string) *jsonapi.ErrorObject
		QueryParamTrue(string) (bool, bool)

		RequiredQueryParams(...string) (map[string]string, error)
		OptionalQueryParams(...string) map[string]string
		QParams(...string) (map[string]string, error)
	}

	apiContext struct {
		echo.Context

		payload     *jsonapi.OnePayload
		manyPayload *jsonapi.ManyPayload
	}

	CommonExtendable interface {
		CommonExtend(interface{}) error
	}

	Extendable interface {
		Extend(interface{}) error
	}

	CommonMetable interface {
		CommonMeta() error
	}

	Metable interface {
		Meta() error
	}

	CommonLinkable interface {
		CommonLinks(*meta.Pagination) error
	}

	Linkable interface {
		Links(*meta.Pagination) error
	}
)

func (c *apiContext) Payload() *jsonapi.OnePayload {
	return c.payload
}

func (c *apiContext) Attrs(permitted ...string) map[string]interface{} {
	//TODO: remove this once all refactoring is complete
	if len(permitted) == 0 {
		return c.payload.Data.Attributes
	}

	permittedAttrs := make(map[string]interface{})
	for _, p := range permitted {
		if val, ok := c.payload.Data.Attributes[p]; ok {
			permittedAttrs[p] = val
		}
	}
	return permittedAttrs
}

func (c *apiContext) AttrKeys() []string {
	return maputil.Keys(c.Attrs())
}

func (c *apiContext) RequireAttrs(required ...string) error {
	missing := make([]string, 0, len(required))

	for _, key := range required {
		if c.payload.Data.Attributes[key] == nil {
			missing = append(missing, key)
			continue
		}
	}

	if len(missing) > 0 {
		return fmt.Errorf("missing required attributes: %v", missing)
	}

	return nil
}

//Before binding we make a copy of the req body and restore it after binding.
//This allows the body to be used again later
func (c *apiContext) Bind(i interface{}) error {
	body, err := c.readRestoreBody()
	if err != nil {
		return err
	}

	ctype := c.Request().Header.Get(echo.HeaderContentType)
	if isJSONAPI(ctype) {
		err = jsonAPIBind(c, i)
	} else {
		err = c.defaultBind(i)
	}

	c.restoreBody(body)

	return err
}

func (c *apiContext) BindMulti(containedType interface{}) ([]interface{}, error) {
	body, err := c.readRestoreBody()
	if err != nil {
		return nil, err
	}

	ctype := c.Request().Header.Get(echo.HeaderContentType)

	if !isJSONAPI(ctype) {
		return nil, errors.New("BindMulti only supports JSONApi, use Bind")
	}

	i, err := jsonAPIBindMulti(c, containedType)

	c.restoreBody(body)

	return i, err
}

func (c *apiContext) readRestoreBody() ([]byte, error) {
	b, err := ioutil.ReadAll(c.Request().Body)
	c.restoreBody(b)
	return b, err
}

func (c *apiContext) restoreBody(b []byte) {
	c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(b))
}

func (c *apiContext) defaultBind(i interface{}) error {
	db := new(echo.DefaultBinder)
	return db.Bind(i, c)
}

func isJSONAPI(s string) bool {
	const MIMEJsonAPI = "application/vnd.api+json"
	return strings.HasPrefix(s, MIMEJsonAPI)
}

func (c *apiContext) BindAndValidate(i interface{}) error {
	if err := c.Bind(i); err != nil {
		return err
	}
	if err := c.Validate(i); err != nil {
		return err
	}
	return nil
}

func (c *apiContext) JsonApi(i interface{}, status int) error {
	var buf bytes.Buffer
	if err := jsonapi.MarshalPayload(&buf, i); err != nil {
		return err
	}

	// These methods have to be the last thing called, *after* any error checks.
	// Once any of the Write methods are called, the response is "committed" and
	// cannot be changed. This causes error responses with 200 statuses.
	c.Response().Header().Set(echo.HeaderContentType, jsonapi.MediaType)
	c.Response().WriteHeader(status)
	c.Response().Write(buf.Bytes())
	return nil
}

func applyCommon(i interface{}, page *meta.Pagination, extendData ...interface{}) error {
	if casted, ok := i.(CommonExtendable); ok {
		if err := casted.CommonExtend(extendData); err != nil {
			return err
		}
	}

	if casted, ok := i.(CommonMetable); ok {
		if err := casted.CommonMeta(); err != nil {
			return err
		}
	}

	if casted, ok := i.(CommonLinkable); ok {
		if err := casted.CommonLinks(page); err != nil {
			return err
		}
	}
	return nil
}

func apply(i interface{}, page *meta.Pagination, extendData interface{}) error {
	if casted, ok := i.(Extendable); ok {
		if err := casted.Extend(extendData); err != nil {
			return err
		}
	}

	if casted, ok := i.(Metable); ok {
		if err := casted.Meta(); err != nil {
			return err
		}
	}

	if casted, ok := i.(Linkable); ok {
		if err := casted.Links(page); err != nil {
			return err
		}
	}
	return nil
}

func extendAndExtract(i interface{}, page *meta.Pagination, extendData interface{}) (data interface{}, err error) {
	if flect.IsSlice(i) {
		slice := reflect.ValueOf(i)
		for idx := 0; idx < slice.Len(); idx++ {
			elementInterface := slice.Index(idx).Interface()
			if err := applyCommon(elementInterface, page, extendData); err != nil {
				return nil, err
			}
		}
		return i, nil
	}

	if err := applyCommon(i, page, extendData); err != nil {
		return nil, err
	}

	if err := apply(i, page, extendData); err != nil {
		return nil, err
	}
	return i, nil
}

func (c *apiContext) JsonApiOK(i interface{}, extendData ...interface{}) error {
	var ed interface{}
	if len(extendData) > 0 {
		ed = extendData[0]
	}
	data, err := extendAndExtract(i, nil, ed)
	if err != nil {
		return err
	}
	return c.JsonApi(data, http.StatusOK)
}

func (c *apiContext) JsonApiOKPaged(i interface{}, page *meta.Pagination, extendData ...interface{}) error {
	var ed interface{}
	if len(extendData) > 0 {
		ed = extendData[0]
	}
	data, err := extendAndExtract(i, page, ed)
	if err != nil {
		return err
	}
	return c.JsonApi(data, http.StatusOK)
}

func (c *apiContext) BindIdParam(idValue *int, named ...string) (err error) {
	paramName := "id"
	if len(named) > 0 {
		paramName = named[0]
	}
	*idValue, err = strconv.Atoi(c.Param(paramName))
	return err
}

func (c *apiContext) QueryParamTrue(name string) (val, ok bool) {
	switch strings.ToLower(c.QueryParam(name)) {
	case "true", "1":
		return true, true
	case "false", "0":
		return false, true
	default:
		return false, false
	}
}

func jsonAPIBindMulti(c *apiContext, elementType interface{}) ([]interface{}, error) {
	buf := new(bytes.Buffer)
	tee := io.TeeReader(c.Request().Body, buf)

	unmarshaled, err := jsonapi.UnmarshalManyPayload(tee, reflect.TypeOf(elementType))
	if err != nil {
		return nil, err
	}

	c.manyPayload = new(jsonapi.ManyPayload)
	return unmarshaled, json.Unmarshal(buf.Bytes(), c.manyPayload)
}

func jsonAPIBind(c *apiContext, i interface{}) error {
	buf := new(bytes.Buffer)
	tee := io.TeeReader(c.Request().Body, buf)

	rType := reflect.TypeOf(i)

	if rType.Kind() == reflect.Slice {
		value := reflect.TypeOf(rType.Elem())

		unmarshaled, err := jsonapi.UnmarshalManyPayload(tee, value)
		if err != nil {
			return err
		}
		i = unmarshaled
	} else {
		if err := jsonapi.UnmarshalPayload(tee, i); err != nil {
			if notJsonApi(err) {
				return c.ApiError("Request Body is not valid JsonAPI")
			}
			return err
		}
	}

	c.payload = new(jsonapi.OnePayload)
	return json.Unmarshal(buf.Bytes(), c.payload)
}

func (c *apiContext) ApiError(msg string, codes ...int) *echo.HTTPError {
	if len(codes) > 0 {
		return echo.NewHTTPError(codes[0], msg)
	}
	// TODO: return jsonapi error instead
	return echo.NewHTTPError(http.StatusBadRequest, msg)
}

func (c *apiContext) JsonAPIError(msg string, code int, param string) *jsonapi.ErrorObject {
	return &jsonapi.ErrorObject{
		Status: fmt.Sprintf("%d", code),
		Title:  http.StatusText(code),
		Detail: msg,
		Meta: &map[string]interface{}{
			"parameter": param,
		},
	}
}

func (c *apiContext) RequiredQueryParams(required ...string) (map[string]string, error) {
	missing := make([]string, 0, len(required))
	params := make(map[string]string)

	for _, key := range required {
		val := c.QueryParam(key)
		if val == "" {
			missing = append(missing, key)
			continue
		}
		params[key] = val
	}

	if len(missing) > 0 {
		return nil, fmt.Errorf("missing required params: %v", missing)
	}

	return params, nil
}

func (c *apiContext) QParams(required ...string) (map[string]string, error) {
	return QParams(c, required...)
}

func QParams(c echo.Context, required ...string) (map[string]string, error) {
	missing := make([]string, 0, len(required))
	params := make(map[string]string)

	for k := range c.QueryParams() {
		params[k] = c.QueryParam(k)
	}

	for _, k := range required {
		if _, ok := params[k]; !ok {
			missing = append(missing, k)
		}
	}

	if len(missing) > 0 {
		return nil, fmt.Errorf("missing required params: %v", missing)
	}

	return params, nil
}

func (c *apiContext) OptionalQueryParams(optional ...string) map[string]string {
	params := make(map[string]string)
	for _, key := range optional {
		val := c.QueryParam(key)
		params[key] = val
	}
	return params
}

func ApiContextMiddleWare() func(echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return next(&apiContext{c, nil, nil})
		}
	}
}

func restrictedValue(value string, allowed []string, errorText string) (string, error) {
	if contains(allowed, value) {
		return value, nil
	}
	return "", fmt.Errorf(errorText, value)
}

func contains(set []string, s string) bool {
	for _, v := range set {
		if s == v {
			return true
		}
	}
	return false
}
