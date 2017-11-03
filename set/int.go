package set

import "sort"

//Int is a set of ints
type Int map[int]signal

//Contains shows whether a given key is in the Int.
func (s Int) Contains(key int) bool {
	_, ok := s[key]
	return ok
}

//Copy copies the set of ints.
func (s Int) Copy() Int {
	copy := make(Int)
	for k := range s {
		copy[k] = yes
	}
	return copy
}

//Intersection returns a new set containing the intersection of the ints;
func (s Int) Intersection(ints ...Int) (intersection Int) {
	intersection = s.Copy()
	for key := range s {
		for _, set := range append(ints, s) {
			if !set.Contains(key) {
				delete(intersection, key)
			}
		}
	}
	return intersection
}

//XOR returns the keys in one set but not the other.
func (s Int) XOR(other Int) Int {
	union := s.Union(other)
	xor := make(Int)
	for k := range union {
		if (s.Contains(k) && !other.Contains(k)) || (other.Contains(k) && !s.Contains(k)) {
			xor[k] = yes
		}
	}
	return xor
}

//Equal shows whether two Ints are equal; i.e, they contain the same items.
func (s Int) Equal(other Int) bool {
	if len(s) != len(other) {
		return false
	}
	for key := range s {
		if !other.Contains(key) {
			return false
		}
	}
	return true
}

//Union returns a new set containing the union of the int sets.
func (s Int) Union(ints ...Int) (union Int) {
	union = s.Copy()
	for _, set := range ints {
		for key := range set {
			union[key] = yes
		}
	}
	return union
}

//IUnion modifies the Int in place rather than returning a copy.
func (s Int) IUnion(ints ...Int) {
	for _, set := range ints {
		for key := range set {
			s[key] = yes
		}
	}
}

//Difference returns the items in the reciever but not any other arguments
func (s Int) Difference(ints ...Int) (difference Int) {
	difference = s.Copy()
	for _, set := range ints {
		for key := range set {
			delete(difference, key)
		}
	}
	return difference
}

//FromInts creates a set from ints
func FromInts(ints ...int) Int {
	s := make(Int)
	for _, key := range ints {
		s[key] = yes
	}
	return s
}

//Add a key or key(s) to the set, in-place. Don't call on a nil set.
func (s Int) Add(keys ...int) {
	for _, key := range keys {
		s[key] = yes
	}
}

//Remove a key from the set, in-place.
func (s Int) Remove(keys ...int) {
	for _, key := range keys {
		delete(s, key)
	}
}

//ToSlice returns a slice containing the keys of a set. No order is guaranteed.
func (s Int) ToSlice() []int {
	slice := make([]int, len(s))
	var i = 0
	for k := range s {
		slice[i] = k
		i++
	}
	return slice
}

//Sorted returns a slice containing the sorted keys of the set.
func (s Int) Sorted() []int {
	slice := s.ToSlice()
	sort.Ints(slice)
	return slice
}
