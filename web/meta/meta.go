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
	ActionHolder []*JsonAPIAction

	Extendable interface {
		Extend() error
	}

	JsonAPIAction struct {
		Method string         `json:"method"`
		Name   string         `json:"name"`
		URL    string         `json:"url"`
		Multi  bool           `json:"multi"`
		Fields []JsonAPIField `json:"fields"`
	}

	JsonAPIField struct {
		Name      string        `json:"name"`
		InputType string        `json:"type"`
		Value     interface{}   `json:"value,omitempty"`
		Required  bool          `json:"required"`
		Options   []FieldOption `json:"options,omitempty"`
		Data      *Pagination   `json:"data,omitempty"`
	}

	FieldOption struct {
		Label string                 `json:"label,omitempty"`
		Value interface{}            `json:"value"`
		Meta  map[string]interface{} `json:"meta,omitempty"`
	}

	Pagination struct {
		Count int `json:"item_count"`
		Max   int `json:"max"`
		Page  int `json:"page"`
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

// AddAction ...
func (ah *ActionHolder) AddAction(m method, name, urlHelper string, params ...interface{}) *JsonAPIAction {
	a := JsonAPIAction{
		Method: string(m),
		Name:   name,
		URL:    APIURL(fmt.Sprintf(urlHelper, params...)),
		Fields: make([]JsonAPIField, 0),
	}
	*ah = append(*ah, &a)
	return &a
}

// Field ...
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

// FieldWithOpts add field with options to action
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

// Pagination add pagination meta to action
func (a *JsonAPIAction) Pagination(data *Pagination) *JsonAPIAction {
	f := JsonAPIField{
		Name:      "page",
		InputType: string(InputNumber),
		Value:     data.Page,
		Required:  false,
		Data:      data,
	}
	a.Fields = append(a.Fields, f)
	return a
}

// RenderActions ...
func (ah *ActionHolder) RenderActions() *jsonapi.Meta {
	return &jsonapi.Meta{"actions": *ah}
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
