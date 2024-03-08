package main

const MOD = 1_000_000_007

func minimumPossibleSum(n int, k int) int {
	mid := k / 2
	if mid >= n {
		return Sn(1, n, n)
	}

	d := n - mid
	x, y := Sn(1, mid, mid), Sn(k, k+d-1, d)
	return (x + y) % MOD
}

// Sn 等差数列求和
func Sn(a1, an, n int) int {
	return (n * (a1 + an)) / 2 % MOD
}
