package main

import (
	"math"

	. "github.com/niluan304/leetcode/container"
)

// #### 方法一：堆
//
// **思路**
//
// 题目要求 $y_i+y_j+|x_i-x_j|$，其中 $|x_i-x_j| \leq k$ 的最大值。根据题目条件，$i \lt j$ 时，$x_i \leq x_j $，可以拆去绝对值符号，得到 $(-x_i+y_i)+(x_j+y_j)$，其中 $x_j-x_i \leq k$。
//
// 根据这个等式，可以遍历 $\textit{points}$ 所有点，使每个点作为 $[x_j, y_j]$，并且求出满足 $x_j-x_i \leq k$ 的最大的 $(-x_i+y_i)$，而这一步，可以用堆来完成。用一个最小堆，堆的元素是 $[x-y,x]$，堆顶元素的 $(x-y)$ 值最小，即 $(-x+y)$ 值最大。每次遍历一个点时，先弹出堆顶不满足当前 $x_j-x_i \leq k$ 的元素，然后用堆顶
// 元素和当前元素计算 $(-x_i+y_i)+(x_j+y_j)$，再将当前元素放入堆。遍历完后，即得到了式子的最大值。
//
// **复杂度分析**
// - 时间复杂度：$O(n \\times \\log n)$，其中 $n$ 是 $\\textit{points}$ 的长度，每个元素最多进入并离开 $\\textit{heap}$ 一次。
// - 空间复杂度：$O(n)$，是 $\\textit{heap}$ 的空间复杂度。
func leetcode1(points [][]int, k int) int {
	res := math.MinInt

	hp := NewEmptyHeap(func(x, y Item) bool {
		if x.value == y.value {
			return x.index < y.index
		}
		return x.value < y.value
	})

	for _, point := range points {
		x, y := point[0], point[1]
		for hp.Len() > 0 && x-hp.Head().index > k {
			hp.PopHead()
		}
		if hp.Len() > 0 {
			res = max(res, x+y-hp.Head().value)
		}
		hp.Insert(Item{index: x, value: x - y})
	}
	return res
}

type Item struct {
	index int
	value int
}
