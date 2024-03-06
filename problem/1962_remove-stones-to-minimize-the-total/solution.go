package main

import (
	"cmp"

	. "github.com/niluan304/leetcode/container"
)

func minStoneSum(piles []int, k int) int {
	hp := NewMaxIntHeap(piles...)
	for i := 0; i < k; i++ {
		head := hp.Head()
		hp.Update(0, head-head/2)
	}
	return Sum(hp.Data())
}

func Sum[S ~[]E, E cmp.Ordered](x S) E {
	var m E
	for i := 0; i < len(x); i++ {
		m += x[i]
	}
	return m
}
