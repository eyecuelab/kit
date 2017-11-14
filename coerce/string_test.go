package coerce

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	const want = "foo"

	got, ok := String([]rune(want))
	assert.True(t, ok)
	assert.Equal(t, want, got)

	got, ok = String([]byte(want))
	assert.True(t, ok)
	assert.Equal(t, want, got)

	_, ok = String(22)
	assert.False(t, ok)

	got, ok = String(want)
	assert.Equal(t, want, got)
	assert.True(t, ok)

}
