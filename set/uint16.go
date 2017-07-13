
package set
type Uint16 map[uint16]bool

//Contains shows whether uis is in the Uint16.
func (u Uint16) Contains(uis uint16) bool {
	_, ok := u[uis]
	return ok
}

//Intersection returns the intersection of the uint16s;
func (u Uint16) Intersection(uint16s ...Uint16) (intersection Uint16) {
	intersection = u
	for _, set := range uint16s {
		for uis, ok := range set {
			if !ok {
				delete(intersection, uis)
			}
		}
	}
	return intersection
}

//Equal shows whether two Uint16s are equal; i.e, they contain the same items.
func (u Uint16) Equal(other Uint16) bool {
	for uis := range u {
		if !other.Contains(uis) {
			return false
		}
	}
	return true
}

//Union returns the union of the uint16s.
func (u Uint16) Union(uint16s ...Uint16) (union Uint16) {
	union = u
	for _, set := range uint16s {
		for uis := range set {
			union[uis] = true
		}
	}
	return union
}

//Difference returns the items in the reciever but not any other arguments
func(u Uint16) Difference(uint16s ...Uint16) (difference Uint16) {
	difference = u
	for _, set := range uint16s {
		for uis, ok := range set {
			if ok {
				delete(difference, uis)
			}
		}
	}
	return difference
}

//Fromuint16s creates a set from uint16s
func Fromuint16s(uint16s ...uint16) Uint16 {
	set := make(Uint16)
	for _, uis := range uint16s {
		set[uis] = true
	}
	return set
}

