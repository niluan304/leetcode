package main

import (
	"slices"
)

// 排序 + 分组循环
//
// 本题用到一个经典结论。
// 把序列中的每个元素看成图中的一个点，如果两个元素可以相互交换，则在两个元素之间连一条边。
// 结论：处于同一连通块内的所有元素可以通过若干次操作排成任意顺序。
//
// - 时间复杂度：$\mathcal{O}(n \log n)$。
// - 空间复杂度：$\mathcal{O}(n)$。
func lexicographicallySmallestArray(nums []int, limit int) []int {

	var n = len(nums)
	var ans = make([]int, n)

	type Pair struct{ Index, Value int }
	var pairs = make([]Pair, 0, n)
	for i, num := range nums {
		pairs = append(pairs, Pair{Index: i, Value: num})
	}
	slices.SortFunc(pairs, func(a, b Pair) int {
		return a.Value - b.Value
	})

	// 分组循环
	ids := slices.Clone(pairs)
	for i := 0; i < n; {
		j := i
		for i++; i < n; i++ {
			if pairs[i].Value-pairs[i-1].Value > limit {
				break
			}
		}
		// 说明区间 [j : i-1] 内的元素可以排列成任意顺序

		slices.SortFunc(ids[j:i], func(a, b Pair) int {
			return a.Index - b.Index
		})
		for k := j; k < i; k++ {
			ans[ids[k].Index] = pairs[k].Value
		}
	}

	return ans
}
