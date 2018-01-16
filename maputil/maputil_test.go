package maputil

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

var m = map[string]interface{}{"a": 1, "b": 2, "c": "bar"}

func TestSortedKeys(t *testing.T) {
	want := []string{"a", "b", "c"}
	got := SortedKeys(m)
	assert.Equal(t, want, got)
}

func TestVals(t *testing.T) {
	want := []interface{}{1, 2, "bar"}
	got := Vals(m)
	sort.Sort(sorter(got))
	assert.Equal(t, want, got)
}

//sorter considers ints to be smaller than strings
type sorter []interface{}

func (s sorter) Less(i, j int) bool {
	if a, isInt := s[i].(int); isInt {
		if b, isInt := s[j].(int); isInt {
			return a < b
		}
		return true
	}
	if a, isStr := s[i].(string); isStr {
		if b, isStr := s[j].(string); isStr {
			return a < b
		}
		return false
	}
	panic("this shouldn't happen")
}
func (s sorter) Len() int      { return len(s) }
func (s sorter) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func TestCopy(t *testing.T) {
	start := map[string]interface{}{"1": 1}
	got := Copy(start)
	assert.Equal(t, start, got)
	start["2"] = 2
	assert.NotEqual(t, start, got) // no leak
}
