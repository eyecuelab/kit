
package set
type Int map[int]bool

//Contains shows whether is is in the Int.
func (n Int) Contains(is int) bool {
	_, ok := n[is]
	return ok
}

//Intersection returns the intersection of the ints;
func (n Int) Intersection(ints ...Int) (intersection Int) {
	intersection = n
	for _, set := range ints {
		for is, ok := range set {
			if !ok {
				delete(intersection, is)
			}
		}
	}
	return intersection
}

//Equal shows whether two Ints are equal; i.e, they contain the same items.
func (n Int) Equal(other Int) bool {
	for is := range n {
		if !other.Contains(is) {
			return false
		}
	}
	return true
}

//Union returns the union of the ints.
func (n Int) Union(ints ...Int) (union Int) {
	union = n
	for _, set := range ints {
		for is := range set {
			union[is] = true
		}
	}
	return union
}

//Difference returns the items in the reciever but not any other arguments
func(n Int) Difference(ints ...Int) (difference Int) {
	difference = n
	for _, set := range ints {
		for is, ok := range set {
			if ok {
				delete(difference, is)
			}
		}
	}
	return difference
}

//Fromints creates a set from ints
func Fromints(ints ...int) Int {
	set := make(Int)
	for _, is := range ints {
		set[is] = true
	}
	return set
}

