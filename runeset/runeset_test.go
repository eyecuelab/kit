package runeset

import (
	"reflect"
	"testing"

	"golang.org/x/text/unicode/norm"
)

const ascii_lowercase = "abcdefghijklmnopqrstuvwxyz"

var (
	abc = FromString("abc")
	a   = FromString("a")
	ab  = FromString("ab")
	bc  = FromString("bc")
	ac  = FromString("ac")
)

func TestRuneSet_Contains(t *testing.T) {
	type args struct {
		r rune
	}
	tests := []struct {
		name string
		rs   RuneSet
		args args
		want bool
	}{
		{
			name: "yes",
			rs:   abc,
			args: args{'c'},
			want: true,
		},
		{
			name: "no",
			rs:   abc,
			args: args{'d'},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.rs.Contains(tt.args.r); got != tt.want {
				t.Errorf("RuneSet.Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRuneSet_Intersection(t *testing.T) {
	type args struct {
		sets []RuneSet
	}
	tests := []struct {
		name             string
		rs               RuneSet
		args             args
		wantIntersection RuneSet
	}{
		{

			name:             "ok",
			rs:               abc,
			args:             args{[]RuneSet{ab, a}},
			wantIntersection: a,
		}, {
			name:             "empty",
			rs:               abc,
			args:             args{[]RuneSet{ab, bc, ac}},
			wantIntersection: RuneSet{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotIntersection := tt.rs.Intersection(tt.args.sets...); !reflect.DeepEqual(gotIntersection, tt.wantIntersection) {
				t.Errorf("RuneSet.Intersection() = %v, want %v", gotIntersection, tt.wantIntersection)
			}
		})
	}
}

func TestRuneSet_Equal(t *testing.T) {
	type args struct {
		other RuneSet
	}

	tests := []struct {
		name string
		rs   RuneSet
		args args
		want bool
	}{

		{name: "ok",
			rs:   FromString("abcabc"),
			args: args{FromString("abc")},
			want: true,
		},
		{
			name: "no - len",
			rs:   FromString(norm.NFC.String("Finé")),
			args: args{FromString(norm.NFD.String("Finé"))},
			want: false,
		},
		{
			name: "no - different chars",
			rs:   FromString("s"),
			args: args{FromString("p")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.rs.Equal(tt.args.other); got != tt.want {
				t.Errorf("RuneSet.Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRuneSet_Union(t *testing.T) {
	type args struct {
		sets []RuneSet
	}
	tests := []struct {
		name      string
		rs        RuneSet
		args      args
		wantUnion RuneSet
	}{
		{
			name:      "ok",
			rs:        RuneSet{},
			args:      args{[]RuneSet{bc, a}},
			wantUnion: abc,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotUnion := tt.rs.Union(tt.args.sets...); !reflect.DeepEqual(gotUnion, tt.wantUnion) {
				t.Errorf("RuneSet.Union() = %v, want %v", gotUnion, tt.wantUnion)
			}
		})
	}
}

func TestRuneSet_Difference(t *testing.T) {
	type args struct {
		sets []RuneSet
	}

	abc := FromString("abc")
	a := FromString("a")
	bc := FromString("bc")
	tests := []struct {
		name           string
		rs             RuneSet
		args           args
		wantDifference RuneSet
	}{
		{
			name:           "two diff",
			rs:             abc,
			args:           args{[]RuneSet{bc}},
			wantDifference: a,
		},
		{name: "three diff",
			rs:             abc,
			args:           args{[]RuneSet{a, bc}},
			wantDifference: RuneSet{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDifference := tt.rs.Difference(tt.args.sets...); !reflect.DeepEqual(gotDifference, tt.wantDifference) {
				t.Errorf("RuneSet.Difference() = %v, want %v", gotDifference, tt.wantDifference)
			}
		})
	}
}

func TestFromRunes(t *testing.T) {
	type args struct {
		runes []rune
	}
	const bigRune = rune(0xFFFF)
	tests := []struct {
		name string
		args args
		want RuneSet
	}{
		{
			name: "dedupe",
			args: args{append([]rune("abcabc"), bigRune)},
			want: RuneSet{
				'a':     yes,
				'b':     yes,
				'c':     yes,
				bigRune: yes,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromRunes(tt.args.runes...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromRunes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromString(t *testing.T) {
	const bigRune = rune(0xFFFF)
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantSet RuneSet
	}{
		{name: "ok",
			args: args{"abcabcabc" + string(bigRune)},
			wantSet: RuneSet{
				'a':     yes,
				'b':     yes,
				'c':     yes,
				bigRune: yes,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSet := FromString(tt.args.s); !reflect.DeepEqual(gotSet, tt.wantSet) {
				t.Errorf("FromString() = %v, want %v", gotSet, tt.wantSet)
			}
		})
	}
}

func TestIntersection(t *testing.T) {
	type args struct {
		set  RuneSet
		sets []RuneSet
	}
	tests := []struct {
		name string
		args args
		want RuneSet
	}{
		{
			name: "ok",
			args: args{
				set:  abc,
				sets: []RuneSet{bc},
			},
			want: bc,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Intersection(tt.args.set, tt.args.sets...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Intersection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnion(t *testing.T) {
	type args struct {
		sets []RuneSet
	}
	tests := []struct {
		name string
		args args
		want RuneSet
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Union(tt.args.sets...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Union() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRuneSet_Copy(t *testing.T) {
	tests := []struct {
		name string
		rs   RuneSet
		want RuneSet
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.rs.Copy(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RuneSet.Copy() = %v, want %v", got, tt.want)
			}
		})
	}
}
