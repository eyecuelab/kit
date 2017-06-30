package str

import "github.com/eyecuelab/kit/runes"

//RemoveWhiteSpace removes whitespace {'\n' '\t' ' ' '\r`} from a string.
//Note that this converts to runes and back to UTF-8, so RemoveWhiteSpace(s) == s
//for a non-whitespace string does not necessarially hold, since the code points may differ.
func RemoveWhiteSpace(s string) string {
	var bytes []rune
	for _, char := range s {
		switch char {
		case '\n', '\t', ' ', '\r':
			// pass
		default:
			bytes = append(bytes, char)
		}

	}
	return string(bytes)
}

//RemoveRunes removes any runes listed from the string.
//Note that this converts to runes and back to UTF-8, so RemoveRunes(s) == s
//does not necessarially hold, since the code points may differ.
func RemoveRunes(s string, toRemove ...rune) string {
	var bytes []rune
	set := runes.Set(toRemove...)
	for _, r := range s {
		_, unwanted := set[r]
		if !unwanted {
			bytes = append(bytes, r)
		}
	}
	return string(bytes)

}

//ToRuneSet converts a string to a set of the runes inside. i.e, a map of bools
func ToRuneSet(s string) map[rune]bool {
	set := make(map[rune]bool)
	for _, r := range s {
		set[r] = true
	}
	return set
}
