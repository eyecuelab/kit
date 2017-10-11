package str

import (
	"testing"
)

func TestRemoveASCIIWhitespace(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"spaces", args{"foo   "}, "foo"},
		{"tabs", args{"\tbar\t"}, "bar"},
		{"newlines", args{"\nba\nz\n"}, "baz"},
		{"none", args{"python"}, "python"},
		{"all", args{ASCIIWhitespace}, ""},
	}

	// TODO: Add test cases.

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveASCIIWhiteSpace(tt.args.s); got != tt.want {
				t.Errorf("RemoveWhitespace() = %v, want %v", got, tt.want)
			}
		})
	}
}



func TestASCIIHeavyNormalization(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ASCIIHeavyNormalization(tt.args.s); got != tt.want {
				t.Errorf("ASCIIHeavyNormalization() = %v, want %v", got, tt.want)
			}
		})
	}
}
