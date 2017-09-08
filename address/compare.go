package address

import "github.com/eyecuelab/kit/stringslice"

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
