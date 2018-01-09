package sortlib

import (
	"sort"
)

//ByteSlice is a sortable slice of bytes.
type ByteSlice []byte

func (b ByteSlice) Less(i, j int) bool {
	return b[i] < b[j]
}

func (b ByteSlice) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b ByteSlice) Len() int {
	return len(b)
}

type RuneSlice []rune

func (r RuneSlice) Less(i, j int) bool {
	return r[i] < r[j]
}

func (r RuneSlice) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r RuneSlice) Len() int {
	return len(r)
}

//Bytes returns a sorted copy of the string, in ascending order (by bytes).
func ByBytes(s string) string {
	b := ByteSlice(s)
	sort.Sort(b)
	return string(b)
}

//Runes returns a sorted copy of the string, in ascending order (by runes).
func ByRunes(s string) string {
	r := RuneSlice(s)
	sort.Sort(r)
	return string(r)
}
