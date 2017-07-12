"""setgenerator.py generates a set type for a golang type
why does go not have a set type"""

import os

BASE_STR = """
package _PACKAGE
//_TYPE_NAME is a set type for _BASE_TYPE which supports the following set methods:
//intersection, union, equal, contains, difference, 
type _TYPE_NAME map[_BASE_TYPE]bool

//Contains shows whether _SHORTHAND is in the _TYPE_NAME.
func (_RECIEVER _TYPE_NAME) Contains(_SHORTHAND _BASE_TYPE) bool {
	_, ok := _RECIEVER[_SHORTHAND]
	return ok
}

//Intersection returns the intersection of the _TYPE_NAMEs;
func (_RECIEVER _TYPE_NAME) Intersection(_TYPE_NAMEs ..._TYPE_NAME) (intersection _TYPE_NAME) {
	intersection = _RECIEVER
	for _, set := range _TYPE_NAMEs {
		for _SHORTHAND, ok := range set {
			if !ok {
				delete(intersection, _SHORTHAND)
			}
		}
	}
	return intersection
}

//Equal shows whether two _TYPE_NAMEs are equal; i.e, they contain the same items.
func (_RECIEVER _TYPE_NAME) Equal(other _TYPE_NAME) bool {
	for _SHORTHAND := range _RECIEVER {
		if !other.Contains(_SHORTHAND) {
			return false
		}
	}
	return true
}

//Union returns the union of the _TYPE_NAMES.
func (_RECIEVER _TYPE_NAME) Union(_TYPE_NAMEs ..._TYPE_NAME) (union _TYPE_NAME) {
	union = _RECIEVER
	for _, set := range _TYPE_NAMEs {
		for _SHORTHAND := range set {
			union[_SHORTHAND] = true
		}
	}
	return union
}

//Difference returns the items in the reciever but not any other arguments
func(_RECIEVER _TYPE_NAME) Difference(_TYPE_NAMEs ..._TYPE_NAME) (difference _TYPE_NAME) {
	difference = _RECIEVER
	for _, set := range _TYPE_NAMEs {
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


def make_set_type(type_, base_type, reciever, shorthand, outfile=None, package=None):
    code = BASE_STR.replace("_TYPE_NAME", type_)
    code = code.replace("_BASE_TYPE", base_type)
    code = code.replace("_RECIEVER", reciever)
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


if __name__ == '__main__':
    types = {
        'int': ('Int', 'int', 'n',  'ints'),
        'int8': ('Int8', 'int8', 'n',  'int8s'),
        'int16': ('Int16', 'int16', 'n',  'int16s'),
        'int32': ('Int32', 'int32', 'n',  'int32s'),
        'int64': ('Int64', 'int64', 'n',  'int64s'),
        'uint8': ('Uint8', 'uint8', 'u',  'uint8s'),
        'uint16': ('Uint16', 'uint16', 'u',  'uint16s'),
        'uint32': ('Uint32', 'uint32', 'u',  'uint32s'),
        'uint64': ('Uint64', 'uint64', 'u',  'uint64s'),
        'float32': ('Float32', 'float32', 'x',  'float32s'),
        'float64': ('Float64', 'float64', 'x',  'float64s'),
        'string': ('String', 'string', 's',  'strings')}
    for type_, args in types.items():
        outfile = os.curdir + type_ + '.go'
        make_set_type(*args, package='set',
                      outfile=f'C:/users/efron/desktop/{type_}.go')
