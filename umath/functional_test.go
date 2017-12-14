package umath

// import (
// 	"testing"
//
// 	"github.com/stretchr/testify/assert"
// )
//
// func isEven(n uint) bool { return n%2 == 0 }
// func TestFilter(t *testing.T) {
// 	a := []uint{2, 8, 7}
// 	assert.Equal(t, []uint{2, 8}, Filter(isEven, a))
//
// }
//
// func leftshift(n uint) uint { return n << 1 }
//
// func TestMap(t *testing.T) {
// 	a := Range(0, 5, 1)
// 	want := Range(0, 10, 2)
// 	assert.Equal(t, want, Map(leftshift, a))
// }
//
// func add(a, b uint) uint { return a + b }
// func TestReduce(t *testing.T) {
// 	a := Range(0, 5, 1)
// 	assert.Equal(t, Sum(a...), Reduce(add, a[0], a[1:]...))
// }
//
// func TestAccumulate(t *testing.T) {
//
// 	a := Range(0, 5, 1)
// 	want := []uint{0, 1, 3, 6, 10}
// 	assert.Equal(t, want, Accumulate(add, a))
//
// 	assert.Empty(t, Accumulate(add, nil))
// }
