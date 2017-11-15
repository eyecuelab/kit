package stringmap

import "sort"

func Keys(m map[string]string) []string {
	keys := make([]string, len(m))
	var i int
	for k := range m {
		keys[i] = k
		i++
	}
	return keys
}

func Vals(m map[string]string) []string {
	vals := make([]string, len(m))
	var i int
	for _, v := range m {
		vals[i] = v
		i++
	}
	return vals
}

func SortedKeys(m map[string]string) []string {
	keys := Keys(m)
	sort.Strings(keys)
	return keys
}

func Copy(m map[string]string) map[string]string {
	copy := make(map[string]string)
	for k, v := range m {
		copy[k] = v
	}
	return copy
}
