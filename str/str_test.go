package str

import (
	"reflect"
	"strings"
	"testing"

	"github.com/eyecuelab/kit/pretty"
)

func TestRemoveRunes(t *testing.T) {
	var alphabet = []rune(ASCIILowercase)
	var digits = []rune(ASCIINumerics)
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

func TestDiffs_String(t *testing.T) {
	tests := []struct {
		name string
		d    Diffs
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.String(); got != tt.want {
				t.Errorf("Diffs.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_runeDiff(t *testing.T) {
	type args struct {
		s []rune
		q []rune
	}
	tests := []struct {
		name string
		args args
		want Diffs
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := runeDiff(tt.args.s, tt.args.q); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("runeDiff() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDiff(t *testing.T) {
	type args struct {
		a string
		b string
	}
	const accent rune = 769
	tests := []struct {
		name string
		args args
		want Diffs
	}{
		{
			name: "ok",
			args: args{a: "Foo", b: "foo"},
			want: Diffs{RuneDiff{position: 0, a: 'F', b: 'f'}},
		},
		{
			name: "don't fold combining",
			args: args{a: NFC("Finé"), b: NFD("Finé")},
			want: Diffs{RuneDiff{position: 3, a: 'é', b: 'e'}, RuneDiff{position: 4, a: noChar, b: accent}},
		}, {
			name: "don't fold combining pt 2",
			args: args{a: NFD("Finé"), b: NFC("Finé")},
			want: Diffs{RuneDiff{position: 3, a: 'e', b: 'é'}, RuneDiff{position: 4, a: accent, b: noChar}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Diff(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				pretty.Print(pretty.Diff(got, tt.want))
				t.Errorf("got %v, wanted %v", got, tt.want)

			}
		})
	}
}

func Test_subIfNoChar(t *testing.T) {
	type args struct {
		r rune
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "ok",
			args: args{noChar},
			want: "NO_CHAR",
		},
		{name: "self",
			args: args{'f'},
			want: "f",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subIfNoChar(tt.args.r); got != tt.want {
				t.Errorf("subIfNoChar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRuneDiff_String(t *testing.T) {
	type fields struct {
		a        rune
		b        rune
		position int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rd := RuneDiff{
				a:        tt.fields.a,
				b:        tt.fields.b,
				position: tt.fields.position,
			}
			if got := rd.String(); got != tt.want {
				t.Errorf("RuneDiff.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_min(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := min(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("min() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMap(t *testing.T) {
	type args struct {
		f func(string) string
		a []string
	}

	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "capitalize",
			args: args{strings.ToUpper, []string{"lower", "upper"}},
			want: []string{"LOWER", "UPPER"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Map(tt.args.f, tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}
