package flect

import (
	"reflect"
	"strings"

	"github.com/eyecuelab/kit/functools"
)

type tagOpts []string

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

func GroupValuesByTagOption(tag string, structs ...interface{}) map[string]map[string]interface{} {
	optsMap := make(map[string]map[string]interface{})

	for _, val := range Values(structs...) {
		for i, field := range Fields(val) {
			tagValue, tagOpts, ok := TagValues(field.Tag, tag)
			if ok {
				attrValue := val.Field(i).Interface()
				for _, tagOpt := range tagOpts {
					if optsMap[tagOpt] == nil {
						optsMap[tagOpt] = make(map[string]interface{})
					}
					optsMap[tagOpt][tagValue] = attrValue
				}
			}
		}
	}
	return optsMap
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

func GroupNonEmptyValuesByTagOption(tag string, structs ...interface{}) map[string]map[string]interface{} {
	optsMap := make(map[string]map[string]interface{})
	for _, val := range Values(structs...) {
		for i, field := range Fields(val) {
			tagValue, tagOpts, ok := TagValues(field.Tag, tag)
			if ok {
				attrValue := val.Field(i).Interface()
				if IsZeroOfType(attrValue) {
					continue
				}
				for _, tagOpt := range tagOpts {
					if optsMap[tagOpt] == nil {
						optsMap[tagOpt] = make(map[string]interface{})
					}
					optsMap[tagOpt][tagValue] = attrValue
				}
			}
		}
	}
	return optsMap
}

func IsZeroOfType(x interface{}) bool {
	return reflect.DeepEqual(x, reflect.Zero(reflect.TypeOf(x)).Interface())
}
