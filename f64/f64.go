package f64

//f64 is a library of helper functions for float64 math.

func Max(floats ...float64) (max float64, ok bool) {
	if len(floats) == 0 {
		return 0, false
	}
	max = floats[0]
	for _, x := range floats[1:] {
		if x > max {
			max = x
		}
	}
	return max, true
}

//Min returns the minimum
func Min(floats ...float64) (min float64, ok bool) {
	if len(floats) == 0 {
		return 0, false
	}
	min = floats[0]
	for _, x := range floats[1:] {
		if x < min {
			min = x
		}
	}
	return min, true
}

func Sum(floats ...float64) (sum float64) {
	for _, x := range floats {
		sum += x
	}
	return sum
}

func Product(floats ...float64) float64 {
	product := 1.0
	for _, x := range floats {
		product *= x
	}
	return product
}

func Map(f func(float64) float64, floats []float64) []float64 {
	output := make([]float64, len(floats))
	for i, x := range floats {
		output[i] = f(x)
	}
	return output
}

func ReduceInit(f func(float64, float64) float64, floats []float64, init float64) float64 {
	reduced := init
	for _, x := range floats {
		reduced = f(reduced, x)
	}
	return reduced
}

func Reduce(f func(float64, float64) float64, floats []float64) (reduced float64, ok bool) {
	if len(floats) == 0 {
		return 0.0, false
	}
	reduced = floats[0]
	for _, x := range floats {
		reduced = f(reduced, x)
	}
	return 0.0, true
}

//SortableFloat64 is a sortable float64 type
type SortableFloat64 []float64
