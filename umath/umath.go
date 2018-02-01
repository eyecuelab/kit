//Package umath contains helper functions for math on unsigned integers
package umath

import (
	"github.com/eyecuelab/kit/imath"
)

const (
	is64bit = uint64(^uint(0)) == ^uint64(0)
)

//Min returns the smallest integer argument.
func Min(n uint, a ...uint) uint {
	min := n
	for _, m := range a {
		if m < min {
			min = m
		}
	}
	return min
}

//Sum returns the sum of it's arguments. Sum() is 0
func Sum(a ...uint) uint {
	var sum uint
	for _, u := range a {
		sum += u
	}
	return sum
}

func Range(start, stop uint, step int) []uint {
	if (start > stop && step > 0) ||
		(start < stop && step < 0) || step == 0 {
		return nil
	}
	a := make([]uint, 0, imath.Abs(int(start-stop)/imath.Abs(step)))
	if step < 0 {
		for n := int(start); n > int(stop); n += step {
			a = append(a, uint(n))
		}
		return a
	}

	for n := int(start); n < int(stop); n += step {
		a = append(a, uint(n))
	}

	return a
}

//Product returns the product of it's arguments. Product() is 1.
func Product(a ...uint) uint {
	product := uint(1)
	for _, u := range a {
		product *= u
	}
	return product
}

//Max returns the largest integer argument.
func Max(n uint, a ...uint) uint {
	max := n
	for _, m := range a {
		if m > max {
			max = m
		}
	}
	return max
}

//Clamp takes an uint n, returns low if n < low, high if n > high, and n otherwise.
func Clamp(n, low, high uint) uint {
	if n < low {
		return low
	} else if n > high {
		return high
	}
	return n
}

//Pow is an efficent implementation of exponentiation by squaring.
func Pow(base, exp uint) uint {
	result := uint(1)
	for ; exp > 0; exp >>= 1 {
		if exp&1 > 0 {
			result *= base
		}
		base *= base
	}
	return result
}
