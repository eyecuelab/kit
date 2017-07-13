
package set
type Int64 map[int64]bool

//Contains shows whether is is in the Int64.
func (n Int64) Contains(is int64) bool {
	_, ok := n[is]
	return ok
}

//Intersection returns the intersection of the int64s;
func (n Int64) Intersection(int64s ...Int64) (intersection Int64) {
	intersection = n
	for _, set := range int64s {
		for is, ok := range set {
			if !ok {
				delete(intersection, is)
			}
		}
	}
	return intersection
}

//Equal shows whether two Int64s are equal; i.e, they contain the same items.
func (n Int64) Equal(other Int64) bool {
	for is := range n {
		if !other.Contains(is) {
			return false
		}
	}
	return true
}

//Union returns the union of the int64s.
func (n Int64) Union(int64s ...Int64) (union Int64) {
	union = n
	for _, set := range int64s {
		for is := range set {
			union[is] = true
		}
	}
	return union
}

//Difference returns the items in the reciever but not any other arguments
func(n Int64) Difference(int64s ...Int64) (difference Int64) {
	difference = n
	for _, set := range int64s {
		for is, ok := range set {
			if ok {
				delete(difference, is)
			}
		}
	}
	return difference
}

//Fromint64s creates a set from int64s
func Fromint64s(int64s ...int64) Int64 {
	set := make(Int64)
	for _, is := range int64s {
		set[is] = true
	}
	return set
}

