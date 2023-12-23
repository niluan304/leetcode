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
