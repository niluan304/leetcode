package main

import (
	"container/heap"
	"sort"
)

func pickGifts(gifts []int, k int) int64 {
	var hp IntHeap = gifts
	heap.Init(&hp)

	for i := 0; i < k; i++ {
		hp[0] = mySqrt(hp[0])
		heap.Fix(&hp, 0)
	}

	var total = 0
	for _, gift := range hp {
		total += gift
	}
	return int64(total)
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x any)        {}
func (h *IntHeap) Pop() any          { return nil }

// [69. x 的平方根](https://leetcode.cn/problems/sqrtx/description/)
func mySqrt(x int) int {
	return sort.Search(x+1, func(i int) bool { return i*i > x }) - 1
}
