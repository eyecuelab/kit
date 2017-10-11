package str

import (
	"testing"
)

func TestRemoveDiacriticsNFC(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{

		{
			name: "ok",
			arg:  "Finé",
			want: "Fine",
		},
		{
			name: "accented letters",
			arg:  "áéíóúüñ",
			want: "aeiouun",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveDiacriticsNFC(tt.arg); got != tt.want {
				t.Errorf("RemoveDiacriticsNFC() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveASCIIPunctuation(t *testing.T) {
	t.Run("ascii punctuation", func(t *testing.T) {
		if got := RemoveASCIIPunctuation(ASCIIPunct + "foo"); got != "foo" {
			t.Errorf("RemovePunctuation() = %v, want %v", got, "foo")
		}
	})
}

func TestExtremeNormalization(t *testing.T) {

	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "no whitespace",
			s:    ASCIIWhitespace,
			want: "",
		}, {
			name: "lowercase capitals",
			s:    "FOO",
			want: "foo",
		},
		{
			name: "no diacritics",
			s:    `Finé`,
			want: "fine",
		},
		{
			name: "no control characters",
			s:    "\x07", //bell
			want: "",
		},
		{
			name: "no punctuation",
			s:    ASCIIPunct,
			want: "",
		},
		{
			name: "no null",
			s:    "\x00",
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExtremeNormalization(tt.s); got != tt.want {
				t.Errorf("ExtremeNormalization() = %v, want %v", got, tt.want)
			}
		})
	}
}
