package islice

func FromStringSlice(strs []string) []interface{} {
	converted := make([]interface{}, len(strs))
	for i, s := range strs {
		converted[i] = s
	}
	return converted
}

func FromIntSlice(ints []int) []interface{} {
	converted := make([]interface{}, len(ints))
	for i, int := range ints {
		converted[i] = int
	}
	return converted
}

