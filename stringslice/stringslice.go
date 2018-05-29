//Package stringslice contains various functions to work with slices of strings and the strings contained within.
package stringslice

import "sort"

type predicate func(string) bool

//IndexOf returns the index of the search string.  -1 if not found
func IndexOf(a []string, search string) int {
	for i, s := range a {
		if s == search {
			return i
		}
	}
	return -1
}


//NonEmpty returns a slice containing the non-empty elements of a
func NonEmpty(a []string) []string {
	var filtered []string
	return AppendIfNonEmpty(filtered, a...)
}

//All returns true if f(s) is true for all s in a. All([], f) is true.
func All(a []string, f predicate) bool {
	for _, s := range a {
		if !f(s) {
			return false
		}
	}
	return true
}

//Copy makes a copy of the []string.
func Copy(src []string) []string {
	dst := make([]string, len(src))
	copy(dst, src)
	return dst
}

//Any (returns true if f(s) is true for any s in a. Any([], f) is false.
func Any(a []string, f func(string) bool) bool {
	for _, s := range a {
		if f(s) {
			return true
		}
	}
	return false
}

//Filter returns the elements of a where f(a) is true.
func Filter(a []string, f func(string) bool) []string {
	filtered := make([]string, 0, len(a))
	for _, s := range a {
		if f(s) {
			filtered = append(filtered, s)
		}
	}
	return filtered
}

//FilterFalse returns the elements of a where f(a) is false.
func FilterFalse(a []string, f func(string) bool) []string {
	filtered := make([]string, 0, len(a))
	for _, s := range a {
		if !f(s) {
			filtered = append(filtered, s)
		}
	}
	return filtered
}

//Map maps a function across the slice, returning a new slice containing f(s) for s in a
func Map(a []string, f func(string) string) []string {
	mapped := make([]string, len(a))
	for i, s := range a {
		mapped[i] = f(s)
	}
	return mapped
}

//AppendIfNonEmpty appends the given string(s) to the slice if they are not the empty string.
func AppendIfNonEmpty(a []string, strings ...string) []string {
	for _, s := range strings {
		if s != "" {
			a = append(a, s)
		}
	}
	return a
}

//Combine multiple slices into a single slice. This is equivalent to appending each slice in turn.
func Combine(slices ...[]string) []string {
	var L int
	for _, a := range slices {
		L += len(a)
	}
	combined := make([]string, 0, L)
	for _, a := range slices {
		combined = append(combined, a...)
	}
	return combined
}

//Sorted takes any number of input strings and returns a sorted stringslice.
func Sorted(a ...string) []string {
	sort.Strings(a)
	return a
}

//TotalDistance returns the sum of the levanshtein distance of the strings. If the string slices are not the same length, OK is false.
func TotalDistance(a []string, b []string) (distance int, ok bool) {
	if len(a) != len(b) {
		return 0, false
	}
	for i := range a {
		distance += levenshteinDistance(a[i], b[i])
	}
	return distance, true
}

// levenshteinDistance measures the difference between two strings.
// The Levenshtein distance between two words is the minimum number of
// single-character edits (i.e. insertions, deletions or substitutions)
// required to change one word into the other.
//
// This implemention is optimized to use O(min(m,n)) space and is based on the
// optimized C version found here:
// http://en.wikibooks.org/wiki/Algorithm_implementation/Strings/Levenshtein_distance#C
// This implentation Copyright (c) 2015 Peter Renström under the MIT license: https://github.com/renstrom/fuzzysearch/blob/master/LICENSE
func levenshteinDistance(s, t string) int {
	r1, r2 := []rune(s), []rune(t)
	column := make([]int, len(r1)+1)

	for y := 1; y <= len(r1); y++ {
		column[y] = y
	}

	for x := 1; x <= len(r2); x++ {
		column[0] = x

		for y, lastDiag := 1, x-1; y <= len(r1); y++ {
			oldDiag := column[y]
			cost := 0
			if r1[y-1] != r2[x-1] {
				cost = 1
			}
			column[y] = min(column[y]+1, column[y-1]+1, lastDiag+cost)
			lastDiag = oldDiag
		}
	}
	return column[len(r1)]
}

func min(a int, ints ...int) int {
	min := a
	for _, n := range ints {
		if n < min {
			min = n
		}
	}
	return min
}
