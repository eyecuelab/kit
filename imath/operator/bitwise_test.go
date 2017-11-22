package operator

import (
	"testing"

	"github.com/eyecuelab/kit/random"
	"github.com/stretchr/testify/assert"
)

func TestBitAnd(t *testing.T) {
	A, _ := random.Int64s(25)
	B, _ := random.Int64s(25)
	for i := range A {
		a, b := int(A[i]), int(B[i])
		assert.Equal(t, int(uint(a)&uint(b)), BitAnd(a, b))
	}
}

func TestBitOr(t *testing.T) {
	A, _ := random.Int64s(25)
	B, _ := random.Int64s(25)
	for i := range A {
		a, b := int(A[i]), int(B[i])
		assert.Equal(t, int(uint(a)|uint(b)), BitOr(a, b))
	}
}

func TestBitXor(t *testing.T) {
	A, _ := random.Int64s(25)
	B, _ := random.Int64s(25)
	for i := range A {
		a, b := int(A[i]), int(B[i])
		assert.Equal(t, int(uint(a)^uint(b)), BitXor(a, b))
	}
}

func TestBitInvert(t *testing.T) {
	A, _ := random.Int64s(25)
	for i := range A {
		a := int(A[i])
		assert.Equal(t, int(^uint(a)), BitInvert(a))
	}
}
