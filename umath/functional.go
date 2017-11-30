package umath

func Map(f func(uint) uint, a []uint) []uint {
	mapped := make([]uint, len(a))
	for i, n := range a {
		mapped[i] = f(n)
	}
	return mapped
}

func Filter(f func(uint) bool, a []uint) []uint {
	filtered := make([]uint, 0, len(a))
	for _, n := range a {
		if f(n) {
			filtered = append(filtered, n)
		}
	}
	return filtered
}

func Reduce(f func(uint, uint) uint, start uint, a ...uint) uint {
	reduced := start
	for _, n := range a {
		reduced = f(reduced, n)
	}
	return reduced
}

func Accumulate(f func(uint, uint) uint, a []uint) []uint {
	accumulated := make([]uint, len(a))
	if len(a) == 0 {
		return accumulated
	}
	accumulated[0] = a[0]
	for i, n := range a[1:] {
		accumulated[i+1] = f(accumulated[i], n)
	}
	return accumulated
}
