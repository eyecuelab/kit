package integration

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/eyecuelab/kit/web/meta"
	"github.com/parnurzeal/gorequest"
	"github.com/stretchr/testify/assert"
)

const (
	apiURL = "http://localhost:3000"
)

type (
	// JSONAPIRespMeta ...
	JSONAPIRespMeta struct {
		Actions []meta.JsonAPIAction
	}
	// JSONAPIRespData ...
	JSONAPIRespData struct {
		Attributes interface{}     `json:"attributes"`
		Meta       JSONAPIRespMeta `json:"meta"`
		Links      interface{}     `json:"links"`
	}
	// JSONAPIOneResp ...
	JSONAPIOneResp struct {
		Data JSONAPIRespData `json:"data"`
	}
	// JSONAPIManyResp ...
	JSONAPIManyResp struct {
		Data []JSONAPIRespData `json:"data"`
		Meta JSONAPIRespMeta   `json:"meta"`
	}
)

// Request generic json api request with optional auth token
func Request(method string, path string, token string) (req *gorequest.SuperAgent) {
	r := gorequest.New().
		CustomMethod(method, apiURL+path).
		Set("Content-Type", "application/vnd.api+json")
	if token != "" {
		r = r.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	}
	return r
}

// Get jsonapi get request
func Get(t *testing.T, path string, token string) (gorequest.Response, *JSONAPIOneResp, []error) {
	resp, body, errs := Request("GET", path, token).End()
	data := unmarshalOne(t, body)

	return resp, data, errs
}

// Post jsonapi post request
func Post(t *testing.T, path string, attrs map[string]interface{}, token string) (gorequest.Response, *JSONAPIOneResp, []error) {
	resp, body, errs := Request("POST", path, token).Send(jsonAPIPayload(t, attrs)).End()
	data := unmarshalOne(t, body)

	return resp, data, errs
}

// Patch jsonapi patch request
func Patch(t *testing.T, path string, attrs map[string]interface{}, token string) (gorequest.Response, *JSONAPIOneResp, []error) {
	resp, body, errs := Request("PATCH", path, token).Send(jsonAPIPayload(t, attrs)).End()
	data := unmarshalOne(t, body)

	return resp, data, errs
}

// AssertGetOK assert GET request is OK
func AssertGetOK(t *testing.T, path string, token string) (resp gorequest.Response, body string, errs []error) {
	resp, body, errs = Request("GET", path, token).End()
	assert.Empty(t, errs)
	AssertStatusOK(t, resp)

	return resp, body, errs
}

// AssertAttr ...
func AssertAttr(t *testing.T, i interface{}, m map[string]interface{}) {
	attrs := i.(map[string]interface{})
	for k, v := range m {
		assert.Equal(t, v, attrs[k])
	}
}

// AssertLink ...
func AssertLink(t *testing.T, i interface{}, names ...string) {
	links := i.(map[string]interface{})
	var link string
	for _, name := range names {
		link = links[name].(string)
		assert.True(t, len([]rune(link)) > 0)
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

// AssertStatusOK assert response status code is OK
func AssertStatusOK(t *testing.T, resp gorequest.Response) {
	ok := resp.StatusCode == 200 || resp.StatusCode == 201
	if !ok {
		t.Errorf("Expected status 200/201 but is %d", resp.StatusCode)
	}
}

func jsonAPIPayload(t *testing.T, attrs map[string]interface{}) (payload string) {
	b, err := json.Marshal(attrs)
	if err != nil {
		t.Errorf("Error marshaling payload %v", err)
	}

	return fmt.Sprintf(`
	  {
	    "data": {
	      "attributes": %s
	    }
	  }
	`, b)
}

func unmarshalOne(t *testing.T, body string) *JSONAPIOneResp {
	var data JSONAPIOneResp
	err := json.Unmarshal([]byte(body), &data)
	if err != nil {
		t.Errorf("Error unmarshaling response %v", err)
	}

	return &data
}
