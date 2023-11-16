package main

import "math"

// leetcode1 袖珍计算器算法
// 袖珍计算器算法是一种用指数函数 exp 和对数函数 ln 代替平方根函数的方法。
// 我们通过有限的可以使用的数学函数，得到我们想要计算的结果。
func leetcode1(x int) int {
	if x == 0 {
		return 0
	}
	ans := int(math.Exp(0.5 * math.Log(float64(x))))
	if (ans+1)*(ans+1) <= x {
		return ans + 1
	}
	return ans
}

// leetcode2 牛顿迭代
// 牛顿迭代法是一种可以用来快速求解函数零点的方法。
func leetcode2(x int) int {
	if x == 0 {
		return 0
	}
	C, x0 := float64(x), float64(x)
	for {
		xi := 0.5 * (x0 + C/x0)
		if math.Abs(x0-xi) < 1e-7 {
			break
		}
		x0 = xi
	}
	return int(x0)
}

// minMemory 1.91 MB 的代码示例
func minMemory(x int) int {
	V := float64(x)
	var root float64 = V
	for math.Abs(root*root-V) > 0.000001 {
		root = (root + V/root) / 2
	}
	return int(root)
}
