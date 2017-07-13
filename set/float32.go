
package set
type Float32 map[float32]bool

//Contains shows whether fs is in the Float32.
func (x Float32) Contains(fs float32) bool {
	_, ok := x[fs]
	return ok
}

//Intersection returns the intersection of the float32s;
func (x Float32) Intersection(float32s ...Float32) (intersection Float32) {
	intersection = x
	for _, set := range float32s {
		for fs, ok := range set {
			if !ok {
				delete(intersection, fs)
			}
		}
	}
	return intersection
}

//Equal shows whether two Float32s are equal; i.e, they contain the same items.
func (x Float32) Equal(other Float32) bool {
	for fs := range x {
		if !other.Contains(fs) {
			return false
		}
	}
	return true
}

//Union returns the union of the float32s.
func (x Float32) Union(float32s ...Float32) (union Float32) {
	union = x
	for _, set := range float32s {
		for fs := range set {
			union[fs] = true
		}
	}
	return union
}

//Difference returns the items in the reciever but not any other arguments
func(x Float32) Difference(float32s ...Float32) (difference Float32) {
	difference = x
	for _, set := range float32s {
		for fs, ok := range set {
			if ok {
				delete(difference, fs)
			}
		}
	}
	return difference
}

//Fromfloat32s creates a set from float32s
func Fromfloat32s(float32s ...float32) Float32 {
	set := make(Float32)
	for _, fs := range float32s {
		set[fs] = true
	}
	return set
}

