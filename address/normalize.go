package address

import (
	"strings"

	"github.com/eyecuelab/kit/str"
)

func normalizeComponent(s string) string {
	s = str.ExtremeNormalization(s)
	return NormalizeSuffixes(s)
}

var suffixReplacer = makeSuffixReplacer()

func makeSuffixReplacer() *strings.Replacer {
	pairs := make([]string, 0, len(suffixVariations)*2)
	for key, val := range suffixVariations {
		pairs = append(pairs, key, val)
	}
	return strings.NewReplacer(pairs...)

}

//NormalizeSuffixes normalizes lowercase address suffixes to the standard format used by the United States Postal Service.
//i.e, "pkwy": "parkway",
//see http://www.pb.com/docs/us/pdf/sis/mail-services/usps-suffix-abbreviations.pdf
func NormalizeSuffixes(s string) string {
	return suffixReplacer.Replace(s)
}
