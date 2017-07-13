
package set
type Uint32 map[uint32]bool

//Contains shows whether uis is in the Uint32.
func (u Uint32) Contains(uis uint32) bool {
	_, ok := u[uis]
	return ok
}

//Intersection returns the intersection of the uint32s;
func (u Uint32) Intersection(uint32s ...Uint32) (intersection Uint32) {
	intersection = u
	for _, set := range uint32s {
		for uis, ok := range set {
			if !ok {
				delete(intersection, uis)
			}
		}
	}
	return intersection
}

//Equal shows whether two Uint32s are equal; i.e, they contain the same items.
func (u Uint32) Equal(other Uint32) bool {
	for uis := range u {
		if !other.Contains(uis) {
			return false
		}
	}
	return true
}

//Union returns the union of the uint32s.
func (u Uint32) Union(uint32s ...Uint32) (union Uint32) {
	union = u
	for _, set := range uint32s {
		for uis := range set {
			union[uis] = true
		}
	}
	return union
}

//Difference returns the items in the reciever but not any other arguments
func(u Uint32) Difference(uint32s ...Uint32) (difference Uint32) {
	difference = u
	for _, set := range uint32s {
		for uis, ok := range set {
			if ok {
				delete(difference, uis)
			}
		}
	}
	return difference
}

//Fromuint32s creates a set from uint32s
func Fromuint32s(uint32s ...uint32) Uint32 {
	set := make(Uint32)
	for _, uis := range uint32s {
		set[uis] = true
	}
	return set
}

