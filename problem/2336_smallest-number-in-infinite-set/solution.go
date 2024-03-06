package main

import (
	"sort"

	. "github.com/niluan304/leetcode/container"
)

type SmallestInfiniteSet struct {
	set []int
}

// Constructor
// 这不是通用的实现，因为题目保证 1 <= num <= 1000，且调用次数 <= 1000次，才能这样实现。
func Constructor() SmallestInfiniteSet {
	set := make([]int, 1000)
	for i := range set {
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
	set := append([]int{}, s.set[:i]...)
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
	heap  *Heap[int]   // 用 小根堆 + 哈希表 替代有序集合
	exist map[int]bool // 用于 heap 去重

	smallest int
}

func Constructor2() SmallestInfiniteSet2 {
	return SmallestInfiniteSet2{
		heap:     NewMinIntHeap(),
		exist:    make(map[int]bool),
		smallest: 1,
	}
}

func (s *SmallestInfiniteSet2) PopSmallest() int {
	if s.heap.Len() == 0 {
		ans := s.smallest
		s.smallest++
		return ans
	}
	ans := s.heap.PopHead()
	s.exist[ans] = false
	return ans
}

func (s *SmallestInfiniteSet2) AddBack(num int) {
	if num < s.smallest && !s.exist[num] {
		s.exist[num] = true
		s.heap.Insert(num)
	}
}
