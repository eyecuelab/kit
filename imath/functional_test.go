package imath

import (
	"testing"

	"github.com/eyecuelab/kit/imath/operator"
	"github.com/stretchr/testify/assert"
)

func TestAccumulate(t *testing.T) {
	a := []int{1, 2, 3, 4}
	want := []int{1, 3, 6, 10}
	assert.Equal(t, want, Accumulate(operator.Add, a))
	assert.Empty(t, Accumulate(operator.Mul, []int(nil)))
}

func TestFilter(t *testing.T) {
	isEven := func(n int) bool { return n%2 == 0 }
	want := []int{0, 2, 4}
	assert.Equal(t, want, Filter(isEven, Range(0, 5, 1)))
}

func TestReduce(t *testing.T) {
	a := Range(1, 10, 1)
	got := Reduce(operator.Add, 0, a...)
	want := Sum(Range(0, 10, 1)...)
	assert.Equal(t, want, got)
}

func TestMap(t *testing.T) {
	double := func(n int) int { return 2 * n }
	want := []int{2, 4, 6, 8}
	assert.Equal(t, want, Map(double, []int{1, 2, 3, 4}))
}
