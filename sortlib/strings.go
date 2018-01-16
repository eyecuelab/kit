package sortlib

import (
	"sort"

	"github.com/eyecuelab/kit/sortlib/sortable"
)

//Bytes returns a sorted copy of the string, in ascending order (by bytes).
func ByBytes(s string) string {
	b := sortable.Bytes(s)
	sort.Sort(b)
	return string(b)
}

//Runes returns a sorted copy of the string, in ascending order (by runes).
func ByRunes(s string) string {
	r := sortable.Runes(s)
	sort.Sort(r)
	return string(r)
}
