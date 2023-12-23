package main

import (
	"cmp"
	"container/heap"
)

func minStoneSum(piles []int, k int) int {
	hp := IntHeap(piles)
	heap.Init(&hp)
	for i := 0; i < k; i++ {
		hp[0] -= hp[0] / 2
		heap.Fix(&hp, 0)
	}
	return Sum(hp)
}

func Sum[S ~[]E, E cmp.Ordered](x S) E {
	var m E
	for i := 0; i < len(x); i++ {
		m += x[i]
	}
	return m
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
