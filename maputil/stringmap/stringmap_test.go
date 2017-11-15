package stringmap

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

var m = map[string]string{
	"a": "a",
	"b": "b",
	"c": "c",
}

func TestSortedKeys(t *testing.T) {
	want := []string{"a", "b", "c"}
	got := SortedKeys(m)
	assert.Equal(t, want, got)
}

func TestVals(t *testing.T) {
	want := []string{"a", "b", "c"}
	got := Vals(m)
	sort.Strings(got)
	assert.Equal(t, want, got)
}
