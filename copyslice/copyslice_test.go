package copyslice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInt(t *testing.T) {
	a := []int{2, 5}
	b := Int(a)
	assert.Equal(t, a, b)
	b[0] = 3
	assert.NotEqual(t, a, b)
}

func TestString(t *testing.T) {
	a := []string{"2", "5"}
	b := String(a)
	assert.Equal(t, a, b)
	b[0] = "3"
	assert.NotEqual(t, a, b)
}

func TestRune(t *testing.T) {
	a := []rune{2, 5}
	b := Rune(a)
	assert.Equal(t, a, b)
	b[0] = 3
	assert.NotEqual(t, a, b)
}

func TestByte(t *testing.T) {
	a := []byte{'2', '5'}
	b := Byte(a)
	assert.Equal(t, a, b)
	b[0] = '3'
	assert.NotEqual(t, a, b)
}

func TestInt64(t *testing.T) {
	a := []int64{2, 5}
	b := Int64(a)
	assert.Equal(t, a, b)
	b[0] = 3
	assert.NotEqual(t, a, b)
}

func TestFloat64(t *testing.T) {
	a := []float64{2, 5}
	b := Float64(a)
	assert.Equal(t, a, b)
	b[0] = 3
	assert.NotEqual(t, a, b)
}
