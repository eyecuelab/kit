package runeset

//RuneSet is a map[rune]bool with a contains method
type RuneSet map[rune]inter

//Contains shows whether r is in the RuneSet.
func (rs RuneSet) Contains(r rune) bool {
	_, ok := rs[r]
	return ok
}

//Intersection returns the intersection of the sets;
func (rs RuneSet) Intersection(sets ...RuneSet) (intersection RuneSet) {
	intersection = rs
	for _, set := range sets {
		for r, ok := range set {
			if !ok {
				delete(intersection, r)
			}
		}
	}
	return intersection
}

//Equal shows whether two RuneSets are equal; i.e, they contain the same items.
func (rs RuneSet) Equal(other RuneSet) bool {
	for r := range rs {
		if !other.Contains(r) {
			return false
		}
	}
	return true
}

//Union returns the union of the sets.
func (rs RuneSet) Union(sets ...RuneSet) (union RuneSet) {
	union = rs
	for _, set := range sets {
		for r := range set {
			union[r] = true
		}
	}
	return union
}

//Difference returns the items in the reciever but not any other arguments
//i.e, if set = {'a', b' 'c'}; set.Difference({'b', 'c'}) = {'a'}
func (rs RuneSet) Difference(sets ...RuneSet) (difference RuneSet) {
	difference = rs
	for _, set := range sets {
		for r, ok := range set {
			if ok {
				delete(difference, r)
			}
		}
	}
	return difference
}

//FromRunes creates a set from runes
func FromRunes(runes ...rune) RuneSet {
	set := make(RuneSet)
	for _, r := range runes {
		set[r] = true
	}
	return set
}

//FromString converts a string to a RuneSet of the runes inside.
func FromString(s string) (set RuneSet) {
	var r rune
	for _, r = range s {
		set[r] = true
	}
	return set
}
func all()
