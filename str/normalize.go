package str

import (
	"bytes"
	"strings"

	"github.com/eyecuelab/kit/runeset"
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

func removeRunesNotInSet(s string, set runeset.RuneSet) string {
	var buf bytes.Buffer
	for _, r := range s {
		if set.Contains(r) {
			buf.WriteRune(r)
		}
	}
	return buf.String()
}
func removeRunesInSet(s string, set runeset.RuneSet) string {
	var buf bytes.Buffer
	for _, r := range s {
		if !set.Contains(r) {
			buf.WriteRune(r)
		}
	}
	return buf.String()
}
