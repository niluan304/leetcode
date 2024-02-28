package main

import (
	"math"
	"slices"
	"sort"
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
	//
	ans := make([]int, m)
	st := NewSegmentTree(left, right)
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
			st.Update(py, px+py)
		}
		ans[i] = st.Query(y, right)
		if ans[i] == 0 { // 1 <= nums1[i], nums2[i] <= 1e9，如果存在符合条件的，值不可能为 0，因此值为 0 代表查询不存在
			ans[i] = -1
		}
	}
	return ans
}

func NewSegmentTree(left, right int) *SegmentTree {
	return &SegmentTree{
		left:  left,
		right: right,
		mx:    0,
		lo:    nil,
		ro:    nil,
	}
}

// SegmentTree
// 动态开点线段树·其一·单点修改
type SegmentTree struct {
	left, right int // 该点的区间范围：[left,right]

	mx int // 根据题目要求，可能求：区间最值 / 区间和

	lo, ro *SegmentTree // 左右儿子节点
}

func (t *SegmentTree) Query(left, right int) int {
	if t == nil || left > t.right || right < t.left {
		return 0 // 本题目要求，不存在为 -1，因此也可以 return -1，相应的，NewSegmentTree 时也需要设置 mx = -1
	}
	if left <= t.left && t.right <= right {
		return t.mx
	}
	return max(t.lo.Query(left, right), t.ro.Query(left, right))
}

// 动态开点
func (t *SegmentTree) spread() {
	mid := (t.left + t.right) / 2
	if t.lo == nil {
		t.lo = NewSegmentTree(t.left, mid)
	}
	if t.ro == nil {
		t.ro = NewSegmentTree(mid+1, t.right)
	}
}

// 单点修改，没有延迟操作（只有区间修改需要延迟操作）
// 用于整合 Add, Update 操作
func (t *SegmentTree) operator(i int, op func(node *SegmentTree)) {
	if t.left == t.right {
		op(t)
		return
	}

	t.spread() // 动态开点

	mid := (t.left + t.right) / 2
	if i <= mid {
		t.lo.operator(i, op)
	} else {
		t.ro.operator(i, op)
	}
	t.mx = max(t.lo.mx, t.ro.mx) // 根据题目要求修改
}

func (t *SegmentTree) Update(i int, value int) {
	t.operator(i, func(node *SegmentTree) {
		node.mx = max(node.mx, value) // 根据题目要求修改
	})
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
