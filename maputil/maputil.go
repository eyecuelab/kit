package maputil

import "sort"

func Keys(m map[string]interface{}) []string {
	keys := make([]string, len(m))
	var i int
	for k := range m {
		keys[i] = k
		i++
	}
	return keys
}

func Vals(m map[string]interface{}) []interface{} {
	vals := make([]interface{}, len(m))
	var i int
	for _, v := range m {
		vals[i] = v
		i++
	}
	return vals
}

func SortedKeys(m map[string]interface{}) []string {
	keys := Keys(m)
	sort.Strings(keys)
	return keys
}
