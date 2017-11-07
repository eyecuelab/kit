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
	for i := 1; i < len(splitValue); i++ {
		opts = append(opts, splitValue[i])
	}
	return splitValue[0], opts, true
}

func GroupValuesByTagOption(tag string, excludeEmpties bool, structs ...interface{}) map[string]map[string]interface{} {
	toScan := make([]reflect.Value, len(structs))
	optsMap := make(map[string]map[string]interface{})

	for i, s := range structs {
		toScan[i] = reflect.ValueOf(s)
	}

	for _, scan := range toScan {
		for i := 0; i < scan.NumField(); i++ {
			tagValue, tagOpts, ok := TagValues(scan.Type().Field(i).Tag, tag)
			if ok {
				attrValue := scan.Field(i).Interface()
				if IsZeroOfType(attrValue) && excludeEmpties {
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