package operator

func BuildClamp(low, high int) func(int) int {
	return func(n int) int {
		if n < low {
			return low
		} else if n > high {
			return high
		}
		return n
	}
}

func BuildLowClamp(low int) func(int) int {
	return func(n int) int {
		if n < low {
			return low
		}
		return n
	}
}

func BuildHighClamp(high int) func(int) int {
	return func(n int) int {
		if n > high {
			return high
		}
		return n
	}
}
