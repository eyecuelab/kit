//Package imath is a contains helper functions for integer math.
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

func Map(f func(int) int, a []int) []int {
	mapped := make([]int, len(a))
	for i, n := range a {
		mapped[i] = f(n)
	}
	return mapped
}

func Filter(f func(int) bool, a []int) []int {
	filtered := make([]int, 0, len(a))
	for _, n := range a {
		if f(n) {
			filtered = append(filtered, n)
		}
	}
	return filtered
}

func Reduce(f func(int, int) int, start int, a ...int) int {
	reduced := start
	for _, n := range a {
		reduced = f(reduced, n)
	}
	return reduced
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

func Accumulate(f func(int, int) int, a []int) []int {
	accumulated := make([]int, len(a))
	if len(a) == 0 {
		return accumulated
	}
	accumulated[0] = a[0]
	for i, n := range a[1:] {
		accumulated[i+1] = f(accumulated[i], n)
	}
	return accumulated
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
