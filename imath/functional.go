package imath

//Map (f, a) returns a slice containing f(n) for every n in a
func Map(f func(int) int, a []int) []int {
	mapped := make([]int, len(a))
	for i, n := range a {
		mapped[i] = f(n)
	}
	return mapped
}

//Filter (f, a) returns a slice containing n for all n where f(n) == true
func Filter(f func(int) bool, a []int) []int {
	filtered := make([]int, 0, len(a))
	for _, n := range a {
		if f(n) {
			filtered = append(filtered, n)
		}
	}
	return filtered
}

func Reduce(f func(int, int) int, start int, a ...int) int {
	reduced := start
	for _, n := range a {
		reduced = f(reduced, n)
	}
	return reduced
}

//Accumulate returns a slice B where B[i] = Reduce(f, a[:i+1]).
func Accumulate(f func(int, int) int, a []int) []int {
	accumulated := make([]int, len(a))
	if len(a) == 0 {
		return accumulated
	}
	accumulated[0] = a[0]
	for i, n := range a[1:] {
		accumulated[i+1] = f(accumulated[i], n)
	}
	return accumulated
}
