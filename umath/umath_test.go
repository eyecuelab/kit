package umath

import (
	"math/rand"
	"testing"

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

func naivePow(base, exp uint) uint {
	result := uint(1)
	for ; exp > 0; exp-- {
		result *= base
	}
	return result
}

func TestPow(t *testing.T) {

	t.Run("random small pairs", func(t *testing.T) {
		for i := 0; i < 50; i++ {
			base := uint(rand.Intn(2 << 20))
			exp := uint(rand.Intn(2 << 16))
			if got, want := Pow(base, exp), naivePow(base, exp); got != want {
				t.Errorf("Pow got %v, but want %v", got, want)
			}
		}
	})
}
