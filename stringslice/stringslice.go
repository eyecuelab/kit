package stringslice

//NonEmpty returns a slice containing the non-empty elements of a
func NonEmpty(a []string) []string {
	var nonempty []string
	for _, s := range a {
		if s != "" {
			nonempty = append(nonempty, s)
		}
	}
	return nonempty
}
