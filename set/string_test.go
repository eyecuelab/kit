package set

import (
	"reflect"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	_str_foo = "a"
	_str_bar = "b"
	_str_baz = "c"
	_str_moo = "d"
)

var (
	_str_fooBar    = FromStrings(_str_foo, _str_bar)
	_str_barBaz    = FromStrings(_str_bar, _str_baz)
	_str_fooBaz    = FromStrings(_str_foo, _str_baz)
	_str_fooBarBaz = FromStrings(_str_foo, _str_bar, _str_baz)
)

func TestString_Union(t *testing.T) {
	type args struct {
		strings []String
	}

	tests := []struct {
		name      string
		s         String
		args      args
		wantUnion String
	}{
		{
			name:      "ok",
			s:         _str_fooBar,
			args:      args{[]String{_str_barBaz}},
			wantUnion: _str_fooBarBaz,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotUnion := tt.s.Union(tt.args.strings...); !reflect.DeepEqual(gotUnion, tt.wantUnion) {
				t.Errorf("String.Union() = %v, want %v", gotUnion, tt.wantUnion)
			}
		})
	}
}

func TestString_Difference(t *testing.T) {
	type args struct {
		strings []String
	}
	tests := []struct {
		name           string
		s              String
		args           args
		wantDifference String
	}{
		{
			name:           "ok",
			s:              _str_fooBarBaz,
			args:           args{[]String{_str_fooBar}},
			wantDifference: FromStrings(_str_baz),
		},
		{
			name:           "return self",
			s:              _str_fooBarBaz,
			wantDifference: _str_fooBarBaz,
		},
		{
			name: "multiples ok",
			s:    _str_fooBarBaz,
			args: args{[]String{_str_fooBar, _str_barBaz}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.s.Difference(tt.args.strings...)
			if !got.Equal(tt.wantDifference) {
				t.Errorf("got %v, but want %v", got, tt.wantDifference)
			}
		})
	}
}

func TestString_Add(t *testing.T) {
	type args struct {
		keys []string
	}
	tests := []struct {
		name string
		s    String
		args args
		want String
	}{
		{
			name: "add",
			s:    String{_str_foo: yes},
			args: args{[]string{_str_foo, _str_foo, _str_bar, _str_baz}},
			want: _str_fooBarBaz,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Add(tt.args.keys...)
			if !tt.s.Equal(tt.want) {
				t.Errorf("reciever of s.Add() is %v, but should be %v", tt.s, tt.want)
			}
		})
	}
}

func TestString_Sorted(t *testing.T) {
	tests := []struct {
		name     string
		s        String
		want     []string
		wantFail bool
	}{
		{
			name: "ok",
			s:    _str_fooBar,
			want: []string{_str_foo, _str_bar},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.s.Sorted()
			sort.Strings(got)
			sort.Strings(tt.want)
			if !reflect.DeepEqual(got, tt.want) && !tt.wantFail {
				t.Errorf("String.ToSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestString_Intersection(t *testing.T) {
	type args struct {
		strings []String
	}
	tests := []struct {
		name             string
		s                String
		args             args
		wantIntersection String
	}{
		{"ok", _str_fooBar, args{[]String{_str_barBaz}}, FromStrings(_str_bar)},
		{"multiples", _str_fooBar, args{[]String{_str_fooBarBaz, _str_barBaz}}, FromStrings(_str_bar)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Intersection(tt.args.strings...); !got.Equal(tt.wantIntersection) {
				t.Errorf("String.Intersection() = %v, want %v", got, tt.wantIntersection)
			}
		})
	}
}
func TestString_Remove(t *testing.T) {
	var set = make(String)
	set.Add(_str_foo, _str_bar, _str_baz)
	assert.Equal(t, 3, len(set))
	set.Remove(_str_foo, _str_bar)
	assert.Equal(t, 1, len(set))
	set.Remove(_str_foo)
	assert.Equal(t, 1, len(set))
}

func TestString_Equal(t *testing.T) {
	type args struct {
		other String
	}
	tests := []struct {
		name string
		s    String
		args args
		want bool
	}{
		{
			"yes",
			String{_str_foo: yes, _str_bar: yes},
			args{String{_str_foo: yes, _str_bar: yes}},
			true,
		},
		{
			"no",
			String{_str_foo: yes, _str_bar: yes},
			args{String{}},
			false,
		}, {
			"no - nonoverlapping keys",
			FromStrings(_str_foo, _str_bar),
			args{other: FromStrings(_str_baz, _str_moo)},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Equal(tt.args.other); got != tt.want {
				t.Errorf("String.Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestString_IUnion(t *testing.T) {
	set := make(String)
	a := FromStrings(_str_foo, _str_bar)
	b := FromStrings(_str_baz)

	set.IUnion(a)
	assert.Equal(t, set, a)
	set.IUnion(b)
	b.IUnion(a)
	assert.Equal(t, set, b)
}
func TestString_XOR(t *testing.T) {
	type args struct {
		a String
		b String
	}
	tests := []struct {
		name string
		args args
		want String
	}{
		{
			"ok",
			args{String{_str_foo: yes, _str_bar: yes}, String{_str_foo: yes, _str_baz: yes}},
			String{_str_bar: yes, _str_baz: yes},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.a.XOR(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("XOR() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestString_Reduce(t *testing.T) {
	min := func(a, b string) string {
		if a < b {
			return a
		}
		return b
	}

	set := FromStrings("aaa", "ab", "ccc")
	want := "aaa"
	got, _ := set.Reduce(min)
	if want != got {
		t.Errorf("reduce broken: should get %s, but got %s", want, got)
	}
	var emptySet String
	if _, ok := emptySet.Reduce(min); ok {
		t.Errorf("should not get OK on empty set")
	}
}

func TestString_Map(t *testing.T) {
	double := func(s string) string {
		out := ""
		for _, r := range s {
			out += string(r) + string(r)
		}
		return out
	}
	set := FromStrings("ab", "cd")
	want := FromStrings("aabb", "ccdd")
	got := set.Map(double)
	if !got.Equal(want) {
		t.Errorf("map broken: %v", got.XOR(want))
	}
}

func TestString_Filter(t *testing.T) {
	capital := func(s string) bool {
		for _, r := range s {
			if r < 'A' || r > 'Z' {
				return false
			}
		}
		return true
	}
	set := FromStrings("AAA", "BBB", "crasn", "44s")
	want := FromStrings("AAA", "BBB")
	got := set.Filter(capital)
	if !want.Equal(got) {
		t.Errorf("should get %v, but got %v", want, got)
	}
}
