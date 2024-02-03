package main

import "slices"

// LIS
// 常见变形，拼接的条件不一样
//
// - 时间复杂度：$\mathcal{O}(n^2)$。
// - 空间复杂度：$\mathcal{O}(n)$。
func bestTeamScore(scores []int, ages []int) int {
	var n = len(scores)

	type Pair struct{ Score, Age int }
	var paris = make([]Pair, n)
	for i, _ := range paris {
		paris[i] = Pair{Score: scores[i], Age: ages[i]}
	}

	slices.SortFunc(paris[:], func(a, b Pair) int {
		if a.Age == b.Age {
			return a.Score - b.Score
		}
		return a.Age - b.Age
	})

	var ans = 0 // math.MaxInt32 // math.MinInt32
	dp := make([]int, n)
	for i, pair := range paris {
		for j, p := range paris[:i] {
			if p.Score <= pair.Score {
				dp[i] = max(dp[i], dp[j])
			}
		}

		dp[i] += pair.Score
		ans = max(ans, dp[i])
	}

	return ans
}
