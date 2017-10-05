package str

import "testing"

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
