"""
mathhelpers.py generates helper functions for a mathematical type.
"""

BASE_STR = """

//Max checks to see if a slice is non-empty and returns the maximum elements and true if so.
//Returns 0, false if slice is empty.
func Max(_SLICE_NAME ..._TYPE) (max _TYPE, ok bool) {
	if len(_SLICE_NAME) == 0 {
		return 0, false
	}
	max = _SLICE_NAME[0]
	for _, x := range _SLICE_NAME[1:] {
		if x > max {
			max = x
		}
	}
	return max, true
}

//Min checks to see if a slice is non-empty and returns the minimum element and true if so.
//Returns 0, false if slice is empty. Equivalent to reduce(math.min, _SLICE_NAME)
})
func Min(_SLICE_NAME ..._TYPE) (min _TYPE, ok bool) {
	if len(_SLICE_NAME) == 0 {
		return 0, false
	}
	min = _SLICE_NAME[0]
	for _, x := range _SLICE_NAME[1:] {
		if x < min {
			min = x
		}
	}
	return min, true
}


//Sum outputs the sum of the given slice. Sum([]) == 0.
//Equivalent to ReduceInit(func(a, b _TYPE)_TYPE {return a+b}, _SLICE_NAME, 0)
func Sum(_SLICE_NAME ..._TYPE) (sum _TYPE) {
	for _, x := range _SLICE_NAME {
		sum += x
	}
	return sum
}

//Product outputs the product of the given slice. Product([]) == 1.
//Equivalent to ReduceInit(func(a, b _TYPE)_TYPE {return a*b}, _SLICE_NAME, 1)
func Product(_SLICE_NAME ..._TYPE) _TYPE {
	product := 1.0
	for _, x := range _SLICE_NAME {
		product *= x
	}
	return product
}

func Map(f func(_TYPE) _TYPE, _SLICE_NAME []_TYPE) []_TYPE {
	output := make([]_TYPE, len(_SLICE_NAME))
	for i, x := range _SLICE_NAME {
		output[i] = f(x)
	}
	return output
}

//ReduceInit starts with init and applies f(init, x) for each x in _SLICE_NAME.
func ReduceInit(f func(_TYPE, _TYPE) _TYPE, _SLICE_NAME []_TYPE, init _TYPE) _TYPE {
	reduced := init
	for _, x := range _SLICE_NAME {
		reduced = f(reduced, x)
	}
	return reduced
}

//Reduce checks to see if a slice is non-empty. If so, it reduces with f across the slice,
//and returns reduced, true.
//Otherwise, it returns 0.0, false.
func Reduce(f func(_TYPE, _TYPE) _TYPE, _SLICE_NAME []_TYPE) (reduced _TYPE, ok bool) {
	if len(_SLICE_NAME) == 0 {
		return 0.0, false
	}
	reduced = _SLICE_NAME[0]
	for _, x := range _SLICE_NAME {
		reduced = f(reduced, x)
	}
	return 0.0, true
}

//Sortable_TYPE is a sortable _TYPE type
type Sortable_TYPE []_TYPE
"""


def make_math_helpers(type_, slice_, outfile=None, package=None):    
        code = BASE_STR.replace("_TYPE", type_)
        code = code.replace("_SLICE_NAME", slice_)
        if package is None:
            code = code.replace("_PACKAGE", type_.lower())
        else:
            code = code.replace("_PACKAGE", package)
        if outfile is not None:
            with open(outfile, 'w+') as out:
                print(code, file=out)
        else:
            print(code)