package str

import (
	"bytes"

	"github.com/eyecuelab/kit/runeset"
)

const (
	//ASCIIPunct is contains all ASCII punctuation, identical to string.punctuation in python 3.6
	ASCIIPunct = `$+<=>^|~!"#$%&\'()*+,-./:;<=>?@[\\]^_{|}~` + "`"

	//ASCIIWhitespace is a list of all ASCII whitespace, identical to string.Whitespace in python 3.6
	ASCIIWhitespace = " \t\n\r\x0b\x0c"

	//ASCIIPrintable is a list of all ASCII printable characters, identical to string.printable in python 3.6
	ASCIIPrintable = `0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!"#$%&\'()*+,-./:;<=>?@[\\]^_{|}~ \t\n\r\x0b\x0c` + "`"

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

	ASCIISet = runeset.FromString(ASCII)
)

func printableASCII() runeset.RuneSet {
	printable := make(runeset.RuneSet)
	var yes runeset.Signal
	for b := rune(32); b < 128; b++ {
		printable[b] = yes
	}
	return printable
}

func nonprintableASCII() runeset.RuneSet {
	nonprintable := make(runeset.RuneSet)
	var yes runeset.Signal
	for b := rune(0); b < 32; b++ {
		nonprintable[b] = yes
	}
	return nonprintable
}

//RemovePunctuation removes punctuation (as defined by unicode) from a string.
//Note that this converts to runes and back to UTF-8, so RemoveWhiteSpace(s) == s
//for a string that contains non-punctuation characters does not necessarially hold, since the code points may differ.
func RemoveASCIIPunctuation(s string) string {
	return removeRunesInSet(s, ASCIIPunctSet)
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

func RemoveNonASCII(s string) string {
	return removeRunesNotInSet(s, ASCIISet)
}
