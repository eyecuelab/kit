package functools

//All returns false if any predicate is false. All() is True.
func All(predicates ...bool) bool {
	for _, predicate := range predicates {
		if !predicate {
			return false
		}
	}
	return true
}

//Any returns true if any predicate is true. Any() is False.
func Any(predicates ...bool) bool {
	for _, predicate := range predicates {
		if predicate {
			return true
		}
	}
	return false
}
