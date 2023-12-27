package main

import (
	"sort"

	"github.com/emirpasic/gods/trees/redblacktree"
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

type CountIntervals4 struct{ st *LazySegmentTree }

func Constructor4() CountIntervals4 {
	return CountIntervals4{st: &LazySegmentTree{left: 1, right: 1e9}}
}
func (i *CountIntervals4) Add(left, right int) { i.st.Update(left, right, 1) }
func (i *CountIntervals4) Count() int          { return i.st.sum }

// NewLazySegmentTree
// 根据数组初始化线段树
func NewLazySegmentTree(nums []int) *LazySegmentTree {
	root := &LazySegmentTree{}
	var dfs func(st *LazySegmentTree, left, right int)
	dfs = func(st *LazySegmentTree, left, right int) {
		*st = LazySegmentTree{
			left:   left,
			right:  right,
			lChild: &LazySegmentTree{},
			rChild: &LazySegmentTree{},
		}
		if left == right {
			st.sum = nums[left-1]
			return
		}

		mid := (left + right) / 2
		dfs(st.lChild, left, mid)
		dfs(st.rChild, mid+1, right)
		st.sum = st.lChild.sum + st.rChild.sum
	}

	// 线段树的下标一般从 1 开始，即 [1, len(nums)]
	dfs(root, 1, len(nums))
	return root
}

// LazySegmentTree
// 动态开点线段树·其二·延迟标记（区间修改 + 区间覆盖）
// 同时支持区间修改和区间覆盖的，会比只支持一种的繁琐一些。
type LazySegmentTree struct {
	left, right, sum int

	lChild, rChild *LazySegmentTree

	todoAdd    int  // 区间增加延迟标记
	todoUpdate *int // 区间覆盖延迟标记
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

func (st *LazySegmentTree) Add(left, right int, add int) {
	if add == 0 {
		return
	}

	// 当前节点已被区间 [left, right] 完整覆盖，不再继续递归
	if left <= st.left && st.right <= right {
		st.doAdd(add)
		return
	}

	st.spread()

	mid := (st.left + st.right) / 2
	if left <= mid {
		st.lChild.Add(left, right, add)
	}
	if mid < right {
		st.rChild.Add(left, right, add)
	}
	st.sum = st.lChild.sum + st.rChild.sum
}

// 也常叫做 pushdown
// 1. 动态开点
// 2. 把自身的 lazy 标记传递给子节点，并清除自身的 lazy 标记
func (st *LazySegmentTree) spread() {
	if st.left == st.right {
		return
	}

	mid := (st.left + st.right) / 2
	if st.lChild == nil {
		st.lChild = &LazySegmentTree{left: st.left, right: mid}
	}
	if st.rChild == nil {
		st.rChild = &LazySegmentTree{left: mid + 1, right: st.right}
	}

	// 1. 必须先传递区间覆盖的 lazy 标记给子节点
	st.lChild.doUpdate(st.todoUpdate)
	st.rChild.doUpdate(st.todoUpdate)
	st.todoUpdate = nil

	// 2. 再处理区间增加的 lazy 标记
	st.lChild.doAdd(st.todoAdd)
	st.rChild.doAdd(st.todoAdd)
	st.todoAdd = 0
}

// 更新本区间结点的值，并将值保存为 lazy 标记，等到需要更新子节点时，再下放 lazy 标记
func (st *LazySegmentTree) doUpdate(value *int) {
	if value == nil {
		return
	}

	st.todoAdd = 0
	st.todoUpdate = value
	st.sum = *value * (st.right - st.left + 1)
}

// 更新本区间结点的值，并将值保存为 lazy 标记，等到需要更新子节点时，再下放 lazy 标记
func (st *LazySegmentTree) doAdd(value int) {
	if value == 0 {
		return
	}

	st.todoAdd += value
	st.sum += value * (st.right - st.left + 1)
}
