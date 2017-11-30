package runeset

//RuneSet is a map[rune]bool with the methods you would expect from a set type.
//Eg, Contains, Union, Intersection, and Difference.
//I will make code generation for further set types in the future.
type Signal struct{}

var yes Signal

type RuneSet map[rune]Signal

//Contains shows whether r is in the RuneSet.
func (rs RuneSet) Contains(r rune) bool {
	_, ok := rs[r]
	return ok
}

//Intersection returns the intersection of the sets;
func (rs RuneSet) Intersection(sets ...RuneSet) (intersection RuneSet) {
	intersection = rs.Copy()
	for _, set := range sets {
		for key := range intersection {
			if _, ok := set[key]; !ok {
				delete(intersection, key)
			}
		}
	}
	return intersection
}

func Intersection(set RuneSet, sets ...RuneSet) RuneSet {
	return set.Intersection(sets...)
}

//Equal shows whether two RuneSets are equal; i.e, they contain the same items.
func (rs RuneSet) Equal(other RuneSet) bool {
	if len(rs) != len(other) {
		return false
	}
	for r := range rs {
		if !other.Contains(r) {
			return false
		}
	}
	return true
}

//Union returns the union of the sets.
func (rs RuneSet) Union(sets ...RuneSet) (union RuneSet) {
	sets = append(sets, rs)
	return Union(sets...)
}

func Union(sets ...RuneSet) RuneSet {
	union := make(RuneSet)
	for _, set := range sets {
		for r := range set {
			union[r] = yes
		}
	}
	return union
}

//Difference returns the items in the reciever but not any other arguments
//i.e, if set = {'a', b' 'c'}; set.Difference({'b', 'c'}) = {'a'}
func (rs RuneSet) Difference(sets ...RuneSet) (difference RuneSet) {
	difference = rs.Copy()
	for _, set := range sets {
		for key := range difference {
			if _, ok := set[key]; ok {
				delete(difference, key)
			}
		}
	}
	return difference
}

//FromRunes creates a set from runes
func FromRunes(runes ...rune) RuneSet {
	set := make(RuneSet)
	for _, r := range runes {
		set[r] = yes
	}
	return set
}

//FromString converts a string to a RuneSet of the runes inside.
func FromString(s string) (set RuneSet) {
	set = make(RuneSet)
	for _, r := range s {
		set[r] = yes
	}
	return set
}

//Copy returns a copy of the RuneSet.
func (rs RuneSet) Copy() RuneSet {
	copy := make(RuneSet)
	for k, v := range rs {
		copy[k] = v
	}
	return copy
}
