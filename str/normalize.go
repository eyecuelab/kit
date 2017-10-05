package str

import (
	"strings"
	"unicode"

	"golang.org/x/text/runes"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

//RemoveWhiteSpace removes whitespace {'\n' '\t' ' ' '\r`} from a string.
//Note that this converts to runes and back to UTF-8, so RemoveWhiteSpace(s) == s
//for a non-whitespace string does not necessarially hold, since the code points may differ.
func RemoveWhiteSpace(s string) string {
	return RemoveRunes(s, '\n', ' ', '\r', '\t')
}

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

var removeNonSpacingMarks = runes.Remove(runes.In(unicode.Mn))
var diacriticRemover = transform.Chain(norm.NFD, removeNonSpacingMarks, norm.NFC)

//RemoveDiacriticsNFC creates a copy of s with the diacritics removed. It also transforms it to NFC.
//Thread Safe
func RemoveDiacriticsNFC(s string) string {
	var diacriticRemover = transform.Chain(norm.NFD, removeNonSpacingMarks, norm.NFC)
	out, _, _ := transform.String(diacriticRemover, s)
	return out
}
