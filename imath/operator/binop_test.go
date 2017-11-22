package operator

import (
	"testing"

	"github.com/eyecuelab/kit/random"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	A, _ := random.Int64s(25)
	B, _ := random.Int64s(25)
	for i := range A {
		a, b := int(A[i]), int(B[i])
		assert.Equal(t, a+b, Add(a, b))
	}
}

func TestMul(t *testing.T) {
	A, _ := random.Int64s(25)
	B, _ := random.Int64s(25)
	for i := range A {
		a, b := int(A[i]), int(B[i])
		assert.Equal(t, a*b, Mul(a, b))
	}
}

func TestDiv(t *testing.T) {
	A, _ := random.Int64s(25)
	B, _ := random.Int64s(25)
	for i := range A {
		a, b := int(A[i]), int(B[i])
		assert.Equal(t, a/b, Div(a, b))
	}
}

func TestMod(t *testing.T) {
	A, _ := random.Int64s(25)
	B, _ := random.Int64s(25)
	for i := range A {
		a, b := int(A[i]), int(B[i])
		assert.Equal(t, a%b, Mod(a, b))
	}
}

func TestSub(t *testing.T) {
	A, _ := random.Int64s(25)
	B, _ := random.Int64s(25)
	for i := range A {
		a, b := int(A[i]), int(B[i])
		assert.Equal(t, a-b, Sub(a, b))
	}
}

func TestLT(t *testing.T) {
	A, _ := random.Int64s(25)
	B, _ := random.Int64s(25)
	for i := range A {
		a, b := int(A[i]), int(B[i])
		assert.Equal(t, a < b, LT(a, b))
	}
}

func TestLTE(t *testing.T) {
	A, _ := random.Int64s(25)
	B, _ := random.Int64s(25)
	for i := range A {
		a, b := int(A[i]), int(B[i])
		assert.Equal(t, a <= b, LTE(a, b))
	}
}

func TestGT(t *testing.T) {
	A, _ := random.Int64s(25)
	B, _ := random.Int64s(25)
	for i := range A {
		a, b := int(A[i]), int(B[i])
		assert.Equal(t, a > b, GT(a, b))
	}
}

func TestGTE(t *testing.T) {
	A, _ := random.Int64s(25)
	B, _ := random.Int64s(25)
	for i := range A {
		a, b := int(A[i]), int(B[i])
		assert.Equal(t, a >= b, GTE(a, b))
	}
}
