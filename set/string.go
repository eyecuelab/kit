package set

import "sort"

//String is a set of strings
type String map[string]signal

//Contains returns true if the given key is in the set.
func (s String) Contains(key string) bool {
	_, ok := s[key]
	return ok
}

//Copy copies the set of strings.
func (s String) Copy() String {
	copy := make(String)
	for k := range s {
		copy[k] = yes
	}
	return copy
}

//Intersection returns a new set containing the intersection of the strings;
func (s String) Intersection(strings ...String) (intersection String) {
	intersection = s.Copy()
	for key := range s {
		for _, set := range append(strings, s) {
			if !set.Contains(key) {
				delete(intersection, key)
			}
		}
	}
	return intersection
}

//XOR returns the keys in one set but not the other.
func (s String) XOR(other String) String {
	union := s.Union(other)
	xor := make(String)
	for k := range union {
		if (s.Contains(k) && !other.Contains(k)) || (other.Contains(k) && !s.Contains(k)) {
			xor[k] = yes
		}
	}
	return xor
}

//Equal shows whether two Strings are equal; i.e, they contain the same items.
func (s String) Equal(other String) bool {
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

//Union returns a new set containing the union of the string sets.
func (s String) Union(strings ...String) (union String) {
	union = s.Copy()
	for _, set := range strings {
		for key := range set {
			union[key] = yes
		}
	}
	return union
}

//IUnion modifies the StringSet in place rather than returning a copy.
func (s String) IUnion(strings ...String) {
	for _, set := range strings {
		for key := range set {
			s[key] = yes
		}
	}
}

//Difference returns the items in the reciever but not any other arguments
func (s String) Difference(strings ...String) (difference String) {
	difference = s.Copy()
	for _, set := range strings {
		for key := range set {
			delete(difference, key)
		}
	}
	return difference
}

//FromStrings creates a set from strings
func FromStrings(strings ...string) String {
	s := make(String)
	for _, key := range strings {
		s[key] = yes
	}
	return s
}

//Add a key or key(s) to the set, in-place. Don't call on a nil set.
func (s String) Add(keys ...string) {
	for _, key := range keys {
		s[key] = yes
	}
}

//ToSlice returns a slice containing the keys of a set. No order is guaranteed.
func (s String) ToSlice() []string {
	slice := make([]string, len(s))
	var i = 0
	for k := range s {
		slice[i] = k
		i++
	}
	return slice
}

//Sorted returns a slice containing the keys of the set in lexigraphic order.
func (s String) Sorted() []string {
	slice := s.ToSlice()
	sort.Strings(slice)
	return slice
}

//Map a function f(x) across the set, returning a new set containing f(x) for all x in the set.
func (s String) Map(f func(string) string) String {
	mapped := make(String)
	for k := range s {
		mapped[f(k)] = yes
	}
	return mapped

}

//Reduce applies a reducing function across the set. It will return (0, false) for a set with zero entries.
func (s String) Reduce(f func(string, string) string) (string, bool) {
	if len(s) == 0 {
		return "", false
	}
	first := true
	var reduced string
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
func (s String) Filter(f func(string) bool) String {
	filtered := make(String)
	for k := range s {
		if f(k) {
			filtered.Add(k)
		}
	}
	return filtered
}

//IsSubset returns true if every key in s is also in other
func (s String) IsSubset(other String) bool {
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
func (s String) Remove(key string) bool {
	_, ok := s[key]
	delete(s, key)
	return ok
}

//Pop an arbitrary element from the set. Returns "", false if no more elements remain. No order, or lack of order, is guaranteed.
func (s String) Pop() (k string, more bool) {
	for k := range s {
		//iterate
		delete(s, k)
		return k, true
	}
	return "", false
}

//IsDisjoint returns true if s shared no elements with other. Note that the empty set is disjoint with everything.
func (s String) IsDisjoint(other String) bool {
	for k := range s {
		if _, ok := other[k]; ok {
			return false
		}
	}
	return true
}

//IsProperSubset returns true if every key in s is also in other and s != other
func (s String) IsProperSubset(other String) bool {
	return len(s) < len(other) && s.IsSubset(other)
}

//IsSuperset returns true if every key in other is also in s
func (s String) IsSuperset(other String) bool {
	return other.IsSubset(s)
}

//IsProperSuperset returns true if every key in other is also in s and s != other
func (s String) IsProperSuperset(other String) bool {
	return len(s) > len(other) && other.IsSuperset(s)
}
