package meta

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/eyecuelab/jsonapi"
	"github.com/eyecuelab/kit/web/pagination"
)

type (
	method       string
	inputType    string
	idType       string
	ActionHolder []*JsonAPIAction

	Extendable interface {
		Extend() error
	}

	JsonAPIAction struct {
		Method       string         `json:"method"`
		Name         string         `json:"name"`
		URL          string         `json:"url"`
		Multi        bool           `json:"multi"`
		Relationship bool           `json:"relationship"`
		Fields       []JsonAPIField `json:"fields"`
	}

	JsonAPIField struct {
		Name      string        `json:"name"`
		InputType string        `json:"type"`
		Value     interface{}   `json:"value,omitempty"`
		Required  bool          `json:"required"`
		Options   []FieldOption `json:"options,omitempty"`
		Data      *pagination.Pagination   `json:"data,omitempty"`
	}

	FieldOption struct {
		Label string                 `json:"label,omitempty"`
		Value interface{}            `json:"value"`
		Meta  map[string]interface{} `json:"meta,omitempty"`
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

	IdString idType = "string"
	IdInt    idType = "int"
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

// RelFields ...
func (a *JsonAPIAction) RelFields(typeName string) *JsonAPIAction {
	fs := []JsonAPIField{
		{
			Name:      "type",
			Value:     typeName,
			Required:  true,
		},
		{
			Name:      "id",
			Required:  true,
		},
	}
	a.Relationship = true
	a.Fields = append(a.Fields, fs...)
	return a
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

func FieldOptionsFromValues(values ...string) []FieldOption{
	opts := make([]FieldOption, len(values))
	for i, v := range values {
		opts[i] = FieldOption{Value: v}
	}
	return opts
}

// Pagination add pagination meta to action
func (a *JsonAPIAction) Pagination(data *pagination.Pagination) *JsonAPIAction {
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