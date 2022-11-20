package math

import "math/rand"

func SolovayStrassenTest(p, k int) bool {
	// p - given number to check
	// k - number of experiments

	if (p%2 == 0 && p != 2) || p < 1 {
		return false
	}
	if p <= 3 {
		return true
	}

	for i := 0; i < k; i++ {
		a := rand.Intn(p-2) + 1
		jacobySymbol := (p + calculateJacobySymbol(a, p)) % p
		mod := calculateResidual(a, (p-1)/2, p)
		if jacobySymbol == 0 || mod != jacobySymbol {
			return false
		}
	}
	return true
}
