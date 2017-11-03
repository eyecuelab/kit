package set

import (
	"reflect"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	_foo = "_foo"
	_bar = "_bar"
	_baz = "_baz"
	_moo = "_moo"
)

var (
	fooBar    = FromStrings(_foo, _bar)
	barBaz    = FromStrings(_bar, _baz)
	fooBaz    = FromStrings(_foo, _baz)
	fooBarBaz = FromStrings(_foo, _bar, _baz)
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
			s:         fooBar,
			args:      args{[]String{barBaz}},
			wantUnion: fooBarBaz,
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
			s:              fooBarBaz,
			args:           args{[]String{fooBar}},
			wantDifference: FromStrings(_baz),
		},
		{
			name:           "return self",
			s:              fooBarBaz,
			wantDifference: fooBarBaz,
		},
		{
			name: "multiples ok",
			s:    fooBarBaz,
			args: args{[]String{fooBar, barBaz}},
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
			s:    String{_foo: yes},
			args: args{[]string{_foo, _foo, _bar, _baz}},
			want: fooBarBaz,
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

func TestString_ToSlice(t *testing.T) {
	tests := []struct {
		name     string
		s        String
		want     []string
		wantFail bool
	}{
		{
			name: "ok",
			s:    fooBar,
			want: []string{_foo, _bar},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.s.ToSlice()
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
		{"ok", fooBar, args{[]String{barBaz}}, FromStrings(_bar)},
		{"multiples", fooBar, args{[]String{fooBarBaz, barBaz}}, FromStrings(_bar)},
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
	set.Add(_foo, _bar, _baz)
	assert.Equal(t, 3, len(set))
	set.Remove(_foo, _bar)
	assert.Equal(t, 1, len(set))
	set.Remove(_foo)
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
			String{_foo: yes, _bar: yes},
			args{String{_foo: yes, _bar: yes}},
			true,
		},
		{
			"no",
			String{_foo: yes, _bar: yes},
			args{String{}},
			false,
		}, {
			"no - nonoverlapping keys",
			FromStrings(_foo, _bar),
			args{other: FromStrings(_baz, _moo)},
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
	a := FromStrings(_foo, _bar)
	b := FromStrings(_baz)

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
			args{String{_foo: yes, _bar: yes}, String{_foo: yes, _baz: yes}},
			String{_bar: yes, _baz: yes},
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
