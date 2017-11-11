package random

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const iterations = 250

/*
Note: it's not possible to test whether or not something is "really" random.
These tests have a very, very small chance of failing by chance.
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
