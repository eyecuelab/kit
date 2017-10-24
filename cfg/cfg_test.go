package cfg

import "testing"

func TestBeenSet(t *testing.T) {

	type unknownType struct {
		a int32
	}

	var defaultUnknown unknownType
	var nonDefaultUnknown = unknownType{5}
	type args struct {
		v interface{}
	}
	var empty interface{}
	tests := []struct {
		name string
		args args
		want bool
	}{

		{
			name: "empty interface",
			args: args{empty},
			want: false,
		},
		{
			name: "nil slice",
			args: args{[]int(nil)},
			want: false,
		}, {
			name: "empty slice",
			args: args{[]int{}},
			want: true,
		},
		{
			name: "unknown type - default",
			args: args{defaultUnknown},
			want: false,
		},
		{
			name: "unknown type - non default",
			args: args{nonDefaultUnknown},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BeenSet(tt.args.v); got != tt.want {
				t.Errorf("BeenSet() = %v, want %v", got, tt.want)
			}
		})
	}
}
