package math

func calculateJacobySymbol(a, p int) int {
	// a/n
	// a - random number
	// n - number to test

	if p == 0 || p%2 == 0 {
		return 0
	}
	a = a % p
	result := 1
	var reminder int
	for a != 0 {
		for a%2 == 0 {
			a >>= 1
			reminder = p % 8
			if reminder == 3 || reminder == 5 {
				result = -result
			}
		}
		reminder = p
		p = a
		a = reminder
		if a%4 == 3 && p%4 == 3 {
			result = -result
		}
		a = a % p
	}
	if p == 1 {
		return result
	} else {
		return 0
	}
}
