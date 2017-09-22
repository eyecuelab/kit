package functools

func StringSliceContains(stringSlice []string, searchTerm string) bool {
	for _, s := range stringSlice {
		if s == searchTerm {
			return true
		}
	}
	return false
}
