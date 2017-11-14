package coerce

//String converts strings, []bytes, and []runes to strings
func String(v interface{}) (string, bool) {
	switch v := v.(type) {
	default:
		return "", false
	case string:
		return v, true
	case []byte:
		return string(v), true
	case []rune:
		return string(v), true
	}
}
