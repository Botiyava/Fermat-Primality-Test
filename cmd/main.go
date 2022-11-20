package main

import (
	modmath "Fermat_Test/math"
	"fmt"
	"math"
	"os"
)

func main() {
	var method int // Code of method that will be used
	var p int      // Number that we want to test
	var k int      // Number of experiments
	fmt.Print("Enter the code of method (1) Fermat, (2) Solovay-Strassen: ")
	fmt.Fscan(os.Stdin, &method)
	fmt.Print("Enter the number p: ")
	fmt.Fscan(os.Stdin, &p)
	fmt.Print("Enter number of experiments: ")
	fmt.Fscan(os.Stdin, &k)

	var isPrime bool
	switch method {
	case 1:
		isPrime = modmath.PrimaryTest(p, k)
	case 2:
		isPrime = modmath.SolovayStrassenTest(p, k)
	default:
		fmt.Println("Unknown method")
	}

	if isPrime {
		fmt.Printf("The number %d is a prime with"+
			" probability of %.2f%%\n", p, (1-(1/math.Pow(2, float64(k))))*100)
	} else {
		fmt.Printf("The number %d is not a prime\n", p)
	}
}
