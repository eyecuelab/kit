
package set
type Float64 map[float64]bool

//Contains shows whether fs is in the Float64.
func (x Float64) Contains(fs float64) bool {
	_, ok := x[fs]
	return ok
}

//Intersection returns the intersection of the float64s;
func (x Float64) Intersection(float64s ...Float64) (intersection Float64) {
	intersection = x
	for _, set := range float64s {
		for fs, ok := range set {
			if !ok {
				delete(intersection, fs)
			}
		}
	}
	return intersection
}

//Equal shows whether two Float64s are equal; i.e, they contain the same items.
func (x Float64) Equal(other Float64) bool {
	for fs := range x {
		if !other.Contains(fs) {
			return false
		}
	}
	return true
}

//Union returns the union of the float64s.
func (x Float64) Union(float64s ...Float64) (union Float64) {
	union = x
	for _, set := range float64s {
		for fs := range set {
			union[fs] = true
		}
	}
	return union
}

//Difference returns the items in the reciever but not any other arguments
func(x Float64) Difference(float64s ...Float64) (difference Float64) {
	difference = x
	for _, set := range float64s {
		for fs, ok := range set {
			if ok {
				delete(difference, fs)
			}
		}
	}
	return difference
}

//Fromfloat64s creates a set from float64s
func Fromfloat64s(float64s ...float64) Float64 {
	set := make(Float64)
	for _, fs := range float64s {
		set[fs] = true
	}
	return set
}

