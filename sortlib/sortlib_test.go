package sortlib

import (
	"testing"
)

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
