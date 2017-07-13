
package set
type Uint8 map[uint8]bool

//Contains shows whether uis is in the Uint8.
func (u Uint8) Contains(uis uint8) bool {
	_, ok := u[uis]
	return ok
}

//Intersection returns the intersection of the uint8s;
func (u Uint8) Intersection(uint8s ...Uint8) (intersection Uint8) {
	intersection = u
	for _, set := range uint8s {
		for uis, ok := range set {
			if !ok {
				delete(intersection, uis)
			}
		}
	}
	return intersection
}

//Equal shows whether two Uint8s are equal; i.e, they contain the same items.
func (u Uint8) Equal(other Uint8) bool {
	for uis := range u {
		if !other.Contains(uis) {
			return false
		}
	}
	return true
}

//Union returns the union of the uint8s.
func (u Uint8) Union(uint8s ...Uint8) (union Uint8) {
	union = u
	for _, set := range uint8s {
		for uis := range set {
			union[uis] = true
		}
	}
	return union
}

//Difference returns the items in the reciever but not any other arguments
func(u Uint8) Difference(uint8s ...Uint8) (difference Uint8) {
	difference = u
	for _, set := range uint8s {
		for uis, ok := range set {
			if ok {
				delete(difference, uis)
			}
		}
	}
	return difference
}

//Fromuint8s creates a set from uint8s
func Fromuint8s(uint8s ...uint8) Uint8 {
	set := make(Uint8)
	for _, uis := range uint8s {
		set[uis] = true
	}
	return set
}

