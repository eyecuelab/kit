package functools

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	foo = "foo"
	bar = "bar"
	baz = "baz"
)

func TestStringSliceContains(t *testing.T) {
	assert.True(t, StringSliceContains([]string{foo, bar, baz}, bar))
	assert.False(t, StringSliceContains([]string{foo, bar, baz}, "moo"))
}
