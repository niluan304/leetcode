package main

import (
	"sort"
)

type RangeModule struct {
	tree chthollyTree
}

func Constructor() RangeModule {
	return RangeModule{
		tree: chthollyTree{{left: 1, right: 1e9 + 1, value: 0}},
	}
}

func (m *RangeModule) AddRange(left int, right int) {
	m.tree.Assign(left, right-1, 1)
}

func (m *RangeModule) QueryRange(left int, right int) bool {
	j := sort.Search(len(m.tree), func(i int) bool {
		return m.tree[i].right >= right-1
	})

	for i := j; i >= 0; i-- {
		node := m.tree[i]
		if node.value == 0 {
			return false
		}
		if node.left <= left {
			return true
		}
	}
	return false
}

func (m *RangeModule) RemoveRange(left int, right int) {
	m.tree.Assign(left, right-1, 0)
}

/**
 * Your RangeModule object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddRange(left,right);
 * param_2 := obj.QueryRange(left,right);
 * obj.RemoveRange(left,right);
 */

type chthollyNode struct {
	left, right int
	value       int
}

type chthollyTree []chthollyNode

func (t *chthollyTree) split(mid int) int {
	i := sort.Search(len(*t), func(i int) bool {
		return (*t)[i].right >= mid
	})
	for ; i >= 0; i-- {
		node := (*t)[i]
		if node.left == mid {
			return i
		}
		if node.left < mid { // node.left <= mid - 1
			x := chthollyNode{left: mid, right: node.right, value: node.value}
			tmp := append(chthollyTree{x}, (*t)[i+1:]...)
			*t = append((*t)[:i+1], tmp...)
			(*t)[i].right = mid - 1
			return i + 1
		}
	}
	return len(*t)
}

func (t *chthollyTree) Assign(left, right, value int) {
	i, j := t.split(left), t.split(right+1) // 顺序不能颠倒，否则可能RE
	(*t)[i].right = right
	(*t)[i].value = value
	if i+1 < j {
		*t = append((*t)[:i+1], (*t)[j:]...)
	}
}

type RangeModule2 struct {
	st *LazySegmentTree
}

func Constructor2() RangeModule2 { return RangeModule2{st: &LazySegmentTree{left: 1, right: 1e9}} }

func (m *RangeModule2) AddRange(left int, right int)    { m.st.Update(left, right-1, 1) }
func (m *RangeModule2) RemoveRange(left int, right int) { m.st.Update(left, right-1, 0) }

func (m *RangeModule2) QueryRange(left int, right int) bool {
	x, y := m.st.Query(left, right-1), (right - left)
	return x == y
}

// LazySegmentTree
// 动态开点线段树·其二·延迟标记（仅支持区间覆盖，不支持区间增加）
type LazySegmentTree struct {
	left, right, sum int

	todoUpdate     *int // 延迟标记
	lChild, rChild *LazySegmentTree
}

func (st *LazySegmentTree) Query(left, right int) int {
	// 对于不在线段树中的点，应按照题意来返回
	if st == nil || left > st.right || right < st.left {
		return 0
	}
	if left <= st.left && st.right <= right {
		return st.sum
	}

	st.spread()
	return st.lChild.Query(left, right) + st.rChild.Query(left, right)
}

// 更新本区间结点的值，并将值保存为 lazy 标记，等到需要更新子节点时，再下放 lazy 标记
func (st *LazySegmentTree) doUpdate(value *int) {
	if value == nil {
		return
	}
	st.todoUpdate = value
	st.sum = *value * (st.right - st.left + 1)
}

// 也常叫做 pushdown
// 1. 动态开点
// 2. 把自身的 lazy 标记传递给子节点，并清除自身的 lazy 标记
func (st *LazySegmentTree) spread() {
	mid := (st.left + st.right) / 2
	if st.lChild == nil {
		st.lChild = &LazySegmentTree{left: st.left, right: mid}
	}
	if st.rChild == nil {
		st.rChild = &LazySegmentTree{left: mid + 1, right: st.right}
	}

	st.lChild.doUpdate(st.todoUpdate)
	st.rChild.doUpdate(st.todoUpdate)
	st.todoUpdate = nil
}

func (st *LazySegmentTree) Update(left, right int, value int) {
	// 当前节点已被区间 [left, right] 完整覆盖，不再继续递归
	if left <= st.left && st.right <= right {
		st.doUpdate(&value)
		return
	}

	st.spread()

	mid := (st.left + st.right) / 2
	if left <= mid {
		st.lChild.Update(left, right, value)
	}
	if mid < right {
		st.rChild.Update(left, right, value)
	}
	st.sum = st.lChild.sum + st.rChild.sum
}
