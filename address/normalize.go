package address

import (
	"strings"

	"github.com/eyecuelab/kit/str"
)

var shortForm = map[string]string{
	"street": "st",
	"drive":  "dr",
	"north":  "n",
	"south":  "s",
	"west":   "w",
	"east":   "e",
}

type signal interface{}

var yes signal

func isNoise(s string) bool {
	switch s {
	case "ste", "apt", "#", "suite":
		return true
	}
	return false
}

func removeNoise(strings []string) []string {
	var components []string
	for _, s := range strings {
		if !isNoise(s) {
			components = append(components, s)
		}
	}
	return components
}

func shortForms(components []string) []string {
	for i, s := range components {
		if short, ok := shortForm[s]; ok {
			components[i] = short
		} else {
			components[i] = s
		}
	}
	return components
}

func removeNoiseAndShorten(s string) string {
	cleaned := shortForms(removeNoise(strings.Split(s, " ")))
	return strings.Join(cleaned, " ")
}

//Normalize all or part of an address component
func Normalize(component string) string {
	component = str.NKFC(component)
	component = str.RemoveRunes(component, ';', ':', '.', ',', '!', '?', '*', '#')
	component = strings.TrimSpace(component)
	component = strings.ToLower(component)
	return removeNoiseAndShorten(component)
}
