package str

import (
	"strings"

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
