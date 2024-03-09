package main

import (
	. "github.com/niluan304/leetcode/container/segtree"
	. "github.com/niluan304/leetcode/copypasta"
)

// 两数之和 + 线段树
// - 时间复杂度：O(nlogC)。C = 1e4 - (-1e4)
// - 空间复杂度：O(n)。
func countSmaller(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)

	tree := NewSumSegmentTree(-1e4, 1e4)
	for i := n - 1; i >= 0; i-- {
		num := nums[i]
		ans[i] = tree.Query(-1e4, num-1)

		tree.Set(num, func(old int) int {
			return old + 1
		})
	}

	return ans
}

// 带索引排序的归并排序
//
// - 时间复杂度：O(nlogn)。
// - 空间复杂度：O(n)。
func countSmaller2(nums []int) []int {
	ans := make([]int, len(nums))
	// 解法：只对数组下标排序
	type Pair struct{ Num, Index int }
	paris := make([]Pair, len(nums))
	for i, num := range nums {
		paris[i] = Pair{Num: num, Index: i}
	}

	MergeSortWith(paris, func(x, y Pair) int {
		return x.Num - y.Num
	}, func(left, right []Pair) {
		j := 0
		for _, pair := range left {
			for j < len(right) && pair.Num > right[j].Num {
				j++
			}
			ans[pair.Index] += j
		}
	})

	return ans
}
