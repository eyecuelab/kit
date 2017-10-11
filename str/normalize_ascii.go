package str

import (
	"bytes"
	"strings"

	"github.com/eyecuelab/kit/runeset"
)

const (
	//ASCIIPunct is contains all ASCII punctuation, identical to string.punctuation in python 3.6
	ASCIIPunct = `$+<=>^|~!"#$%&\'()*+,-./:;<=>?@[\\]^_{|}~` + "`"

	//ASCIIWhitespace is a list of all ASCII whitespace, identical to string.Whitespace in python 3.6
	ASCIIWhitespace = " \t\n\r\x0b\x0c"

	//ASCIIPrintable is a list of all ASCII printable characters, identical to string.printable in python 3.6
	ASCIIPrintable = `0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!"#$%&\'()*+,-./:;<=>?@[\\]^_{|}~ \t\n\r\x0b\x0c` + "`"

	//ASCIINotPrintable is a list of all non-printable ASCII characters, identical ''.join(chr(x) for x in range(128) if chr(x) not in set(string.printable)) in python 3.6
	ASCIINotPrintable = `\x00\x01\x02\x03\x04\x05\x06\x07\x08\x0e\x0f\x10\x11\x12\x13\x14\x15\x16\x17\x18\x19\x1a\x1b\x1c\x1d\x1e\x1f\x7f`

	ASCIILowercase = `abcdefghijklmnopqrstuvwxyz`

	ASCIIUpperCase = `ABCDEFGHIJKLMNOPQRSTUVWXYZ`

	ASCIILetters = ASCIILowercase + ASCIIUpperCase

	ASCIINumerics = "0123456789"

	ASCIIAlphaNumeric = ASCIILowercase + ASCIIUpperCase + ASCIINumerics

	//ASCII is all ASCII characters, comprising the unicode code points 0-127.
	ASCII = "`" + `\x00\x01\x02\x03\x04\x05\x06\x07\x08\t\n\x0b\x0c\r\x0e\x0f\x10\x11\x12\x13\x14\x15\x16\x17\x18\x19\x1a\x1b\x1c\x1d\x1e\x1f !"#$%&\'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_abcdefghijklmnopqrstuvwxyz{|}~\x7f`
)

var (
	//ASCIIPunctSet  contains all ASCII punctuation. Equivalent ot set(string.punctuation) in python 3.6
	ASCIIPunctSet = runeset.FromString(ASCIIPunct)

	//ASCIIWhitespaceSet contains all ASCII whitespace, identical to set(string.whitespace) in python 3.6
	ASCIIWhitespaceSet = runeset.FromString(ASCIIWhitespace)
	//ASCIIPrintableSet contains  all ASCII printable characters, identical to string.printable in python 3.6
	ASCIIPrintableSet = runeset.FromString(ASCIIPrintable)

	//ASCIINotPrintableSet contains all non-printable ASCII characters, identical ''.join(chr(x) for x in range(128) if chr(x) not in set(string.printable)) in python 3.6
	ASCIINotPrintableSet = runeset.FromString(ASCIINotPrintable)

	ASCIIAlphaNumericSet = runeset.FromString(ASCIIAlphaNumeric)

	ASCIISet = runeset.FromString(ASCII)
)

//RemovePunctuation removes punctuation (as defined by unicode) from a string.
//Note that this converts to runes and back to UTF-8, so RemoveWhiteSpace(s) == s
//for a string that contains non-punctuation characters does not necessarially hold, since the code points may differ.
func RemoveASCIIPunctuation(s string) string {
	return removeRunesInSet(s, ASCIIPunctSet)
}

func RemoveASCIINonAlphaNumeric(s string) string {
	return removeRunesNotInSet(s, ASCIIAlphaNumericSet)
}

func RemoveASCIIWhiteSpace(s string) string {
	buf := bytes.Buffer{}
	for _, r := range s {
		if !ASCIIWhitespaceSet.Contains(r) {
			buf.WriteRune(r)
		}
	}
	return buf.String()
}

func RemoveASCIINonPrintable(s string) string {
	buf := bytes.Buffer{}
	for _, r := range s {
		if !ASCIINotPrintableSet.Contains(r) {
			buf.WriteRune(r)
		}
	}
	return buf.String()
}

//ASCIIHeavyNormalization removes all non-alphanumeric runes from s and lowercases the string
func ASCIIHeavyNormalization(s string) string {
	s = RemoveASCIINonAlphaNumeric(s)
	return strings.ToLower(s)
}

func RemoveNonASCII(s string) string {
	return removeRunesNotInSet(s, ASCIISet)
}
