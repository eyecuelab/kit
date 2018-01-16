//Package copyslice contains functions to create copies of various slice types.
package copyslice

//Int copies  a []int
func Int(src []int) []int {
	dst := make([]int, len(src))
	copy(dst, src)
	return dst
}

//Uint copies  a []uint
func Uint(src []uint) []uint {
	dst := make([]uint, len(src))
	copy(dst, src)
	return dst
}

//String copies a []string
func String(src []string) []string {
	dst := make([]string, len(src))
	copy(dst, src)
	return dst
}

//Rune copies a []rune
func Rune(src []rune) []rune {
	dst := make([]rune, len(src))
	copy(dst, src)
	return dst
}

//Byte copies a []byte
func Byte(src []byte) []byte {
	dst := make([]byte, len(src))
	copy(dst, src)
	return dst
}

//Int64 copies a []int64
func Int64(src []int64) []int64 {
	dst := make([]int64, len(src))
	copy(dst, src)
	return dst
}

//Float64 copies a []Float64
func Float64(src []float64) []float64 {
	dst := make([]float64, len(src))
	copy(dst, src)
	return dst
}

//Uint64 copies a []Uint64
func Uint64(src []uint64) []uint64 {
	dst := make([]uint64, len(src))
	copy(dst, src)
	return dst
}
