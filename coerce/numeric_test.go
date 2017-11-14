package coerce

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFloat64(t *testing.T) {
	AllowImpreciseConversion(true)
	a := 2<<55 + 1
	x, err := Float64(a)
	assert.NoError(t, err)
	assert.NotEqual(t, int(x), a)

	AllowImpreciseConversion(false)
	b := 2<<52 - 1
	x, err = Float64(b)
	assert.NoError(t, err)
	assert.Equal(t, float64(b), x)

	_, err = Float64(int64(a))
	assert.Error(t, err)

	_, err = Float64(uint64(a))
	assert.Error(t, err)

	_, err = Float64(uint(a))
	assert.Error(t, err)

	_, err = Float64(a)
	assert.Error(t, err)

	AllowImpreciseConversion(true)
	_, err = Float64("foo")
	assert.Error(t, err)
}

func TestInt64(t *testing.T) {
	AllowOverflow(true)
	a := uint(math.MaxInt64) + 1
	x, err := Int64(a)
	assert.NoError(t, err)
	assert.NotEqual(t, int64(x), a)

	AllowOverflow(false)
	x, err = Int64(a)
	assert.Error(t, err)
}
