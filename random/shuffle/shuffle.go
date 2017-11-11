//Package shuffle implements the fisher-yates shuffle to generate a random permutation of the given sequence.
//This uses crypto/rand as the randomness source.
package shuffle

import (
	"crypto/rand"
	"math/big"
)

//Bytes shuffles a slice of bytes
func Bytes(a []byte) ([]byte, error) {
	b := make([]byte, len(a))
	copy(b, a)
	for i := int64(len(b) - 1); i > 0; i-- {
		n, err := rand.Int(rand.Reader, big.NewInt(i+1))
		if err != nil {
			return nil, err
		}
		j := n.Int64()
		b[i], b[j] = b[j], b[i]
	}

	return b, nil
}

//Ints shuffles a slice of ints
func Int64s(a []int64) ([]int64, error) {
	b := make([]int64, len(a))
	copy(b, a)
	for i := int64(len(b) - 1); i > 0; i-- {
		n, err := rand.Int(rand.Reader, big.NewInt(i+1))
		if err != nil {
			return nil, err
		}
		j := n.Int64()
		b[i], b[j] = b[j], b[i]
	}
	return b, nil
}

//Float64s shuffles a slice of strings
func Float64s(a []float64) ([]float64, error) {
	b := make([]float64, len(a))
	copy(b, a)
	for i := int64(len(b) - 1); i > 0; i-- {
		n, err := rand.Int(rand.Reader, big.NewInt(i+1))
		if err != nil {
			return nil, err
		}
		j := n.Int64()
		b[i], b[j] = b[j], b[i]
	}
	return b, nil
}

//Strings shuffles a slice of strings
func Strings(a []string) ([]string, error) {
	b := make([]string, len(a))
	copy(b, a)
	for i := int64(len(b) - 1); i > 0; i-- {
		n, err := rand.Int(rand.Reader, big.NewInt(i+1))
		if err != nil {
			return nil, err
		}
		j := n.Int64()
		b[i], b[j] = b[j], b[i]
	}
	return b, nil
}
