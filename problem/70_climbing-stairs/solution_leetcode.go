package main

import (
	"math"
)

type matrix [2][2]int

func mul(a, b matrix) (c matrix) {
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			c[i][j] = a[i][0]*b[0][j] + a[i][1]*b[1][j]
		}
	}
	return c
}

func pow(a matrix, n int) matrix {
	res := matrix{{1, 0}, {0, 1}}
	for ; n > 0; n >>= 1 {
		if n&1 == 1 {
			res = mul(res, a)
		}
		a = mul(a, a)
	}
	return res
}

// leetcode 2 矩阵快速幂
func leetcode2(n int) int {
	res := pow(matrix{{1, 1}, {1, 0}}, n)
	return res[0][0]
}

// leetcode 3 通项公式
func leetcode3(n int) int {
	sqrt5 := math.Sqrt(5)
	pow1 := math.Pow((1+sqrt5)/2, float64(n+1))
	pow2 := math.Pow((1-sqrt5)/2, float64(n+1))
	return int(math.Round((pow1 - pow2) / sqrt5))
}
