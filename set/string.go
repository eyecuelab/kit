package set

type String map[string]bool

//Contains shows whether strS is in the String.
func (s String) Contains(strS string) bool {
	_, ok := s[strS]
	return ok
}

func (s String) Copy() String {
	copy := make(String)
	for k, v := range s {
		copy[k] = v
	}
	return copy
}

//Intersection returns the intersection of the strings;
func (s String) Intersection(strings ...String) (intersection String) {
	intersection = s.Copy()
	for _, set := range strings {
		for strS, ok := range set {
			if !ok {
				delete(intersection, strS)
			}
		}
	}
	return intersection
}

//Equal shows whether two Strings are equal; i.e, they contain the same items.
func (s String) Equal(other String) bool {
	for strS := range s {
		if !other.Contains(strS) {
			return false
		}
	}
	return true
}

//Union returns the union of the strings.
func (s String) Union(strings ...String) (union String) {
	union = s.Copy()
	for _, set := range strings {
		for strS := range set {
			union[strS] = true
		}
	}
	return union
}

//Difference returns the items in the reciever but not any other arguments
func (s String) Difference(strings ...String) (difference String) {
	difference = s.Copy()
	for _, set := range strings {
		for strS, ok := range set {
			if ok {
				delete(difference, strS)
			}
		}
	}
	return difference
}

//FromStrings creates a set from strings
func FromStrings(strings ...string) String {
	s := make(String)
	for _, strS := range strings {
		s[strS] = true
	}
	return s
}

func (s String) Add(keys ...string) {
	for _, key := range keys {
		s[key] = true
	}
}

func (s String) Remove(keys ...string) {
	for _, key := range keys {
		delete(s, key)
	}
}

//FromSlices creates a slice of sets from slices of strings
func FromStringSlice(stringSlices ...[]string) []String {
	sets := make([]String, len(stringSlices))
	for i, slice := range stringSlices {
		sets[i] = FromStrings(slice...)
	}
	return sets
}

func (s String) ToSlice() []string {
	slice := make([]string, len(s))
	var i = 0
	for k := range s {
		slice[i] = k
		i++
	}
	return slice
}
