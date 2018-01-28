package functools

import "testing"
import "github.com/stretchr/testify/assert"

const a, b, c = true, true, false

func TestAll(t *testing.T) {
	assert.False(t, All(a, b, c))
	assert.True(t, All(a, b))
	assert.True(t, All())
}

func TestAny(t *testing.T) {
	assert.True(t, Any(a, b, c))
	assert.True(t, Any(a, b))
	assert.False(t, Any())
}

func TestNone(t *testing.T) {
	assert.True(t, None())
	assert.True(t, None(c))
	assert.False(t, None(a, b, c))
}
