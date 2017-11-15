package web

import (
	"testing"

	"github.com/eyecuelab/kit/errorlib"

	"github.com/google/jsonapi"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

const (
	foo = "foo"
	bar = "bar"
	baz = "baz"
)

type mockParamsEchoContext struct {
	queryParams map[string]string
	params      map[string]string
	echo.Context
}

var _ echo.Context = mockParamsEchoContext{}

func (m mockParamsEchoContext) QueryParam(name string) string {
	v, _ := m.queryParams[name]
	return v
}

func (m mockParamsEchoContext) Param(name string) string {
	v, _ := m.params[name]
	return v
}

func newMock(queryParams, params map[string]string) ApiContext {
	if queryParams == nil {
		queryParams = make(map[string]string)
	}
	if params == nil {
		params = make(map[string]string)
	}
	return &apiContext{
		Context: mockParamsEchoContext{
			Context:     echo.New().AcquireContext(),
			params:      params,
			queryParams: queryParams,
		},
		payload: new(jsonapi.OnePayload),
	}
}

func Test_apiContext_RequiredQueryParams(t *testing.T) {
	want := map[string]string{
		foo: foo,
		bar: bar,
	}
	ctx := newMock(want, nil)
	got, err := ctx.RequiredQueryParams(foo, bar)
	assert.NoError(t, err)
	assert.Equal(t, want, got)

	_, err = ctx.RequiredQueryParams(foo, bar, baz)
	assert.Error(t, err)
}

func Test_apiContext_OptionalQueryParams(t *testing.T) {
	want := map[string]string{
		foo: foo,
		bar: bar,
		baz: "", //deliberately missing
	}
	ctx := newMock(map[string]string{foo: foo, bar: bar}, nil)
	assert.Equal(t, want, ctx.OptionalQueryParams(foo, bar, baz))
}

func Test_apiContext_BindIdParam(t *testing.T) {

	var got int
	want := 20

	ctx := newMock(nil, pair("id", "20"))
	assert.NoError(t, ctx.BindIdParam(&got))
	assert.Equal(t, want, got)

	got, want = 0, 40
	ctx = newMock(nil, pair("foobar", "40"))
	assert.NoError(t, ctx.BindIdParam(&got, "foobar"))

}

func pair(key, val string) map[string]string {
	return map[string]string{key: val}
}
func Test_QueryParamTrue(t *testing.T) {
	const foo = foo

	val, ok := newMock(pair(foo, "1"), nil).QueryParamTrue(foo)
	assert.True(t, val)
	assert.True(t, ok)

	val, ok = newMock(pair(foo, "true"), nil).QueryParamTrue(foo)
	assert.True(t, val)
	assert.True(t, ok)

	val, ok = newMock(pair(foo, "0"), nil).QueryParamTrue(foo)
	assert.False(t, val)
	assert.True(t, ok)

	val, ok = newMock(pair(foo, "false"), nil).QueryParamTrue(foo)
	assert.False(t, val)
	assert.True(t, ok)

	val, ok = newMock(pair(foo, "asdaqsd"), nil).QueryParamTrue(foo)
	assert.False(t, val)
	assert.False(t, ok)
}

func Test_RestrictedParam(t *testing.T) {
	allowed := []string{foo, bar, baz}

	val, err := newMock(nil, pair(foo, foo)).RestrictedParam(foo, allowed...)
	assert.Equal(t, val, foo)
	assert.NoError(t, err)

	val, err = newMock(nil, pair(foo, foo)).RestrictedParam(foo, "asdhjasod")
	assert.NotEqual(t, val, foo)
	assert.Error(t, err)
}

func Test_notJsonApi(t *testing.T) {
	assert.True(t, notJsonApi(errorlib.ErrorString("not a jsonapi")))
	assert.True(t, notJsonApi(errorlib.ErrorString("EOF")))
	assert.False(t, notJsonApi(errorlib.ErrorString("foobar")))
}
