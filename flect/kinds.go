package flect

import "reflect"

func IsSlice(i interface{}) bool {
	return IsA(i, reflect.Slice)
}

func IsA(i interface{}, kind reflect.Kind) bool {
	return reflect.TypeOf(i).Kind() == kind
}

func NotA(i interface{}, kind reflect.Kind) bool {
	return !IsA(i, kind)
}