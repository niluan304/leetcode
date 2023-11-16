package main

import "sort"

// dp 动态规划, LIS
// 时间复杂度：O(n^2)
// 空间复杂度：O(n)
func bestTeamScore(scores []int, ages []int) int {
	type Pair struct {
		age   int
		score int
	}

	var n = len(scores)
	var pairs = make([]Pair, n)
	for i := 0; i < n; i++ {
		pairs[i] = Pair{
			age:   ages[i],
			score: scores[i],
		}
	}

	sort.Slice(pairs, func(i, j int) bool {
		x, y := pairs[i], pairs[j]
		if x.age == y.age {
			return x.score < y.score
		}
		return x.age < y.age
	})
	var dp = make([]int, n)
	var ans = 0
	for i, pair := range pairs {
		var score = 0
		for j, p := range pairs[:i] {
			if p.score > pair.score {
				continue
			}
			score = _max(score, dp[j])
		}
		dp[i] = score + pair.score
		ans = _max(ans, dp[i])
	}
	return ans
}

func _max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func _min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func _abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
