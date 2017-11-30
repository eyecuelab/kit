//Package maputil contains utility functions for working with map[string]interface{}
package maputil

import "sort"

//Keys returns the keys of the map as a slice. No order is guaranted.
func Keys(m map[string]interface{}) []string {
	keys := make([]string, len(m))
	var i int
	for k := range m {
		keys[i] = k
		i++
	}
	return keys
}

//Vals returns the values of the map as a slice. No order is guaranteed.
func Vals(m map[string]interface{}) []interface{} {
	vals := make([]interface{}, len(m))
	var i int
	for _, v := range m {
		vals[i] = v
		i++
	}
	return vals
}

//SortedKeys returns the keys of the map sorted in the usual (lexigraphic) order.
func SortedKeys(m map[string]interface{}) []string {
	keys := Keys(m)
	sort.Strings(keys)
	return keys
}

//Copy returns a copy of the map
func Copy(m map[string]interface{}) map[string]interface{} {
	copy := make(map[string]interface{})
	for k, v := range m {
		copy[k] = v
	}
	return copy
}
