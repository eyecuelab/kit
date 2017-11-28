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

// JSONApiAction jsonapi action
type JSONApiAction struct {
	Method string         `json:"method"`
	Name   string         `json:"name"`
	ID     string         `json:"id"`
	Type   string         `json:"type"`
	URL    string         `json:"url"`
	Fields []JSONApiField `json:"fields"`
}

// JSONApiField jsonapi action field
type JSONApiField struct {
	Name      string      `json:"name"`
	InputType string      `json:"type"`
	Value     interface{} `json:"value"`
	Required  bool        `json:"required"`
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
