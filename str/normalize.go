package str

import (
	"strings"
	"unicode"

	"golang.org/x/text/runes"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

//NormEqual returns true if the NKFC normalized forms of both strings are equal.
func NormEqual(s, q string) bool {
	return NKFC(s) == NKFC(q)
}

//NormFoldEqual returns true if the casefolded, NKFC normalized forms of both strings are equal.
func NormFoldEqual(s, q string) bool {
	return strings.EqualFold(NKFC(s), NKFC(q))
}

//NKFC normalizes a string to it's NKFC form
func NKFC(s string) string {
	return norm.NFKC.String(s)
}

//NFKD normalizes a string to it's NFKD form
func NFKD(s string) string {
	return norm.NFKD.String(s)
}

//NFD normalizes a string to it's NFD form
func NFD(s string) string {
	return norm.NFD.String(s)
}

//NFC normalizes a string to it's NFC form
func NFC(s string) string {
	return norm.NFC.String(s)
}

func isNonSpacingMark(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}

var (
	removeNonSpacingMarks = runes.Remove(runes.In(unicode.Mn))
	removeWhiteSpace      = runes.Remove(runes.In(unicode.White_Space))
	removePunctuation     = runes.Remove(runes.In(unicode.Punct))
	removeControl         = runes.Remove(runes.In(unicode.C))
)

//RemoveDiacriticsNFC creates a copy of s with the diacritics removed. It also transforms it to NFC.
//Thread Safe
func RemoveDiacriticsNFC(s string) string {
	var diacriticRemover = transform.Chain(norm.NFD, removeNonSpacingMarks, norm.NFC)
	out, _, _ := transform.String(diacriticRemover, s)
	return out
}

//RemovePunctuationNTS removes punctuation (as defined by unicode) from a string, but is NOT thread safe.
//Note that this converts to runes and back to UTF-8, so RemoveWhiteSpace(s) == s
//for a string that contains non-punctuation characters does not necessarially hold, since the code points may differ.
func RemovePunctuationNTS(s string) string {
	out, _, _ := transform.String(removePunctuation, s)
	return out
}

//RemoveWhiteSpaceNTS removes whitespace (as defined by unicode) from a string, but is NOT thread safe.
//Note that this converts to runes and back to UTF-8, so RemoveWhiteSpace(s) == s
//for a non-whitespace string does not necessarially hold, since the code points may differ.
//Note that this is faster than RemoveWhitespace, but is not thread safe.
func RemoveWhiteSpaceNTS(s string) string {
	out, _, _ := transform.String(removePunctuation, s)
	return out
}

//RemoveControlNTS removes whitespace (as defined by unicode) from a string, but is NOT thread safe.
//Note that this converts to runes and back to UTF-8, so RemoveWhiteSpace(s) == s
//for a non-whitespace string does not necessarially hold, since the code points may differ.
//Note that this is faster than RemoveWhitespace, but is not thread safe.
func RemoveControlNTS(s string) string {
	out, _, _ := transform.String(removeControl, s)
	return out
}

//ExtremeNormalization heavily normalizes a stirng for purposes of comparison and safety.
//It removes ALL nonspacing marks, whitespace, punctuation, control characters, and transforms the
//string to NFKC encoding. This can and will lose a lot of information!
func ExtremeNormalization(s string) string {
	s = strings.ToLower(s)
	normalizer := transform.Chain(
		norm.NFKD,
		removeNonSpacingMarks,
		removeWhiteSpace,
		removePunctuation,
		removeControl,
		norm.NFKC)
	out, _, _ := transform.String(normalizer, s)
	return out
}
