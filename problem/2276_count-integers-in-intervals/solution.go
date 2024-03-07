package main

import (
	"sort"

	"github.com/emirpasic/gods/trees/redblacktree"

	. "github.com/niluan304/leetcode/container/segtree"
)

type CountIntervals struct {
	*redblacktree.Tree
	cnt int // 所有区间长度和
}

func Constructor() CountIntervals {
	return CountIntervals{redblacktree.NewWithIntComparator(), 0}
}

func (t *CountIntervals) Add(left, right int) {
	// 遍历所有被 [left,right] 覆盖到的区间（部分覆盖也算）
	for node, _ := t.Ceiling(left); node != nil && node.Value.(int) <= right; node, _ = t.Ceiling(left) {
		l, r := node.Value.(int), node.Key.(int)
		if l < left {
			left = l
		} // 合并后的新区间，其左端点为所有被覆盖的区间的左端点的最小值
		if r > right {
			right = r
		} // 合并后的新区间，其右端点为所有被覆盖的区间的右端点的最大值
		t.cnt -= r - l + 1
		t.Remove(r)
	}
	t.cnt += right - left + 1
	t.Put(right, left) // 所有被覆盖到的区间与 [left,right] 合并成一个新区间
}

func (t *CountIntervals) Count() int { return t.cnt }

/**
 * Your CountIntervals object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(left,right);
 * param_2 := obj.Count();
 */

type CountIntervals3 struct {
	chthollyTree
}

// Constructor3
//
// chthollyTree
// 时间复杂度: O(nlogn)
// 空间复杂度: O(n)
// Deprecated: 超时
func Constructor3() CountIntervals3 {
	return CountIntervals3{
		chthollyTree: chthollyTree{
			tree:  []chthollyNode{{left: 1, right: 1e9 + 1, value: 0}},
			count: 0,
		},
	}
}

func (t *CountIntervals3) Add(left, right int) { t.chthollyTree.Assign(left, right) }
func (t *CountIntervals3) Count() int          { return t.count }

type chthollyNode struct {
	left, right int
	value       int
}

type chthollyTree struct {
	tree  []chthollyNode
	count int
}

func (t *chthollyTree) split(mid int) int {
	i := sort.Search(len(t.tree), func(i int) bool {
		return t.tree[i].right >= mid
	})

	for ; i >= 0; i-- {
		node := t.tree[i]
		if node.left == mid {
			return i
		}

		if node.left < mid { // node.left <= mid - 1
			x := chthollyNode{left: mid, right: node.right, value: node.value}
			tmp := append([]chthollyNode{x}, t.tree[i+1:]...)
			t.tree = append(t.tree[:i+1], tmp...)
			t.tree[i].right = mid - 1
			return i + 1
		}
	}
	return len(t.tree)
}

func (t *chthollyTree) Assign(left, right int) {
	i, j := t.split(left), t.split(right+1) // 顺序不能颠倒，否则可能RE
	for k := i; k < j; k++ {
		node := t.tree[k]
		t.count -= (node.right - node.left + 1) * node.value
	}

	t.tree[i].right = right
	t.tree[i].value = 1
	t.count += right - left + 1

	if i+1 < j {
		t.tree = append(t.tree[:i+1], t.tree[j:]...)
	}
}

type CountIntervals4 struct {
	tree *LazySegmentTree
}

func Constructor4() CountIntervals4 {
	return CountIntervals4{
		tree: NewSumLazySegmentTree(1, 1e9),
	}
}

func (i *CountIntervals4) Add(left, right int) {
	i.tree.Update(left, right, 1)
}

func (i *CountIntervals4) Count() int {
	return i.tree.Query(1, 1e9)
}
