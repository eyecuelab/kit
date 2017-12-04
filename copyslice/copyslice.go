//Package copyslice contains functions to create copies of various slice types.
package copyslice

//Int creates a copy of a slice of ints
func Int(src []int) []int {
	dst := make([]int, len(src))
	copy(dst, src)
	return dst
}

//String copies a slice of strings
func String(src []string) []string {
	dst := make([]string, len(src))
	copy(dst, src)
	return dst
}

//Rune copies a slice of runes
func Rune(src []rune) []rune {
	dst := make([]rune, len(src))
	copy(dst, src)
	return dst
}
