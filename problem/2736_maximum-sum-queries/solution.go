package main

import (
	"math"
	"slices"
	"sort"

	. "github.com/niluan304/leetcode/container/segtree"
)

// 暴力穷举
// 时间复杂度：O(mn)
// 空间复杂度：O(1)
// Deprecated: Timeout
func bruteForce(nums1 []int, nums2 []int, queries [][]int) []int {
	n := len(nums1)
	ans := make([]int, len(queries))
	for i := range ans {
		ans[i] = -1
		x, y := queries[i][0], queries[i][1]
		for j := 0; j < n; j++ {
			if nums1[j] >= x && nums2[j] >= y {
				ans[i] = max(ans[i], nums1[j]+nums2[j])
			}
		}
	}
	return ans
}

// 排序 + 动态开点线段树
// - 时间复杂度：$\mathcal{O}(n\log n)$。
// - 空间复杂度：$\mathcal{O}(n)$。
func maximumSumQueries(nums1 []int, nums2 []int, queries [][]int) []int {
	n, m := len(nums1), len(queries)
	left, right := math.MaxInt32, math.MinInt32 // 也可以直接设置为 [1, 1e9]

	type Pair struct{ x, y int }
	pairs := make([]Pair, n)
	for i, x := range nums1 {
		y := nums2[i]
		pairs[i] = Pair{x: x, y: y}

		left = min(left, y)
		right = max(right, y)
	}
	slices.SortFunc(pairs, func(a, b Pair) int {
		return b.x - a.x // 降序排列
	})

	for i := range queries {
		queries[i] = append(queries[i], i) // 原地增加索引标记
	}
	slices.SortFunc(queries, func(a, b []int) int {
		return b[0] - a[0] // 降序排列
	})

	ans := make([]int, m)
	tree := NewSegmentTree(left, right, func(lo, ro int) int {
		return max(lo, ro) // 求区间最大值
	})
	j := 0
	for _, query := range queries {
		x, y, i := query[0], query[1], query[2]
		for ; j < n; j++ {
			px, py := pairs[j].x, pairs[j].y
			if px < x {
				break
			}

			// 设线段树的区间节点为 py, 值为 px + py
			//
			// 通过分批插入「实现降维」，或者说「避免数据污染」
			// 先插入 px >= x 的节点，那线段树里的所有节点都满足 node.x >= x
			// 然后查询区间 [y, 1e9] 的最大值
			tree.Set(py, func(old int) int {
				return max(old, px+py) // 更新为更最大值
			})
		}
		ans[i] = tree.Query(y, right)
		if ans[i] == 0 { // 1 <= nums1[i], nums2[i] <= 1e9，如果存在符合条件的，值不可能为 0，因此值为 0 代表查询不存在
			ans[i] = -1
		}
	}
	return ans
}

// 排序 + 单调栈
// - 时间复杂度：$\mathcal{O}(n\log n)$。
// - 空间复杂度：$\mathcal{O}(n)$。
func maximumSumQueries2(nums1 []int, nums2 []int, queries [][]int) []int {
	n := len(nums1)

	type Pair struct{ x, y int }
	pairs := make([]Pair, n)
	for i, x := range nums1 {
		y := nums2[i]
		pairs[i] = Pair{x: x, y: y}
	}
	slices.SortFunc(pairs, func(a, b Pair) int {
		return b.x - a.x // 降序排列
	})

	for i := range queries {
		queries[i] = append(queries[i], i) // 原地增加索引标记
	}

	slices.SortFunc(queries, func(a, b []int) int {
		x0, x1 := a[0], b[0]
		return x1 - x0
	})

	ans := make([]int, len(queries))
	j := 0

	type Item struct{ y, sum int }
	var stack []Item
	for _, query := range queries {
		for ; j < n; j++ {
			x, y := pairs[j].x, pairs[j].y
			if x < query[0] {
				break
			}

			sum := x + y
			m := len(stack)
			for ; m > 0; m-- {
				if stack[m-1].sum > sum {
					break
				}
				stack = stack[:m-1]
			}

			// 由于从大到小遍历 x，x 是逐渐变小的，那么 top.x >= pairs[j].x
			// 而此时有 top.y < pairs[j].y, 那么说明 top.sum > pairs[j].sum
			if m == 0 || stack[m-1].y < y {
				stack = append(stack, Item{
					y:   y,
					sum: sum,
				})
			}
		}

		k := sort.Search(len(stack), func(i int) bool {
			return stack[i].y >= query[1]
		})
		i := query[2]
		if k < len(stack) {
			ans[i] = stack[k].sum
		} else {
			ans[i] = -1
		}
	}
	return ans
}
