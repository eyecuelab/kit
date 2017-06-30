package runes

//Set returns a set of runes: i.e, a map[rune][bool]
func Set(runes ...rune) map[rune]bool {
	set := make(map[rune]bool)
	for _, r := range runes {
		set[r] = true
	}
	return set
}
