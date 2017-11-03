package set

import "sort"

//String is a set of strings
type String map[string]signal

//Contains shows whether a given key is in the String.
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

//Remove a key from the set, in-place.
func (s String) Remove(keys ...string) {
	for _, key := range keys {
		delete(s, key)
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

func (s String) Sorted() []string {
	slice := s.ToSlice()
	sort.Strings(slice)
	return slice
}
