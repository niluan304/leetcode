package main

func matrixScore(grid [][]int) int {
	var m, n = len(grid), len(grid[0])
	var ones = make([]int, n)
	for _, row := range grid {
		for j := 1; j < n; j++ {
			if row[j] == row[0] {
				ones[j]++
			}
		}
	}

	var ans = 0
	for i, one := range ones {
		ans += max(one, m-one) * 1 << (n - 1 - i)
	}
	return ans
}
