package math

// CalculateResidual calculates the result of the following equation:
// x^(p-1) = 1 (mod p)
func calculateResidual(x, e, m int) int {
	result := 1
	// x - exponential base
	// e - exponent
	// m - modular operand
	for e > 0 {
		if e%2 == 1 {
			result = (result * x) % m
			e -= 1

		} else {
			x = (x * x) % m
			e /= 2
		}
	}
	return result
}
