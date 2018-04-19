package integration

import (
	"fmt"
	"testing"

	"github.com/parnurzeal/gorequest"
	"github.com/stretchr/testify/assert"
)

// AssertGetOK assert GET request is OK
func AssertGetOK(t *testing.T, path string, token string) (resp gorequest.Response, body string, errs []error) {
	resp, body, errs = Request("GET", path, token).End()
	assert.Empty(t, errs)
	AssertStatusOK(t, resp)

	return resp, body, errs
}

// AssertAttr ...
func AssertAttr(t *testing.T, i interface{}, m map[string]interface{}) {
	if i == nil {
		assert.True(t, false, "Expected Attributes got nil")
		return
	}
	attrs := i.(map[string]interface{})
	for k, v := range m {
		assert.Equal(t, v, attrs[k])
	}
}

// AssertLink ...
func AssertLink(t *testing.T, i interface{}, names ...string) {
	if i == nil {
		assert.True(t, false, "Expected links got nil")
		return
	}
	links := i.(map[string]interface{})
	var link string
	for _, name := range names {
		if links[name] == nil {
			assert.True(t, false, fmt.Sprintf("Expected links to have '%s'", name))
			return
		}
		link = links[name].(string)
		assert.True(t, len([]rune(link)) > 0)
	}
}

// AssertNoLink ...
func AssertNoLink(t *testing.T, i interface{}, names ...string) {
	if i == nil {
		assert.True(t, false, "Expected links got nil")
		return
	}
	links := i.(map[string]interface{})
	for _, name := range names {
		if links[name] != nil {
			assert.True(t, false, fmt.Sprintf("Expected links not to have '%s'", name))
			return
		}
	}
}

// AssertAction ...
func AssertAction(t *testing.T, meta JSONAPIRespMeta, names ...string) {
	var includes bool
	for _, name := range names {
		includes = false
		for _, action := range meta.Actions {
			if action.Name == name {
				includes = true
			}
		}
		assert.True(t, includes)
	}
}

// AssertNoAction ...
func AssertNoAction(t *testing.T, meta JSONAPIRespMeta, names ...string) {
	var includes bool
	for _, name := range names {
		includes = false
		for _, action := range meta.Actions {
			if action.Name == name {
				includes = true
			}
		}
		assert.True(t, !includes)
	}
}

// AssertStatusSuccess assert response code is success
func AssertStatusSuccess(t *testing.T, resp gorequest.Response) {
	ok := resp.StatusCode >= 200 && resp.StatusCode <= 226
	if !ok {
		t.Errorf("Expected status 2xx but is %d\n%+v\n", resp.StatusCode, resp.Body)
	}
}

// AssertStatusOK assert response status code is OK
func AssertStatusOK(t *testing.T, resp gorequest.Response) {
	ok := resp.StatusCode == 200 || resp.StatusCode == 201
	if !ok {
		t.Errorf("Expected status 200/201 but is %d\n%+v\n", resp.StatusCode, resp.Body)
	}
}

// AssertStatusUnauthorized assert response status code is 401
func AssertStatusUnauthorized(t *testing.T, resp gorequest.Response) {
	ok := resp.StatusCode == 401
	if !ok {
		t.Errorf("Expected status 401 but is %d", resp.StatusCode)
	}
}

// AssertStatusNotFound assert response status code is 404
func AssertStatusNotFound(t *testing.T, resp gorequest.Response) {
	ok := resp.StatusCode == 404
	if !ok {
		t.Errorf("Expected status 404 but is %d", resp.StatusCode)
	}
}
