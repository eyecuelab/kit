package stringslice

import (
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNonEmpty(t *testing.T) {
	type args struct {
		a []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"nil", args{nil}, nil},
		{"full", args{[]string{"foo", "bar", "baz"}}, []string{"foo", "bar", "baz"}},
		{"missing pieces", args{[]string{"0", "", "1"}}, []string{"0", "1"}},
		{"nonempty slice with all blank args", args{[]string{"", "", ""}}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NonEmpty(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NonEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func nonempty(s string) bool {
	return len(s) > 0
}
func TestAll(t *testing.T) {
	yes := []string{"a", "foo", "bar"}
	no := []string{"a", "", "c"}
	var noelems []string
	assert.True(t, All(yes, nonempty))
	assert.False(t, All(no, nonempty))
	assert.True(t, All(noelems, nonempty))

}

func TestAny(t *testing.T) {
	yes := []string{"", "c", ""}
	no := []string{"", "", ""}
	var noelems []string
	assert.True(t, Any(yes, nonempty))
	assert.False(t, Any(no, nonempty))
	assert.False(t, Any(noelems, nonempty))
}

func TestFilter(t *testing.T) {
	a := []string{"", "c", ""}
	want := []string{"c"}
	assert.Equal(t, want, Filter(a, nonempty))
}

func TestFilterFalse(t *testing.T) {
	a := []string{"", "c", ""}
	want := []string{"", ""}
	assert.Equal(t, want, FilterFalse(a, nonempty))
}

func TestMap(t *testing.T) {
	got := Map([]string{"aa", "bcd"}, strings.ToUpper)
	want := []string{"AA", "BCD"}
	assert.Equal(t, want, got)
}

func TestAppendIfNonEmpty(t *testing.T) {
	want := []string{"foo", "aa", "bb"}
	toAppend := []string{"aa", "", "", "bb"}
	assert.Equal(t, want, AppendIfNonEmpty([]string{"foo"}, toAppend...))
}

func TestTotalDistance(t *testing.T) {
	_, ok := TotalDistance([]string{}, []string{"aa"})
	assert.False(t, ok)

	a := []string{"foo", "bar", "baz"}
	b := []string{"food", "BAR", "BASZ"}
	want := 1 + 3 + 4
	got, ok := TotalDistance(a, b)
	assert.True(t, ok)
	assert.Equal(t, want, got)

}

func sum(ints ...int) int {
	total := 0
	for _, n := range ints {
		total += n
	}
	return total
}

func Test_levenshteinDistance(t *testing.T) {
	type args struct {
		s string
		t string
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
			if got := levenshteinDistance(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("levenshteinDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_min(t *testing.T) {
	type args struct {
		a    int
		ints []int
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
			if got := min(tt.args.a, tt.args.ints...); got != tt.want {
				t.Errorf("min() = %v, want %v", got, tt.want)
			}
		})
	}
}
