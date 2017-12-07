package set

import (
	"reflect"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	_int_foo = iota
	_int_bar
	_int_baz
	_int_moo
)

var (
	_int_fooBar    = FromInts(_int_foo, _int_bar)
	_int_barBaz    = FromInts(_int_bar, _int_baz)
	_int_fooBaz    = FromInts(_int_foo, _int_baz)
	_int_fooBarBaz = FromInts(_int_foo, _int_bar, _int_baz)
)

func TestInt_Union(t *testing.T) {
	type args struct {
		strings []Int
	}

	tests := []struct {
		name      string
		s         Int
		args      args
		wantUnion Int
	}{
		{
			name:      "ok",
			s:         _int_fooBar,
			args:      args{[]Int{_int_barBaz}},
			wantUnion: _int_fooBarBaz,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotUnion := tt.s.Union(tt.args.strings...); !reflect.DeepEqual(gotUnion, tt.wantUnion) {
				t.Errorf("Int.Union() = %v, want %v", gotUnion, tt.wantUnion)
			}
		})
	}
}

func TestInt_Difference(t *testing.T) {
	type args struct {
		strings []Int
	}
	tests := []struct {
		name           string
		s              Int
		args           args
		wantDifference Int
	}{
		{
			name:           "ok",
			s:              _int_fooBarBaz,
			args:           args{[]Int{_int_fooBar}},
			wantDifference: FromInts(_int_baz),
		},
		{
			name:           "return self",
			s:              _int_fooBarBaz,
			wantDifference: _int_fooBarBaz,
		},
		{
			name: "multiples ok",
			s:    _int_fooBarBaz,
			args: args{[]Int{_int_fooBar, _int_barBaz}},
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

func TestInt_Add(t *testing.T) {
	type args struct {
		keys []int
	}
	tests := []struct {
		name string
		s    Int
		args args
		want Int
	}{
		{
			name: "add",
			s:    Int{_int_foo: yes},
			args: args{[]int{_int_foo, _int_foo, _int_bar, _int_baz}},
			want: _int_fooBarBaz,
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

func TestInt_ToSlice(t *testing.T) {
	tests := []struct {
		name     string
		s        Int
		want     []int
		wantFail bool
	}{
		{
			name: "ok",
			s:    _int_fooBar,
			want: []int{_int_foo, _int_bar},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.s.Sorted()
			if !reflect.DeepEqual(got, tt.want) && !tt.wantFail {
				t.Errorf("Int.ToSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_Intersection(t *testing.T) {
	type args struct {
		strings []Int
	}
	tests := []struct {
		name             string
		s                Int
		args             args
		wantIntersection Int
	}{
		{"ok", _int_fooBar, args{[]Int{_int_barBaz}}, FromInts(_int_bar)},
		{"multiples", _int_fooBar, args{[]Int{_int_fooBarBaz, _int_barBaz}}, FromInts(_int_bar)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Intersection(tt.args.strings...); !got.Equal(tt.wantIntersection) {
				t.Errorf("Int.Intersection() = %v, want %v", got, tt.wantIntersection)
			}
		})
	}
}
func TestInt_Delete(t *testing.T) {
	var set = make(Int)
	set.Add(_int_foo, _int_bar, _int_baz)
	assert.Equal(t, 3, len(set))
	set.Delete(_int_foo, _int_bar)
	assert.Equal(t, 1, len(set))
	set.Delete(_int_foo)
	assert.Equal(t, 1, len(set))
}

func TestInt_Equal(t *testing.T) {
	type args struct {
		other Int
	}
	tests := []struct {
		name string
		s    Int
		args args
		want bool
	}{
		{
			"yes",
			Int{_int_foo: yes, _int_bar: yes},
			args{Int{_int_foo: yes, _int_bar: yes}},
			true,
		},
		{
			"no",
			Int{_int_foo: yes, _int_bar: yes},
			args{Int{}},
			false,
		}, {
			"no - nonoverlapping keys",
			FromInts(_int_foo, _int_bar),
			args{other: FromInts(_int_baz, _int_moo)},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Equal(tt.args.other); got != tt.want {
				t.Errorf("Int.Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_IUnion(t *testing.T) {
	set := make(Int)
	a := FromInts(_int_foo, _int_bar)
	b := FromInts(_int_baz)

	set.IUnion(a)
	assert.Equal(t, set, a)
	set.IUnion(b)
	b.IUnion(a)
	assert.Equal(t, set, b)
}
func TestInt_XOR(t *testing.T) {
	type args struct {
		a Int
		b Int
	}
	tests := []struct {
		name string
		args args
		want Int
	}{
		{
			"ok",
			args{Int{_int_foo: yes, _int_bar: yes}, Int{_int_foo: yes, _int_baz: yes}},
			Int{_int_bar: yes, _int_baz: yes},
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

func TestInt_Map(t *testing.T) {
	double := func(n int) int {
		return 2 * n
	}
	set := FromInts(1, 2, 4, 8)
	want := FromInts(2, 4, 8, 16)
	got := set.Map(double)
	if !got.Equal(want) {
		t.Errorf("map broken: %v", got.XOR(want))
	}
}

func TestInt_Reduce(t *testing.T) {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	set := FromInts(-1, -5, 2, 8)
	want := -5
	got, _ := set.Reduce(min)
	if want != got {
		t.Errorf("reduce broken: should get %d, but got %d", want, got)
	}
	var emptySet Int
	if _, ok := emptySet.Reduce(min); ok {
		t.Errorf("should not get OK on empty set")
	}
}

func TestInt_Filter(t *testing.T) {
	nonneg := func(n int) bool {
		return n >= 0
	}
	set := FromInts(-1, 0, 2, 3)
	want := FromInts(0, 2, 3)
	got := set.Filter(nonneg)
	if !want.Equal(got) {
		t.Errorf("should get %v, but got %v", want, got)
	}
}

func TestInt_Contains(t *testing.T) {
	type args struct {
		key int
	}
	tests := []struct {
		name string
		s    Int
		args args
		want bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Contains(tt.args.key); got != tt.want {
				t.Errorf("Int.Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_Copy(t *testing.T) {
	tests := []struct {
		name string
		s    Int
		want Int
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Copy(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int.Copy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromInts(t *testing.T) {
	type args struct {
		ints []int
	}
	tests := []struct {
		name string
		args args
		want Int
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromInts(tt.args.ints...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromInts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_Sorted(t *testing.T) {
	tests := []struct {
		name string
		s    Int
		want []int
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Sorted(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int.Sorted() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_IsSubset(t *testing.T) {
	var (
		s = FromInts(2, 3)
		q = FromInts(2, 5)
		r = FromInts(2)
		w = FromInts(2, 5, 8, 9)
	)

	assert.True(t, r.IsSubset(s))
	assert.False(t, s.IsSubset(q))
	assert.False(t, w.IsSubset(s))
}

func TestInt_Pop(t *testing.T) {
	s := FromInts(5, 2, 3)
	b, _ := s.Pop()
	c, _ := s.Pop()
	d, _ := s.Pop()
	got := []int{b, c, d}
	sort.Ints(got)
	assert.Equal(t, []int{2, 3, 5}, got)
}

func TestInt_IsDisjoint(t *testing.T) {
	s, q := FromInts(2, 8), FromInts(3, 7)
	assert.True(t, s.IsDisjoint(q))
	r := FromInts(2)
	assert.False(t, s.IsDisjoint(r))
}

func TestInt_IsProperSubset(t *testing.T) {
	s, q := FromInts(2, 3), FromInts(2, 3, 5)
	assert.True(t, s.IsProperSubset(q))
	assert.False(t, s.IsProperSubset(s))
}

func TestInt_IsSuperset(t *testing.T) {

}

func TestInt_IsProperSuperset(t *testing.T) {

}
