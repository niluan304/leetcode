package main

import (
	"cmp"
)

func maximumRows(matrix [][]int, numSelect int) int {
	m, n := len(matrix), len(matrix[0])
	var ones = make([]int, m)
	for i, row := range matrix {
		ones[i] = Sum(row)
	}

	var ans = 0
	combinations(n, numSelect, func(path []int) {
		ans = max(ans, countRows(matrix, ones, path))
	})
	return ans
}

func combinations(n int, k int, f func(path []int)) {
	var path []int
	var dfs func(start, k int)

	dfs = func(start, k int) {
		if k == 0 {
			f(path)
			return
		}
		for i := start; i <= n-k; i++ {
			path = append(path, i)
			dfs(i+1, k-1)
			path = path[:len(path)-1]
		}
	}
	dfs(0, k)
}

func Sum[S ~[]E, E cmp.Ordered](x S) E {
	var m E
	for i := 0; i < len(x); i++ {
		m += x[i]
	}
	return m
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func ToPtr[T any](v T) *T { return &v }

func maximumRows2(matrix [][]int, numSelect int) int {
	m, n := len(matrix), len(matrix[0])
	var ones = make([]int, m)
	for i, row := range matrix {
		ones[i] = Sum(row)
	}

	var ans = 0
	var path []int
	var dfs func(i int)
	dfs = func(i int) {
		if len(path) == numSelect {
			ans = max(ans, countRows(matrix, ones, path))
			return
		}
		if len(path)+n-i < numSelect {
			return
		}

		for j := i; j < n; j++ {
			path = append(path, j)
			dfs(j + 1)
			path = path[:len(path)-1]
		}
	}

	dfs(0)
	return ans
}

func countRows(matrix [][]int, ones []int, path []int) (count int) {
	for j := 0; j < len(ones); j++ {
		if ones[j] == 0 {
			count++
			continue
		}

		sum := 0
		for _, v := range path {
			sum += matrix[j][v]
		}
		if sum == ones[j] {
			count++
		}
	}
	return count
}
