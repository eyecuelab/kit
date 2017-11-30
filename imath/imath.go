//Package imath contains tools for signed integer math. It largely corresponds with go's built in `math` library for float64s
package imath

import "math/rand"
import "github.com/eyecuelab/kit/imath/operator"

const (
	is64bit = uint64(^uint(0)) == ^uint64(0)
)

//Min returns the smallest integer argument.
func Min(n int, a ...int) int {
	min := n
	for _, m := range a {
		if m < min {
			min = m
		}
	}
	return min
}

//Sum returns the sum of it's arguments. Sum() is 0
func Sum(a ...int) int {
	return Reduce(operator.Add, 0, a...)
}

func Range(start, stop, step int) []int {
	if (start > stop && step > 0) ||
		(start < stop && step < 0) || step == 0 {
		return nil
	}
	a := make([]int, 0, Abs(start-stop)/Abs(step))
	if step < 0 {
		for n := start; n > stop; n += step {
			a = append(a, n)
		}
		return a
	}

	for n := start; n < stop; n += step {
		a = append(a, n)
	}

	return a
}

//Product returns the product of it's arguments. Product() is 1.
func Product(a ...int) int {
	return Reduce(operator.Mul, 1, a...)
}

//Max returns the largest integer argument.
func Max(n int, a ...int) int {
	max := n
	for _, m := range a {
		if m > max {
			max = m
		}
	}
	return max
}

//Abs returns the absolute value of n
func Abs(n int) int {
	if n >= 0 {
		return n
	}
	return -n
}

//RandSign returns -1 or 1 at random, using the default Source of math/rand This is NOT crypto-safe, at all.
func RandSign() int {
	if rand.Intn(2) > 0 {
		return 1
	}
	return -1
}

//Clamp takes an int n, returns low if n < low, high if n > high, and n otherwise.
func Clamp(n, low, high int) int {
	if n < low {
		return low
	} else if n > high {
		return high
	}
	return n
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
