package sortlib

/*
import (
	"fmt"
	"sort"
)

type numericSorter struct {
	numerals []interface{}
	floats   []float64
}


//Numerals returns a sorted copy of a slice if the slice contains numeric types only.
//Otherwise, it returns the slice itself and false.
func Numerals(a []interface{}) ([]interface{}, bool) {
	noOp := func(x float64) float64 { return x }
	sortable, ok := toNumericSorter(a, noOp)
	if !ok {
		return a, false
	}
	sort.Sort(sortable)
	return sortable.numerals, true
}

//NumeralsKey returns a sorted copy of a slice according to the key function, if slice contains numeric types only.
//Otherwise, it returns the slice itself and false.
func NumeralsKey(a []interface{}, f func(float64) float64) ([]interface{}, bool) {
	sortable, ok := toNumericSorter(a, f)
	if !ok {
		return a, false
	}
	sort.Sort(sortable)
	return sortable.numerals, true
}

func toNumericSorter(a []interface{}, f func(float64) float64) (numericSorter, bool) {
	floats := make([]float64, len(a))
	for i, n := range a {
		if x, ok := coerceToFloat(n); ok {
			floats[i] = f(x)
		} else {
			return numericSorter{}, false
		}
	}
	fmt.Println(a, floats)
	return numericSorter{numerals: a, floats: floats}, true
}

func coerceToFloat(a interface{}) (float64, bool) {
	switch a := a.(type) {
	default:
		return 0, false
	case int8:
		return float64(a), true
	case int16:
		return float64(a), true
	case int32:
		return float64(a), true
	case int64:
		return float64(a), true
	case uint8:
		return float64(a), true
	case uint16:
		return float64(a), true
	case uint32:
		return float64(a), true
	case uint64:
		return float64(a), true
	case float32:
		return float64(a), true
	case float64:
		return float64(a), true
	}

}

func (n numericSorter) Less(i, j int) bool {
	return n.floats[i] < n.floats[j]
}

func (n numericSorter) Len() int {
	return len(n.numerals)
}

func (n numericSorter) Swap(i, j int) {
	n.floats[i], n.floats[j], n.numerals[i], n.numerals[j] = n.floats[j], n.floats[i], n.numerals[j], n.numerals[i]
}
*/
