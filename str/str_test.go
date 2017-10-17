package str

import (
	"reflect"
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

func TestSRemoveRunes(t *testing.T) {
	type args struct {
		s        string
		toRemove string
	}
	const accented_e = 'é'
	tests := []struct {
		name string
		args args
		want string
	}{

		{
			name: "ascii",
			args: args{s: NFC("Finé"), toRemove: ASCIILowercase},
			want: "Fé",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SRemoveRunes(tt.args.s, tt.args.toRemove); got != tt.want {
				t.Errorf("SRemoveRunes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapErr(t *testing.T) {
	type args struct {
		f       func(string) (string, error)
		strings []string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MapErr(tt.args.f, tt.args.strings)
			if (err != nil) != tt.wantErr {
				t.Errorf("MapErr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapErr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMap(t *testing.T) {
	type args struct {
		f       func(string) string
		strings []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Map(tt.args.f, tt.args.strings); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
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
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subIfNoChar(tt.args.r); got != tt.want {
				t.Errorf("subIfNoChar() = %v, want %v", got, tt.want)
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

func Test_max(t *testing.T) {
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
			if got := max(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("max() = %v, want %v", got, tt.want)
			}
		})
	}
}
