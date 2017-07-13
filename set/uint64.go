
package set
type Uint64 map[uint64]bool

//Contains shows whether uis is in the Uint64.
func (u Uint64) Contains(uis uint64) bool {
	_, ok := u[uis]
	return ok
}

//Intersection returns the intersection of the uint64s;
func (u Uint64) Intersection(uint64s ...Uint64) (intersection Uint64) {
	intersection = u
	for _, set := range uint64s {
		for uis, ok := range set {
			if !ok {
				delete(intersection, uis)
			}
		}
	}
	return intersection
}

//Equal shows whether two Uint64s are equal; i.e, they contain the same items.
func (u Uint64) Equal(other Uint64) bool {
	for uis := range u {
		if !other.Contains(uis) {
			return false
		}
	}
	return true
}

//Union returns the union of the uint64s.
func (u Uint64) Union(uint64s ...Uint64) (union Uint64) {
	union = u
	for _, set := range uint64s {
		for uis := range set {
			union[uis] = true
		}
	}
	return union
}

//Difference returns the items in the reciever but not any other arguments
func(u Uint64) Difference(uint64s ...Uint64) (difference Uint64) {
	difference = u
	for _, set := range uint64s {
		for uis, ok := range set {
			if ok {
				delete(difference, uis)
			}
		}
	}
	return difference
}

//Fromuint64s creates a set from uint64s
func Fromuint64s(uint64s ...uint64) Uint64 {
	set := make(Uint64)
	for _, uis := range uint64s {
		set[uis] = true
	}
	return set
}

