package main

const MOD = 1_000_000_007

// 贪心 + 数学
// - 时间复杂度：O(p)。
// - 空间复杂度：O(1)。
func minNonZeroProduct(p int) int {
	mx := 1<<p - 1
	pow := Pow(mx-1, mx>>1)
	return pow * (mx % MOD) % MOD
}

func Pow(x int, n int) int {
	if n == 0 {
		return 1
	}

	//if n < 0 {
	//	return 1.0 / Pow(x, -n)
	//}

	ans := 1
	if n%2 == 1 {
		ans = x % MOD
	}

	pow := Pow(x, n/2)
	square := (pow % MOD) * (pow % MOD) % MOD
	return ans * square % MOD
}
