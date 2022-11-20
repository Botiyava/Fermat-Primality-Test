package math

import (
	"math/rand"
)

// PrimaryTest checks if number p is prime.
// If output of the command is true it means that
// p is a prime number with probability of 1/2^k
func PrimaryTest(p, k int) bool {
	// p - given number to check
	// k - number of experiments

	if p%2 == 0 && p != 2 {
		return false
	}
	if p <= 3 {
		return true
	}

	for i := 0; i < k; i++ {
		a := rand.Intn(p-2) + 1
		//fmt.Println(a)
		if calculateResidual(a, p-1, p) != 1 {
			return false
		}
	}
	return true
}
