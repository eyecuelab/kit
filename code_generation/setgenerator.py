"""setgenerator.py generates a set type for a golang type
why does go not have a set type"""


BASE_STR = """
package _PACKAGE
type _TYPE_NAME map[_BASE_TYPE]bool

//Contains shows whether _SHORTHAND is in the _TYPE_NAME.
func (_RECIEVER _TYPE_NAME) Contains(_SHORTHAND _BASE_TYPE) bool {
	_, ok := _RECIEVER[_SHORTHAND]
	return ok
}

//Intersection returns the intersection of the _SLICE_NAME;
func (_RECIEVER _TYPE_NAME) Intersection(_SLICE_NAME ..._TYPE_NAME) (intersection _TYPE_NAME) {
	intersection = _RECIEVER
	for _, set := range _SLICE_NAME {
		for _SHORTHAND, ok := range set {
			if !ok {
				delete(intersection, _SHORTHAND)
			}
		}
	}
	return intersection
}

//Equal shows whether two _TYPENAMEs are equal; i.e, they contain the same items.
func (_RECIEVER _TYPE_NAME) Equal(other _TYPE_NAME) bool {
	for _SHORTHAND := range _RECIEVER {
		if !other.Contains(_SHORTHAND) {
			return false
		}
	}
	return true
}

//Union returns the union of the _SLICE_NAME.
func (_RECIEVER _TYPE_NAME) Union(_SLICE_NAME ..._TYPE_NAME) (union _TYPE_NAME) {
	union = _RECIEVER
	for _, set := range _SLICE_NAME {
		for _SHORTHAND := range set {
			union[_SHORTHAND] = true
		}
	}
	return union
}

//Difference returns the items in the reciever but not any other arguments
func(_RECIEVER _TYPE_NAME) Difference(_SLICE_NAME ..._TYPE_NAME) (difference _TYPE_NAME) {
	difference = _RECIEVER
	for _, set := range _SLICE_NAME {
		for _SHORTHAND, ok := range set {
			if ok {
				delete(difference, _SHORTHAND)
			}
		}
	}
	return difference
}

//From_BASE_TYPEs creates a set from _BASE_TYPEs
func From_BASE_TYPEs(_BASE_TYPEs ..._BASE_TYPE) _TYPE_NAME {
	set := make(_TYPE_NAME)
	for _, _SHORTHAND := range _BASE_TYPEs {
		set[_SHORTHAND] = true
	}
	return set
}
"""


def make_set_type(type_, base_type, reciever, shorthand, slice_, outfile=None, package=None):    
        code = BASE_STR.replace("_TYPE_NAME", type_)
        code = code.replace("_BASE_TYPE", base_type)
        code = code.replace("_RECIEVER", reciever)
        code = code.replace("_SLICE_NAME", slice_)
        code = code.replace("_SHORTHAND", shorthand)
        if package is None:
            code = code.replace("_PACKAGE", type_.lower())
        else:
            code = code.replace("_PACKAGE", package)
        if outfile is not None:
            with open(outfile, 'w+') as out:
                print(code, file=out)
        else:
            print(code)
