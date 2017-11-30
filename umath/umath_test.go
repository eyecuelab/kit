package umath

import (
	"testing"

	"github.com/eyecuelab/kit/random"
	"github.com/stretchr/testify/assert"
)

func TestMin(t *testing.T) {
	assert.Equal(t, uint(3), Min(5, 3))
}

func TestSum(t *testing.T) {
	assert.Equal(t, uint(10), Sum(2, 5, 3))
}

func TestRange(t *testing.T) {
	assert.Equal(t, []uint{0, 1, 2, 3, 4}, Range(0, 5, 1))
	assert.Equal(t, []uint{5, 3, 1}, Range(5, 0, -2))
	assert.Nil(t, Range(1, 0, 1))
}

func TestProduct(t *testing.T) {
	assert.Equal(t, uint(30), Product(5, 2, 3))
}

func TestMax(t *testing.T) {
	assert.Equal(t, uint(10), Max(2, 10, 5))
}

func TestClamp(t *testing.T) {
	const low, high uint = 1, 500
	assert.Equal(t, low, Clamp(0, low, high))
	assert.Equal(t, uint(2), Clamp(2, low, high))
	assert.Equal(t, high, Clamp(5000, low, high))
}

func TestPow(t *testing.T) {
	bases, err := random.Uint64s(10)
	assert.NoError(t, err)
	exps, err := random.Uint64s(10)
	assert.NoError(t, err)
	for i := range bases {
		base, exp := uint(bases[i]), uint(exps[i])
		got := Pow(base, exp)
		want := naivePow(base, exp)
		assert.Equal(t, want, got)
	}
}

func naivePow(base, exp uint) uint {
	result := uint(1)
	for ; exp > 0; exp-- {
		result *= base
	}
	return result
}
