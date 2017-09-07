package stringslice

import (
	"reflect"
	"testing"
)

func TestNonEmpty(t *testing.T) {
	type args struct {
		a []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"nil", args{nil}, nil},
		{"full", args{[]string{"foo", "bar", "baz"}}, []string{"foo", "bar", "baz"}},
		{"missing pieces", args{[]string{"0", "", "1"}}, []string{"0", "1"}},
		{"nonempty slice with all blank args", args{[]string{"", "", ""}}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NonEmpty(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NonEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}
