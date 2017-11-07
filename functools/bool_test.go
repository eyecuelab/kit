package functools

import "testing"
import "github.com/stretchr/testify/assert"

func TestAll(t *testing.T) {
	a, b, c := true, true, false
	assert.False(t, All(a, b, c))
	assert.True(t, All(a, b))
	assert.True(t, All())
}

func TestAny(t *testing.T) {
	a, b, c := true, true, false
	assert.True(t, Any(a, b, c))
	assert.True(t, Any(a, b))
	assert.False(t, Any())
}
