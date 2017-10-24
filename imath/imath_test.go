package imath

import (
	"math"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMin(t *testing.T) {
	type args struct {
		x int
		a []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"ok",
			args{x: 5, a: []int{2, 6, -1}},
			-1,
		},
		{"default",
			args{x: 5},
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Min(tt.args.x, tt.args.a...); got != tt.want {
				t.Errorf("Min() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMax(t *testing.T) {
	type args struct {
		x int
		a []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"ok",
			args{x: 5, a: []int{2, 6, -1}},
			6,
		},
		{"default",
			args{x: 5},
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Max(tt.args.x, tt.args.a...); got != tt.want {
				t.Errorf("Max() = %v, want %v", got, tt.want)
			}
		})
	}
}

type pair struct {
	a, b int
}

func TestAbs(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		type pair struct {
			a, b int
		}
		pairs := []pair{
			{-1, 1},
			{-0, 0},
			{0, 0},
			{1, 1},
			{2, 2},
			{-2, 2},
			{500, 500},
			{-500, 500},
		}
		for _, p := range pairs {
			if A, B := Abs(p.a), p.b; A != B {
				t.Errorf("Abs(%d) == %d: should be  %d", p.a, A, B)
			}
		}
	})
}

func TestPow(t *testing.T) {

	t.Run("random small pairs", func(t *testing.T) {
		for i := 0; i < 50; i++ {
			base := rand.Intn(2<<20) * RandSign()
			exp := rand.Intn(2<<16) * RandSign()
			if got, want := Pow(base, exp), naivePow(base, exp); got != want {
				t.Errorf("Pow got %v, but want %v", got, want)
			}
		}
	})

}

//TODO - test is broken
/*
func TestPowMod(t *testing.T) {

	t.Run("random small triads", func(t *testing.T) {
		for i := 0; i < 50; i++ {
			base := rand.Intn(2<<20) * RandSign()
			exp := rand.Intn(2 << 16)
			mod := rand.Intn(2<<31) * RandSign()
			if got, want := PowMod(base, exp, mod), naivePowMod(base, exp, mod); got != want {
				t.Errorf("Pow(%d, %d, %d) got %v, but want %v", base, exp, mod, got, want)
			}
		}
	})

}
*/

func TestSign(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(-1, Sign(-5), "sign -5 == -1")
	assert.Equal(0, Sign(0), "sign 0 == 0")
	assert.Equal(1, Sign(math.MaxInt32), "sign(math.MaxInt32) == 1")
}
