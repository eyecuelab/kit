package sortlib

import (
	"sort"

	"github.com/eyecuelab/kit/copyslice"
	"github.com/eyecuelab/kit/sortlib/sortable"
)

//Int creates a sorted copy of the slice of ints.
func Int(a []int) []int {
	b := copyslice.Int(a)
	sort.Ints(b)
	return b
}

//Uint creates a sorted copy of the slice of uints.
func Uint(a []uint) []uint {
	b := copyslice.Uint(a)
	sort.Sort(sortable.Uints(b))
	return b
}

//Float64 creates a sorted copy of the slice of float64s
func Float64(a []float64) []float64 {
	b := copyslice.Float64(a)
	sort.Float64s(b)
	return b
}

//Int64 creates a sorted copy of the slice of int64s
func Int64(a []int64) []int64 {
	b := copyslice.Int64(a)
	sort.Sort(sortable.Int64s(b))
	return b
}

//Uint64 creates a sorted copy of the slice of uint64s
func Uint64(a []uint64) []uint64 {
	b := copyslice.Uint64(a)
	sort.Sort(sortable.Uint64s(b))
	return b
}

//Bytes creates a sorted copy of the slice of bytes
func Bytes(a []byte) []byte {
	b := copyslice.Byte(a)
	sort.Sort(sortable.Bytes(b))
	return b
}

//Runes creates a sorted copy of the slice of runes
func Runes(a []rune) []rune {
	b := copyslice.Rune(a)
	sort.Sort(sortable.Runes(b))
	return b
}
