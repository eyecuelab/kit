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
