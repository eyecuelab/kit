package operator

//BitAnd is bitwise and, '&'.
func BitAnd(a, b int) int { return int(uint(a) & uint(b)) }

//BitOr is bitwise or, '|'
func BitOr(a, b int) int { return int(uint(a) | uint(b)) }

//BitXor is bitwise XOR, '^'
func BitXor(a, b int) int { return int(uint(a) ^ uint(b)) }

//BitInvert inverts the bits of n. Go treats integers as two's complement, regardless of the underlying hardware.
func BitInvert(n int) int { return int(^uint(n)) }
