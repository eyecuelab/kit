//Package flect is meant to work alongside go's `reflect` library, containing additional runtime reflection tools
package flect

import (
	"reflect"
	"strings"

	"github.com/eyecuelab/kit/functools"
)

type tagOpts []string

type groupConfig int
const (
	WithZeros groupConfig = iota
	WithoutZeros
)

func (opts tagOpts) HasOption(name string) bool {
	return functools.StringSliceContains(opts, name)
}

func TagValues(t reflect.StructTag, name string) (value string, opts tagOpts, ok bool) {
	fullValue, ok := t.Lookup(name)
	if !ok {
		return
	}
	splitValue := strings.Split(fullValue, ",")
	if len(splitValue) > 0 {
		opts = tagOpts(splitValue[1:])
	}
	return splitValue[0], opts, true
}

func Values(structs ...interface{}) []reflect.Value {
	values := make([]reflect.Value, len(structs))
	for i, s := range structs {
		values[i] = reflect.ValueOf(s)
	}
	return values
}

func Fields(val reflect.Value) []reflect.StructField {
	fields := make([]reflect.StructField, val.NumField())
	for i := range fields {
		fields[i] = val.Type().Field(i)
	}
	return fields
}

func ValuesByTag(tag string, config groupConfig, structs ...interface{}) map[string]interface{} {
	tagValues := make(map[string]interface{})
	for _, data := range structs {
		for _, val := range Values(data) {
			for j, field := range Fields(val) {
				tagValue, _, ok := TagValues(field.Tag, "attr")
				if ok {
					attrValue := val.Field(j).Interface()
					if config == WithoutZeros && IsZeroOfType(attrValue) {
						continue
					}
					tagValues[tagValue] = attrValue
				}
			}
		}
	}
	return tagValues
}

func IsZeroOfType(x interface{}) bool {
	return reflect.DeepEqual(x, reflect.Zero(reflect.TypeOf(x)).Interface())
}
