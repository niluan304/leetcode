package main

import (
	"math"
)

func stdlib(x float64, n int) float64 {
	return math.Pow(x, float64(n))
}

// 快速幂 递归写法
//
// - 时间复杂度：O(logn)。
// - 空间复杂度：O(logn)。
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		return 1 / myPow(x, -n)
	}

	ans := 1.0
	if n%2 == 1 {
		ans = x
	}

	pow := myPow(x, (n-1)/2)
	return ans * pow * pow
}

// 快速幂 迭代写法
//
// - 时间复杂度：O(logn)。
// - 空间复杂度：O(1)。
func myPow2(x float64, n int) float64 {
	if n < 0 {
		return 1 / myPow2(x, -n)
	}

	ans := 1.0
	// 贡献的初始值为 x
	// 在对 n 进行二进制拆分的同时计算答案
	for n > 0 {
		if n%2 == 1 {
			// 如果 n 二进制表示的最低位为 1，那么需要计入贡献
			ans *= x
		}
		// 将贡献不断地平方
		x = x * x
		// 舍弃 n 二进制表示的最低位，这样我们每次只要判断最低位即可
		n /= 2
	}
	return ans
}
