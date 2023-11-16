package main

import (
	"math"
)

func stdlib(x float64, n int) float64 {
	return math.Pow(x, float64(n))
}

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		return 1 / myPow(x, -n)
	}

	if n%2 == 0 {
		return Square(myPow(x, n/2))
	}

	return Square(myPow(x, (n-1)/2)) * x
}

func Square(x float64) float64 { return x * x }
