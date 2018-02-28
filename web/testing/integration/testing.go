package integration

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/eyecuelab/kit/web/meta"
	"github.com/google/jsonapi"
	"github.com/parnurzeal/gorequest"
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
		ID            string
		Type          string
		Attributes    interface{}     `json:"attributes"`
		Meta          JSONAPIRespMeta `json:"meta"`
		Links         interface{}     `json:"links"`
		Relationships interface{}
	}
	// JSONAPIOneResp ...
	JSONAPIOneResp struct {
		Data     JSONAPIRespData        `json:"data"`
		Errors   []*jsonapi.ErrorObject `json:"errors"`
		Included []JSONAPIRespData
	}
	// JSONAPIManyResp ...
	JSONAPIManyResp struct {
		Data     []JSONAPIRespData      `json:"data"`
		Meta     JSONAPIRespMeta        `json:"meta"`
		Errors   []*jsonapi.ErrorObject `json:"errors"`
		Included []JSONAPIRespData
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

// Delete jsonapi delete request
func Delete(t *testing.T, path string, token string) (gorequest.Response, *JSONAPIOneResp, []error) {
	resp, body, errs := Request("DELETE", path, token).End()

	return resp, unmarshalOne(t, body), errs
}

// GetList jsonapi get list request
func GetList(t *testing.T, path string, token string) (gorequest.Response, *JSONAPIManyResp, []error) {
	resp, body, errs := Request("GET", path, token).End()
	data := unmarshalMany(t, body)

	return resp, data, errs
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
	if body == "" {
		return &data
	}
	err := json.Unmarshal([]byte(body), &data)
	if err != nil {
		t.Errorf("Error unmarshaling response %v", err)
	}

	return &data
}

func unmarshalMany(t *testing.T, body string) *JSONAPIManyResp {
	var data JSONAPIManyResp
	err := json.Unmarshal([]byte(body), &data)
	if err != nil {
		t.Errorf("Error unmarshaling response %v", err)
	}

	return &data
}
