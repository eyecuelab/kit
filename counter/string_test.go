package counter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	foo, bar, baz = "foo", "bar", "baz"
)

func TestString_Add(t *testing.T) {
	c := make(String)
	c.Add(foo)
	want := String{foo: 1}
	assert.Equal(t, want, c)
	want = String{foo: 2}
	c.Add(foo)
	assert.Equal(t, want, c)
}

func TestString_String(t *testing.T) {
	a := String{foo: 2, bar: 3}
	want := `{foo:2, bar:3}`
	assert.Equal(t, want, a.String())
}

func TestString_Sorted(t *testing.T) {
	c := String{"c": 1, "b": 2, "a": 2}
	gotKeys, gotCounts := c.Sorted()
	assert.Equal(t, []string{"c", "a", "b"}, gotKeys)
	assert.Equal(t, []int{1, 2, 2}, gotCounts)
}

func TestString_Combine(t *testing.T) {
	a := String{foo: 2, bar: 3}
	b := String{foo: 1, bar: 1, baz: 1}
	want := String{foo: 3, bar: 4, baz: 1}
	assert.Equal(t, want, a.Combine(b))
}

func TestString_Copy(t *testing.T) {
	a := String{foo: 1}
	b := a.Copy()
	b.Add(foo)
	assert.Equal(t, String{foo: 2}, b)
	assert.Equal(t, String{foo: 1}, a) //no leak
}

func TestMin(t *testing.T) {
	a := String{foo: 1, bar: 2, baz: -2}
	b := String{foo: -1}
	c := String{bar: -4}
	want := String{foo: -1, bar: -4, baz: -2}
	assert.Equal(t, want, Min(a, b, c))
	assert.Equal(t, want, Min(b, a, c))
}

func TestMax(t *testing.T) {
	a := String{foo: 1, bar: 2, baz: -2}
	b := String{foo: -1}
	c := String{bar: 4}
	want := String{foo: 1, bar: 4, baz: -2}
	assert.Equal(t, want, Max(a, b, c))
	assert.Equal(t, want, Max(b, a, c))
}

func TestFromStrings(t *testing.T) {
	a := []string{foo, bar, bar, bar, foo, baz}
	want := String{foo: 2, bar: 3, baz: 1}
	assert.Equal(t, want, FromStrings(a...))
}

func Test_StringEqual(t *testing.T) {
	a, b := String{foo: 2, bar: 3}, String{foo: 2, bar: 3}
	assert.True(t, a.Equal(b))

	c, d := String{}, String{foo: 2}
	assert.False(t, c.Equal(d))

	e, f := String{foo: 2, bar: 3}, String{foo: 3, bar: 3}
	assert.False(t, e.Equal(f))
}

func Test_StringPositiveElements(t *testing.T) {
	a := String{foo: 3, bar: -2, baz: 0}
	want := String{foo: 3}
	assert.Equal(t, want, a.PositiveElements())
}

func TestString_MostCommonN(t *testing.T) {
	wantKeys, wantVals := []string{foo, bar}, []int{3, 2}
	counter := String{foo: 3, bar: 2, baz: 2, "a": 1}
	gotKeys, gotVals, _ := counter.MostCommonN(2)
	assert.Equal(t, wantKeys, gotKeys)
	assert.Equal(t, wantVals, gotVals)
}

func TestString_MostCommon(t *testing.T) {
	wantKey, wantVal := foo, 3
	counter := String{foo: 3, bar: 2, baz: -1}
	gotKey, gotVal := counter.MostCommon()
	assert.Equal(t, wantKey, gotKey)
	assert.Equal(t, wantVal, gotVal)
}
