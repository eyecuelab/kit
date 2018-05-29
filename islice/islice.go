package islice


func StringPtrs(length int) []interface{} {
	from := make([]string, length)
	to := make([]interface{}, length)

	for i := 0; i < length; i++ {
		to[i] = &from[i]
	}
	return to
}