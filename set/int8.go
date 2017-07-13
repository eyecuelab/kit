package set

type Int8 map[int8]bool

//Contains shows whether is is in the Int8.
func (n Int8) Contains(is int8) bool {
	_, ok := n[is]
	return ok
}

//Intersection returns the intersection of the set.Int8s;
func (n Int8) Intersection(int8s ...Int8) (intersection Int8) {
	intersection = n
	for _, set := range int8s {
		for is, ok := range set {
			if !ok {
				delete(intersection, is)
			}
		}
	}
	return intersection
}

//Equal shows whether two set.Int8s are equal; i.e, they contain the same items.
func (n Int8) Equal(other Int8) bool {
	for is := range n {
		if !other.Contains(is) {
			return false
		}
	}
	return true
}

//Union returns the union of the int8s.
func (n Int8) Union(int8s ...Int8) (union Int8) {
	union = n
	for _, set := range int8s {
		for is := range set {
			union[is] = true
		}
	}
	return union
}

//Difference returns the items in the reciever but not any other arguments
func (n Int8) Difference(int8s ...Int8) (difference Int8) {
	difference = n
	for _, set := range int8s {
		for is, ok := range set {
			if ok {
				delete(difference, is)
			}
		}
	}
	return difference
}

//Fromint8s creates a set from int8s
func Fromint8s(int8s ...int8) Int8 {
	set := make(Int8)
	for _, is := range int8s {
		set[is] = true
	}
	return set
}
