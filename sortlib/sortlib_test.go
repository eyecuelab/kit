package sortlib

import (
	"reflect"
	"testing"
)

func TestBytes(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{

		{
			name: "ok",
			args: args{[]byte{22, 81, 35}},
			want: []byte{22, 35, 81},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Bytes(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Bytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

const beta = "Β"

func TestByBytes(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want string
	}{

		{
			name: "ok",
			args: args{"aaZZAA"},
			want: "AAZZaa",
		},
		{
			name: "unicode",
			args: args{beta},
			want: string([]byte{146, 206}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ByBytes(tt.args.s); got != tt.want {
				t.Errorf("ByBytes() = %v, want %v", []uint8(got), (tt.want))
			}
		})
	}
}

func TestByRunes(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"ok",
			args{"aaZZAA"},
			"AAZZaa"},
		{
			"unicode",
			args{beta},
			beta,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ByRunes(tt.args.s); got != tt.want {
				t.Errorf("ByRunes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRunes(t *testing.T) {
	type args struct {
		r []rune
	}
	tests := []struct {
		name string
		args args
		want []rune
	}{
		{
			"ok",
			args{[]rune{22, 12312, 44}},
			[]rune{22, 44, 12312},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Runes(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Runes() = %v, want %v", got, tt.want)
			}
		})
	}
}
