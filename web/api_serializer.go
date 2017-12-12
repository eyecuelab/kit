package web

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/eyecuelab/kit/flect"
	"github.com/google/jsonapi"
	"github.com/jinzhu/inflection"
	"github.com/spf13/viper"
)

type (
	method       string
	inputType    string
	actionHolder []*jsonApiAction

	// jsonapi action
	jsonApiAction struct {
		Method string         `json:"method"`
		Name   string         `json:"name"`
		Url    string         `json:"url"`
		Fields []jsonApiField `json:"fields"`
	}

	// jsonapi action field
	jsonApiField struct {
		Name      string `json:"name"`
		InputType string `json:"type"`
		Value     string `json:"value,omitempty"`
		Required  bool   `json:"required"`
	}
)

const (
	DELETE method = "DELETE"
	GET    method = "GET"
	POST   method = "POST"

	InputText inputType = "text"
	InputPass inputType = "password"
)

var ah actionHolder

func AddAction(m method, name, urlHelper string) *jsonApiAction {
	a := jsonApiAction{
		Method: string(m),
		Name:   name,
		Url:    APIURL(urlHelper),
		Fields: make([]jsonApiField, 0),
	}
	ah = append(ah, &a)
	return &a
}

func (a *jsonApiAction) Field(name string, inputType inputType, value string, requred bool) *jsonApiAction {
	f := jsonApiField{
		Name: name,
		InputType: string(inputType),
		Value: value,
		Required: requred,
	}
	a.Fields = append(a.Fields, f)
	return a
}

func RenderActions() *jsonapi.Meta {
	clone := ah
	ah = []*jsonApiAction{}
	return &jsonapi.Meta{"actions": clone}
}

// APIURL full api url for the path
func APIURL(path string) string {
	return fmt.Sprintf("%s%s", viper.GetString("root_url"), path)
}

// JSONApiSelfLink self link helper
func JSONApiSelfLink(i interface{}) *jsonapi.Links {
	name := reflect.TypeOf(i).Name()
	name = inflection.Plural(strings.ToLower(name))

	id := reflect.ValueOf(i).FieldByName("ID").Interface()

	return &jsonapi.Links{
		"self": jsonapi.Link{
			Href: fmt.Sprintf("/%v/%v", name, id)},
	}
}

// JSONApiRefLink ref link helper
func JSONApiRefLink(i interface{}, rel string) *jsonapi.Links {
	fields := flect.Fields(reflect.ValueOf(i))
	name := reflect.TypeOf(i).Name()
	id := reflect.ValueOf(i).FieldByName("ID").Interface()
	for _, f := range fields {
		value, opts, ok := flect.TagValues(f.Tag, "jsonapi")
		if ok && value == "relation" && opts.HasOption(rel) {
			if _, _, ok := flect.TagValues(f.Tag, "link"); ok {
				return &jsonapi.Links{
					"related": fmt.Sprintf(
						"/%v/%v/%v",
						strings.ToLower(inflection.Plural(name)),
						id,
						strings.ToLower(inflection.Plural(f.Name)),
					),
				}
			}
		}
	}
	return &jsonapi.Links{}
}