package main

func isPrime(x int) bool {
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

// 暴力穷举
// - 时间复杂度：$O(n \cdot \sqrt{n})$。
// - 空间复杂度：$\mathcal{O}(n)$。
// Deprecated: 超时
func countPrimes(n int) (cnt int) {
	for i := 2; i < n; i++ {
		if isPrime(i) {
			cnt++
		}
	}
	return
}

func countPrimes2(n int) (count int) {
	notPrime := make([]bool, n)
	for i := 2; i < n; i++ {
		if notPrime[i] {
			continue
		}
		count++
		for j := 2 * i; j < n; j += i {
			notPrime[j] = true
		}
	}
	return count
}
