
package set
type Int16 map[int16]bool

//Contains shows whether is is in the Int16.
func (n Int16) Contains(is int16) bool {
	_, ok := n[is]
	return ok
}

//Intersection returns the intersection of the int16s;
func (n Int16) Intersection(int16s ...Int16) (intersection Int16) {
	intersection = n
	for _, set := range int16s {
		for is, ok := range set {
			if !ok {
				delete(intersection, is)
			}
		}
	}
	return intersection
}

//Equal shows whether two Int16s are equal; i.e, they contain the same items.
func (n Int16) Equal(other Int16) bool {
	for is := range n {
		if !other.Contains(is) {
			return false
		}
	}
	return true
}

//Union returns the union of the int16s.
func (n Int16) Union(int16s ...Int16) (union Int16) {
	union = n
	for _, set := range int16s {
		for is := range set {
			union[is] = true
		}
	}
	return union
}

//Difference returns the items in the reciever but not any other arguments
func(n Int16) Difference(int16s ...Int16) (difference Int16) {
	difference = n
	for _, set := range int16s {
		for is, ok := range set {
			if ok {
				delete(difference, is)
			}
		}
	}
	return difference
}

//Fromint16s creates a set from int16s
func Fromint16s(int16s ...int16) Int16 {
	set := make(Int16)
	for _, is := range int16s {
		set[is] = true
	}
	return set
}

