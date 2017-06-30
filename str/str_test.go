package str

import (
	"testing"
)

func TestRemoveWhitespace(t *testing.T) {
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
		{"all", args{"\n\t\r "}, ""},
	}

	// TODO: Add test cases.

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveWhiteSpace(tt.args.s); got != tt.want {
				t.Errorf("RemoveWhitespace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveRunes(t *testing.T) {
	var alphabet = []rune("abcdefghijklmnopqrstuvwxyz")
	var digits = []rune("0123456789")
	type args struct {
		s     string
		runes []rune
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"remove_alphabet", args{"foo11829bar", alphabet}, "11829"},
		{"remove_digits", args{"foo11829bar", digits}, "foobar"},
		{"no_op", args{"foobar", []rune{}}, "foobar"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveRunes(tt.args.s, tt.args.runes...); got != tt.want {
				t.Errorf("RemoveRunes() = %v, want %v", got, tt.want)
			}
		})
	}
}
