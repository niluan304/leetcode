package seg_tree

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
	if st.todoUpdate != nil {
		st.lChild.doUpdate(*st.todoUpdate)
		st.rChild.doUpdate(*st.todoUpdate)
		st.todoUpdate = nil
	}

	// 2. 再处理区间增加的 lazy 标记
	if st.todoAdd != 0 {
		st.lChild.doAdd(st.todoAdd)
		st.rChild.doAdd(st.todoAdd)
		st.todoAdd = 0
	}

}

// 更新本区间结点的值，并将值保存为 lazy 标记，等到需要更新子节点时，再下放 lazy 标记
func (st *LazySegmentTree) doUpdate(value int) {
	st.todoAdd = 0
	st.todoUpdate = &value
	st.sum = value * (st.right - st.left + 1)
}

// 更新本区间结点的值，并将值保存为 lazy 标记，等到需要更新子节点时，再下放 lazy 标记
func (st *LazySegmentTree) doAdd(add int) {
	st.todoAdd += add
	st.sum += add * (st.right - st.left + 1)
}

func (st *LazySegmentTree) operator(left, right int, op func(node *LazySegmentTree)) {
	// 当前节点已被区间 [left, right] 完整覆盖，不再继续递归
	if left <= st.left && st.right <= right {
		op(st)
		return
	}

	st.spread()

	mid := (st.left + st.right) / 2
	if left <= mid {
		st.lChild.operator(left, right, op)
	}
	if mid < right {
		st.rChild.operator(left, right, op)
	}
	st.sum = st.lChild.sum + st.rChild.sum
}

// Update 区间覆盖
// 把区间 [left, right] 的值都修改为 value
func (st *LazySegmentTree) Update(left, right int, value int) {
	st.operator(left, right, func(node *LazySegmentTree) {
		node.doUpdate(value)
	})
}

// Add 区间增加
// 区间 [left, right] 的值都加上 add
func (st *LazySegmentTree) Add(left, right int, add int) {
	st.operator(left, right, func(node *LazySegmentTree) {
		node.doAdd(add)
	})
}
