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
	a, b = str.Map(normalize, a), str.Map(normalize, b)
	distance, _ = stringslice.TotalDistance(a, b)
	return distance
}

//NormalizedSharedComponentDistance performs heavy normalization on both addresses, then compares the distance.
func NormalizedSharedComponentDistance(placeA, placeB *Address) (distance int) {
	filteredA, filteredB := SharedComponents(*placeA, *placeB)
	a, b := filteredA.StringSliceFromNonempty(), filteredB.StringSliceFromNonempty()
	a, b = str.Map(normalize, a), str.Map(normalize, b)
	distance, _ = stringslice.TotalDistance(a, b)
	return distance
}
