
package set
type Int32 map[int32]bool

//Contains shows whether is is in the Int32.
func (n Int32) Contains(is int32) bool {
	_, ok := n[is]
	return ok
}

//Intersection returns the intersection of the int32s;
func (n Int32) Intersection(int32s ...Int32) (intersection Int32) {
	intersection = n
	for _, set := range int32s {
		for is, ok := range set {
			if !ok {
				delete(intersection, is)
			}
		}
	}
	return intersection
}

//Equal shows whether two Int32s are equal; i.e, they contain the same items.
func (n Int32) Equal(other Int32) bool {
	for is := range n {
		if !other.Contains(is) {
			return false
		}
	}
	return true
}

//Union returns the union of the int32s.
func (n Int32) Union(int32s ...Int32) (union Int32) {
	union = n
	for _, set := range int32s {
		for is := range set {
			union[is] = true
		}
	}
	return union
}

//Difference returns the items in the reciever but not any other arguments
func(n Int32) Difference(int32s ...Int32) (difference Int32) {
	difference = n
	for _, set := range int32s {
		for is, ok := range set {
			if ok {
				delete(difference, is)
			}
		}
	}
	return difference
}

//Fromint32s creates a set from int32s
func Fromint32s(int32s ...int32) Int32 {
	set := make(Int32)
	for _, is := range int32s {
		set[is] = true
	}
	return set
}

