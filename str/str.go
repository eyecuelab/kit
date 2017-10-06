package str

import (
	"bytes"
	"fmt"

	"github.com/eyecuelab/kit/runeset"
)

//Runes removes any runes listed from the string.
//Note that this converts to runes and back to UTF-8, so RemoveRunes(s) == s
//does not necessarially hold, since the code points may differ.
func RemoveRunes(s string, toRemove ...rune) string {
	buf := bytes.Buffer{}
	set := runeset.FromRunes(toRemove...)
	var r rune
	for _, r = range s {
		if !set.Contains(r) {
			buf.WriteRune(r)
		}
	}
	return buf.String()
}

//SRemoveRunes removes any of the runes contained in the string toRemove from s.
//Equivalent to RemoveRunes(s, []rune(toRemove))
func SRemoveRunes(s string, toRemove string) string {
	buf := bytes.Buffer{}
	set := runeset.FromString(s)
	for _, r := range s {
		if !set.Contains(r) {
			buf.WriteRune(r)
		}
	}
	return buf.String()
}

//MapErr maps f(string)string, err across a slice. If any error results, it
// it immediately returns an empty slice and the error.
func MapErr(f func(string) (string, error), strings []string) ([]string, error) {
	for i, str := range strings {
		result, err := f(str)
		if err != nil {
			return []string{}, fmt.Errorf("MapErr: %v", err)
		}
		strings[i] = result
	}
	return strings, nil
}

//Map maps f(string) across the remaining arguments, returning [f(str) for str in str]
func Map(f func(string) string, strings []string) []string {
	for i, str := range strings {
		strings[i] = f(str)
	}
	return strings
}
