package stringmap

import "sort"

//Keys return a slice containing the keys of the map[string]string. No order is guaranteed.
func Keys(m map[string]string) []string {
	keys := make([]string, len(m))
	var i int
	for k := range m {
		keys[i] = k
		i++
	}
	return keys
}

//Vals returns a slice containing the values of the map[string]string. No order is guaranteed. Note that values may not be unique.
func Vals(m map[string]string) []string {
	vals := make([]string, len(m))
	var i int
	for _, v := range m {
		vals[i] = v
		i++
	}
	return vals
}

//SortedKeys returns a slice containin the (sorted) keys of the map[string]string, in the usual lexigraphic order provided by sort.Strings
func SortedKeys(m map[string]string) []string {
	keys := Keys(m)
	sort.Strings(keys)
	return keys
}

//SortedVals returns a slice containin the (sorted) valuess of the map[string]string, in the usual lexigraphic order provided by sort.Strings. Values may not be unique.
func SortedVals(m map[string]string) []string {
	vals := Vals(m)
	sort.Strings(vals)
	return vals
}

//Copy returns a copy of the map[string]string.
func Copy(m map[string]string) map[string]string {
	copy := make(map[string]string)
	for k, v := range m {
		copy[k] = v
	}
	return copy
}
