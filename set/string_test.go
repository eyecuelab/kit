package set

import (
	"reflect"
	"testing"
)

const (
	foo = "foo"
	bar = "bar"
	baz = "baz"
)

var (
	fooBar    = FromStrings(foo, bar)
	barBaz    = FromStrings(bar, baz)
	fooBaz    = FromStrings(foo, baz)
	fooBarBaz = FromStrings(foo, bar, baz)
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
			wantDifference: FromStrings(baz),
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

		// TODO: Add test cases.
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
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Add(tt.args.keys...)
		})
	}
}

func TestFromStringSlice(t *testing.T) {
	type args struct {
		stringSlices [][]string
	}
	tests := []struct {
		name string
		args args
		want []String
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromStringSlice(tt.args.stringSlices...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromStringSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestString_ToSlice(t *testing.T) {
	tests := []struct {
		name string
		s    String
		want []string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ToSlice(); !reflect.DeepEqual(got, tt.want) {
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
		{"ok", fooBar, args{[]String{barBaz}}, FromStrings(bar)},
		{"multiples", fooBar, args{[]String{fooBarBaz, barBaz}}, FromStrings(bar)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Intersection(tt.args.strings...); !got.Equal(tt.wantIntersection) {
				t.Errorf("String.Intersection() = %v, want %v", got, tt.wantIntersection)
			}
		})
	}
}
