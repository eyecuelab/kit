package address

import (
	"github.com/eyecuelab/kit/str"
	"github.com/eyecuelab/kit/stringslice"
)

//TotalDistance returns the sum of the levenshtein distance of the components of the addresses.
func TotalDistance(placeA, placeB *Address) (distance int) {
	a, b := placeA.StringSlice(), placeB.StringSlice()
	distance, _ = stringslice.TotalDistance(a, b)
	return distance
}

//SharedComponentDistance returns the sum of the levenshtein distance of the shared components of the addresses
func SharedComponentDistance(placeA, placeB *Address) (distance int) {
	filteredA, filteredB := SharedComponents(*placeA, *placeB)
	a, b := filteredA.StringSliceFromNonempty(), filteredB.StringSliceFromNonempty()
	distance, _ = stringslice.TotalDistance(a, b)
	return distance
}

//NormalizedTotalDistance Performs heavy normalization on both addresses, then compares the distance.
func NormalizedTotalDistance(placeA, placeB *Address) (distance int) {
	filteredA, filteredB := SharedComponents(*placeA, *placeB)
	a, b := filteredA.StringSlice(), filteredB.StringSlice()
	a, b = str.Map(str.ExtremeNormalization, a), str.Map(str.ExtremeNormalization, b)
	distance, _ = stringslice.TotalDistance(a, b)
	return distance
}

//NormalizedSharedComponentDistance performs heavy normalization on both addresses, then compares the distance.
func NormalizedSharedComponentDistance(placeA, placeB *Address) (distance int) {
	filteredA, filteredB := SharedComponents(*placeA, *placeB)
	a, b := filteredA.StringSliceFromNonempty(), filteredB.StringSliceFromNonempty()
	a, b = str.Map(str.ExtremeNormalization, a), str.Map(str.ExtremeNormalization, b)
	distance, _ = stringslice.TotalDistance(a, b)
	return distance
}

func NormalizedSharedComponentDistanceSlice(placeA, placeB Address) (distances []int) {
	filteredA, filteredB := SharedComponents(placeA, placeB)
	a, b := filteredA.StringSliceFromNonempty(), filteredB.StringSliceFromNonempty()
	a, b = str.Map(str.ExtremeNormalization, a), str.Map(str.ExtremeNormalization, b)
	distances = make([]int, len(a))
	for i := range distances {
		distances[i] = levenshteinDistance(a[i], b[i])
	}
	return distances
}

func NonOverlappingComponents(placeA, placeB Address) (uniqueToA, uniqueToB int) {
	filteredA, filteredB := SharedComponents(placeA, placeB)
	uniqueToA = len(placeA.StringSliceFromNonempty()) - len(filteredA.StringSliceFromNonempty())
	uniqueToB = len(placeB.StringSliceFromNonempty()) - len(filteredB.StringSliceFromNonempty())
	return uniqueToA, uniqueToB
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
		min = min2(min, n)
	}
	return min
}

func min2(a, b int) int {
	if a < b {
		return a
	}
	return b
}
