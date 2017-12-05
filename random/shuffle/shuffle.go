//Package shuffle implements the fisher-yates shuffle to generate a random permutation of the given sequence.
//This uses crypto/rand as the randomness source.
package shuffle

import (
	"crypto/rand"
	"math/big"

	"github.com/eyecuelab/kit/copyslice"
)

//rint returns an int64 in the half-open interval [0, n)
func rint(n int) (int64, error) {
	big, err := rand.Int(rand.Reader, big.NewInt(int64(n)))
	if err != nil {
		return 0, err
	}
	return big.Int64(), nil
}

//Bytes shuffles a slice of bytes
func Bytes(a []byte) ([]byte, error) {
	b := copyslice.Byte(a)
	for i := len(b) - 1; i > 0; i-- {
		j, err := rint(i + 1)
		if err != nil {
			return nil, err
		}
		b[i], b[j] = b[j], b[i]
	}

	return b, nil
}

//Int64s shuffles a slice of int64s
func Int64s(a []int64) ([]int64, error) {
	b := copyslice.Int64(a)
	for i := len(b) - 1; i > 0; i-- {
		j, err := rint(i + 1)
		if err != nil {
			return nil, err
		}
		b[i], b[j] = b[j], b[i]
	}
	return b, nil
}

//Float64s shuffles a slice of strings
func Float64s(a []float64) ([]float64, error) {
	b := copyslice.Float64(a)
	copy(b, a)
	for i := (len(b) - 1); i > 0; i-- {
		j, err := rint(i + 1)
		if err != nil {
			return nil, err
		}
		b[i], b[j] = b[j], b[i]
	}
	return b, nil
}

//Strings shuffles a slice of strings
func Strings(a []string) ([]string, error) {
	b := make([]string, len(a))
	copy(b, a)
	for i := (len(b) - 1); i > 0; i-- {
		j, err := rint(i + 1)
		if err != nil {
			return nil, err
		}
		b[i], b[j] = b[j], b[i]
	}
	return b, nil
}
