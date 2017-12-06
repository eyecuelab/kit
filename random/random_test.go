package random

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const iterations = 250

/*
Note: it's not possible to test whether or not something is "really" random.
These tests have a very, very small chance of failing by chance.
You're more likely to win the lottery.
*/
func TestRandomBytes(t *testing.T) {
	var seen = make([][]byte, 0, iterations)
	for i := 0; i < iterations; i++ {
		b, err := RandomBytes(10)
		assert.NoError(t, err)
		for _, prev := range seen {
			assert.NotEqual(t, prev, b, "should not repeat")
		}
		seen = append(seen, b)
	}
}

func TestRandomString(t *testing.T) {
	var seen = make([]string, 0, iterations)
	for i := 0; i < iterations; i++ {
		s, err := RandomString(10)
		assert.NoError(t, err)
		for _, prev := range seen {
			assert.NotEqual(t, prev, s, "should not repeat")
		}
		seen = append(seen, s)
	}
}

func TestInt64s(t *testing.T) {
	var seen = make([][]int64, 0, iterations)
	for i := 0; i < iterations; i++ {
		s, err := Int64s(10)
		assert.NoError(t, err)
		for _, prev := range seen {
			assert.NotEqual(t, prev, s, "should not repeat")
		}
		seen = append(seen, s)
	}
}

func Test_uint64(t *testing.T) {
	const (
		b0 = 0x01
		b1 = 0x02
		b2 = 0x03
		b3 = 0x04
		b4 = 0x05
		b5 = 0x06
		b6 = 0x07
		b7 = 0x08

		want uint64 = b0<<000 + b1<<010 + b2<<020 +
			b3<<030 + b4<<040 + b5<<050 + b6<<060 + b7<<070
	)
	got := _uint64([]byte{b0, b1, b2, b3, b4, b5, b6, b7})
	assert.Equal(t, want, got)
}
