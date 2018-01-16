package functools

//All returns false if any predicate is false. All() is True.
func All(bools ...bool) bool {
	for _, b := range bools {
		if !b {
			return false
		}
	}
	return true
}

//Any returns true if any predicate is true. Any() is False.
func Any(bools ...bool) bool {
	for _, b := range bools {
		if b {
			return true
		}
	}
	return false
}

//None returns true if all predicates are false. None() is True.
func None(bools ...bool) bool {
	for _, b := range bools {
		if b {
			return false
		}
	}
	return true
}
