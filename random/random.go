//Package random provides tools for generating cryptographically secure random elements. it uses golang's built in `crypto/rand` for it's RNG.
package random

import (
	"crypto/rand"
	"encoding/base64"
	"math"
)

//RandomBytes returns a random slice of n bytes
func RandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

//RandomString returns a random string comprised of n bytes
func RandomString(n int) (string, error) {
	b, err := RandomBytes(n)
	return base64.URLEncoding.EncodeToString(b), err
}

//Int64s returns a slice containing n random int64s
func Int64s(n int) ([]int64, error) {
	bytes, err := RandomBytes(8 * n)
	if err != nil {
		return nil, err
	}

	ints := make([]int64, 0, n)
	for i := 0; i < len(bytes); i += 8 {
		var out uint64
		for i, b := range bytes[i : i+8] {
			out |= uint64(b) << uint64(i)
		}
		ints = append(ints, int64(out))
	}

	return ints, nil
}

//Uint64s returns a slice containing n random uint64s
func Uint64s(n int) ([]uint64, error) {
	bytes, err := RandomBytes(8 * n)
	if err != nil {
		return nil, err
	}
	uints := make([]uint64, 0, n)
	for i := 0; i < len(bytes); i += 8 {
		var out uint64
		for i, b := range bytes[i : i+8] {
			out |= uint64(b) << uint64(i)
		}
		uints = append(uints, out)
	}

	return uints, nil
}

//Float64s returns a slice containing n random float64s
func Float64s(n int) ([]float64, error) {
	bytes, err := RandomBytes(8 * n)
	if err != nil {
		return nil, err
	}
	floats := make([]float64, 0, n)
	for i := 0; i < len(bytes); i += 8 {
		var out uint64
		for i, b := range bytes[i : i+8] {
			out |= uint64(b) << uint64(i)
		}
		floats = append(floats, math.Float64frombits(out))
	}
	return floats, nil
}
