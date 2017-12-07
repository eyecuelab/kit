package set

import "sort"

//Int is a set of ints. Remember to initialize with make(Int) or Int{}
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

//Delete a key or keys from the set, in-place.
func (s Int) Delete(keys ...int) {
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

//Sorted returns a slice containing the sorted keys of the set in the usual order.
func (s Int) Sorted() []int {
	slice := s.ToSlice()
	sort.Ints(slice)
	return slice
}

//Map a function f(x) across the set, returning a new set containing f(x) for all x in the set.
func (s Int) Map(f func(int) int) Int {
	mapped := make(Int)
	for k := range s {
		mapped[f(k)] = yes
	}
	return mapped

}

//Reduce applies a reducing function across the set. It will return (0, false) for a set with zero entries.
func (s Int) Reduce(f func(int, int) int) (int, bool) {
	if len(s) == 0 {
		return 0, false
	}
	first := true
	var reduced int
	for k := range s {
		if first {
			reduced = k
			first = false

		} else {
			reduced = f(reduced, k)
		}
	}
	return reduced, true
}

//Filter applies a filtering function across the set, returning a set containing x where f(x) is true.
func (s Int) Filter(f func(int) bool) Int {
	filtered := make(Int)
	for k := range s {
		if f(k) {
			filtered.Add(k)
		}
	}
	return filtered
}

//IsSubset returns true if s⊆other; every key in s is also in other
func (s Int) IsSubset(other Int) bool {
	if len(s) > len(other) {
		return false
	}

	for k := range s {
		if _, ok := other[k]; !ok {
			return false
		}

	}
	return true
}

//Remove removes a key from the set, returning the presence of the key.
func (s Int) Remove(key int) bool {
	_, ok := s[key]
	delete(s, key)
	return ok
}

//Pop an arbitrary element from the set. Returns 0, false if no more elements remain. No order (or lack of order) is guaranteed.
func (s Int) Pop() (k int, more bool) {
	for k := range s {
		delete(s, k)
		return k, true
	}
	return 0, false
}

//IsDisjoint returns true if s∩other == Ø; that is, s shares no elements with other. Note that the empty set is disjoint with everything.
func (s Int) IsDisjoint(other Int) bool {
	for k := range s {
		if _, ok := other[k]; ok {
			return false
		}
	}
	return true
}

//IsProperSubset returns true if every s ⊂ other: every key in s is also in other and s != other
func (s Int) IsProperSubset(other Int) bool {
	return len(s) < len(other) && s.IsSubset(other)
}

//IsSuperset returns true if other ⊆ s; every key in other is also in s
func (s Int) IsSuperset(other Int) bool {
	return other.IsSubset(s)
}

//IsProperSuperset returns true if other ⊂ s; every key in other is also in s and s != other
func (s Int) IsProperSuperset(other Int) bool {
	return len(s) > len(other) && other.IsSuperset(s)
}
