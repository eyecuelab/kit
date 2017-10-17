//Package imath is a contains helper functions for integer math.
package imath

import "math/rand"

const (
	is64bit = uint64(^uint(0)) == ^uint64(0)
)

func Min(x int, a ...int) int {
	min := x
	for _, n := range a {
		if n < min {
			min = n
		}
	}
	return min
}

func Max(x int, a ...int) int {
	max := x
	for _, n := range a {
		if n > max {
			max = n
		}
	}
	return max
}

func Abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

//RandSign returns -1 or 1 at random, using the default Source of math/rand This is NOT crypto-safe, at all.
func RandSign() int {
	if rand.Intn(2) > 0 {
		return 1
	}
	return -1
}

//Sign returns the sign of the operand. The sign of zero is zero.
func Sign(n int) int {
	switch {
	case n < 0:
		return -1

	case n == 0:
		return 0

	default:
		return 1
	}
}

//Pow is an efficent implementation of exponentiation by squaring.
func Pow(base, exp int) int {
	result := 1
	for ; exp > 0; exp >>= 1 {
		if exp&1 > 0 {
			result *= base
		}
		base *= base
	}
	return result
}

func naivePow(base, exp int) int {
	result := 1
	for ; exp > 0; exp-- {
		result *= base
	}
	return result
}

//TODO - test is broken
/*
func PowMod(base, exp, mod int) int {
	base %= mod
	exp %= mod

	result := 1
	for ; exp > 0; exp >>= 1 {
		if exp&1 > 0 {
			result = result * base % mod
		}
		base = base * base % mod

	}
	return result
}


func PowModSafe(base, exp, mod int) (int, bool) {
	base %= mod
	if exp < 0 {
		return 0, false // negative exponent is unclear
	}
	exp %= mod

	if is64bit && (math.MaxInt64/base) < base {
		return 0, false
	} else if math.MaxInt32/base < base {
		return 0, false
	}
	// cannot exponentiate base because it will overflow int
	result := 1
	for ; exp > 0; exp >>= 1 {
		if exp&1 > 0 {
			result = result * base % mod
		}
		base = base * base % mod
	}
	return result, true
}
*/
