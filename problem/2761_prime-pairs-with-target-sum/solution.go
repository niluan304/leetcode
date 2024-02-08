package main

const mx = 1e6

var notPrime = make([]bool, mx+1)
var primes []int

func init() {
	for i := 2; i <= mx; i++ {
		if notPrime[i] {
			continue
		}
		primes = append(primes, i)
		for j := 2 * i; j <= mx; j += i {
			notPrime[j] = true
		}
	}
}

func findPrimePairs(n int) [][]int {
	var ans [][]int
	for _, x := range primes {
		y := n - x
		if y < x {
			break
		}
		if !notPrime[y] {
			ans = append(ans, []int{x, y})
		}
	}
	return ans
}
