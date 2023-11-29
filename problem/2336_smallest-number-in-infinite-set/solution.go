package main

import (
	"container/heap"
	"sort"
)

type SmallestInfiniteSet struct {
	set []int
}

// Constructor
// 这不是通用的实现，因为题目保证 1 <= num <= 1000，且调用次数 <= 1000次，才能这样实现。
func Constructor() SmallestInfiniteSet {
	var set = make([]int, 1000)
	for i, _ := range set {
		set[i] = i + 1
	}
	return SmallestInfiniteSet{set: set}
}

func (s *SmallestInfiniteSet) PopSmallest() int {
	n := len(s.set)
	if n == 0 {
		return 0
	}
	ans := s.set[0]
	s.set = s.set[1:]
	return ans
}

func (s *SmallestInfiniteSet) AddBack(num int) {
	i := sort.SearchInts(s.set, num)
	if i < len(s.set) && s.set[i] == num {
		return
	}
	var set = append([]int{}, s.set[:i]...)
	set = append(set, num)
	s.set = append(set, s.set[i:]...)
}

/**
 * Your SmallestInfiniteSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.PopSmallest();
 * obj.AddBack(num);
 */

type SmallestInfiniteSet2 struct {
	heap  IntHeap      // 用 优先队列 + 哈希表 替代有序集合
	exist map[int]bool // 用于 heap 去重

	smallest int
}

func Constructor2() SmallestInfiniteSet2 {
	return SmallestInfiniteSet2{
		heap:     make(IntHeap, 0),
		exist:    make(map[int]bool),
		smallest: 1,
	}
}

func (s *SmallestInfiniteSet2) PopSmallest() int {
	if len(s.heap) == 0 {
		ans := s.smallest
		s.smallest++
		return ans
	}
	ans := heap.Pop(&s.heap).(int)
	s.exist[ans] = false
	return ans
}

func (s *SmallestInfiniteSet2) AddBack(num int) {
	if num < s.smallest && !s.exist[num] {
		s.exist[num] = true
		heap.Push(&s.heap, num)
	}
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
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
