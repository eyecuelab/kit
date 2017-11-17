package functools

import "strings"

func StringContainsAny(str string, subStrs ...string) bool {
	for _, sub := range subStrs {
		if strings.Contains(str, sub) {
			return true
		}
	}
	return false
}
