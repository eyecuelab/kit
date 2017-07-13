
package set
type String map[string]bool

//Contains shows whether strS is in the String.
func (s String) Contains(strS string) bool {
	_, ok := s[strS]
	return ok
}

//Intersection returns the intersection of the strings;
func (s String) Intersection(strings ...String) (intersection String) {
	intersection = s
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
	union = s
	for _, set := range strings {
		for strS := range set {
			union[strS] = true
		}
	}
	return union
}

//Difference returns the items in the reciever but not any other arguments
func(s String) Difference(strings ...String) (difference String) {
	difference = s
	for _, set := range strings {
		for strS, ok := range set {
			if ok {
				delete(difference, strS)
			}
		}
	}
	return difference
}

//Fromstrings creates a set from strings
func Fromstrings(strings ...string) String {
	set := make(String)
	for _, strS := range strings {
		set[strS] = true
	}
	return set
}

