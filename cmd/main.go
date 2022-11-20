package main

import (
	modmath "Fermat_Test/math"
	"fmt"
	"math"
	"os"
)

func main() {
	var p int // Number that we want to test
	var k int // Number of experiments
	fmt.Print("Enter the number p: ")
	fmt.Fscan(os.Stdin, &p)
	fmt.Print("Enter number of experiments: ")
	fmt.Fscan(os.Stdin, &k)

	isPrime := modmath.PrimaryTest(p, k)
	if isPrime {
		fmt.Printf("The number %d is a prime with"+
			" probability of %.2f%%\n", p, (1-(1/math.Pow(2, float64(k))))*100)
	} else {
		fmt.Printf("The number %d is not a prime\n", p)
	}
}
