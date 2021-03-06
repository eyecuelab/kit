package str

import (
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"golang.org/x/text/unicode/rangetable"
)

func isNonSpacingMark(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}

var (
	printable                 = rangetable.Merge(unicode.PrintRanges...)
	UnicodeNonSpacingMarksSet = runes.In(unicode.Mn)
	UnicodePuncuationSet      = runes.In(unicode.Punct)
	UnicodeControlSet         = runes.In(unicode.C)
	UnicodePrintable          = runes.In(printable)

	nonSpacingOrPunctuationOrControl = runes.In(rangetable.Merge(unicode.Mn, unicode.Punct, unicode.C))
	UnicodeNonPrintable              = runes.NotIn(printable)

	removeNonSpacingMarksPunctuationAndControl = runes.Remove(nonSpacingOrPunctuationOrControl)
	removeUnicodeNonSpacingMarks               = runes.Remove(UnicodeNonSpacingMarksSet)
	removeUnicodePunctuation                   = runes.Remove(UnicodePuncuationSet)
	removeUnicodeControl                       = runes.Remove(UnicodeControlSet)
	removeUnicodeNonPrintable                  = runes.Remove(UnicodeNonPrintable)
)

//RemoveDiacriticsNFC creates a copy of s with the diacritics removed. It also transforms it to NFC.
//It is NOT thread Safe
func RemoveDiacriticsNFC(s string) string {
	var diacriticRemover = transform.Chain(norm.NFD, removeUnicodeNonSpacingMarks, norm.NFC)
	out, _, _ := transform.String(diacriticRemover, s)
	return out
}

//ExtremeNormalization heavily normalizes a string for purposes of comparison and safety.
//It lowercases the string, removes ALL nonspacing marks, nonprinting marks, whitespace, control characters, and punctuation, and transforms the string to NFKC encoding. This can and will lose a lot of information!
func ExtremeNormalization(s string) string {

	extremeNormalizer := transform.Chain( //this is created here because transform.Chain is not thread-safe
		norm.NFKD,
		removeNonSpacingMarksPunctuationAndControl,
		removeUnicodeNonPrintable,
		norm.NFKC,
	)
	s = strings.ToLower(s)
	s = RemoveASCIIWhiteSpace(s)
	s = RemoveASCIIPunctuation(s)
	s, _, _ = transform.String(extremeNormalizer, s)
	return s
}
