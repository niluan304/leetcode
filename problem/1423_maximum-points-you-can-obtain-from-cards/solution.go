package main

import "cmp"

func maxScore(cardPoints []int, k int) int {
	var n = len(cardPoints)
	sum := Sum(cardPoints[:n-k])
	mn := sum
	diff := 0

	for i := n - k; i < n; i++ {
		sum += cardPoints[i]
		diff += cardPoints[i-(n-k)]
		mn = min(mn, sum-diff)
	}
	return sum - mn
}

func Sum[S ~[]E, E cmp.Ordered](x S) E {
	var m E
	for i := 0; i < len(x); i++ {
		m += x[i]
	}
	return m
}

func maxScore2(cardPoints []int, k int) int {
	var n = len(cardPoints)
	sum := Sum(cardPoints[:k])
	var ans = sum
	for i := 1; i <= k; i++ {
		sum += cardPoints[n-i] - cardPoints[k-i]
		ans = max(ans, sum)
	}
	return ans
}
