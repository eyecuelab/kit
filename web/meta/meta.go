package meta

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
	actionHolder []*JsonAPIAction

	JsonAPIAction struct {
		Method string         `json:"method"`
		Name   string         `json:"name"`
		URL    string         `json:"url"`
		Fields []JsonAPIField `json:"fields"`
	}

	JsonAPIField struct {
		Name      string        `json:"name"`
		InputType string        `json:"type"`
		Value     interface{}   `json:"value,omitempty"`
		Required  bool          `json:"required"`
		Options   []FieldOption `json:"options,omitempty"`
	}

	FieldOption struct {
		Label string      `json:"label"`
		Value interface{} `json:"value"`
	}
)

const (
	DELETE method = "DELETE"
	GET    method = "GET"
	POST   method = "POST"
	PATCH  method = "PATCH"
	PUT    method = "PUT"

	InputText   inputType = "text"
	InputPass   inputType = "password"
	InputBool   inputType = "bool"
	InputNumber inputType = "number"
	InputSelect inputType = "select"
)

var ah actionHolder

func AddAction(m method, name, urlHelper string) *JsonAPIAction {
	a := JsonAPIAction{
		Method: string(m),
		Name:   name,
		URL:    APIURL(urlHelper),
		Fields: make([]JsonAPIField, 0),
	}
	ah = append(ah, &a)
	return &a
}

func (a *JsonAPIAction) Field(name string, inputType inputType, value interface{}, requred bool) *JsonAPIAction {
	f := JsonAPIField{
		Name:      name,
		InputType: string(inputType),
		Value:     value,
		Required:  requred,
	}
	a.Fields = append(a.Fields, f)
	return a
}

func (a *JsonAPIAction) FieldWithOpts(name string, inputType inputType, value interface{}, requred bool, options []FieldOption) *JsonAPIAction {
	f := JsonAPIField{
		Name:      name,
		InputType: string(inputType),
		Value:     value,
		Required:  requred,
		Options:   options,
	}
	a.Fields = append(a.Fields, f)
	return a
}

func RenderActions() *jsonapi.Meta {
	clone := ah
	ah = []*JsonAPIAction{}
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
