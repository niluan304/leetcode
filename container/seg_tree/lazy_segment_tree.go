package seg_tree

// NewLazySegmentTreeWithNums
// 根据数组初始化线段树
func NewLazySegmentTreeWithNums(nums []int) *LazySegmentTree {
	root := &LazySegmentTree{}
	var dfs func(node *LazySegmentTree, left, right int)
	dfs = func(node *LazySegmentTree, left, right int) {
		*node = LazySegmentTree{
			left:  left,
			right: right,
			lo:    &LazySegmentTree{},
			ro:    &LazySegmentTree{},
		}
		if left == right {
			node.sum = nums[left-1]
			return
		}

		mid := (left + right) / 2
		dfs(node.lo, left, mid)
		dfs(node.ro, mid+1, right)
		node.sum = node.lo.sum + node.ro.sum
	}

	// 线段树的下标一般从 1 开始，即 [1, len(nums)]
	dfs(root, 1, len(nums))
	return root
}

// NewLazySegmentTree 初始化线段树节点
func NewLazySegmentTree(left, right int) *LazySegmentTree {
	return &LazySegmentTree{
		left:  left,
		right: right,

		sum: 0, // 根据题目要求修改

		lo: nil,
		ro: nil,

		todoAdd:    0,
		todoUpdate: nil,
	}
}

// LazySegmentTree
// 动态开点线段树·其二·延迟标记（区间修改 + 区间覆盖）
// 同时支持区间修改和区间覆盖的，会比只支持一种的繁琐一些。
type LazySegmentTree struct {
	left, right int // 该点的区间范围：[left,right]

	sum int // 根据题目要求，可能求：区间最值 / 区间和

	lo, ro *LazySegmentTree // 左右儿子节点

	todoAdd    int  // 区间增加延迟标记
	todoUpdate *int // 区间覆盖延迟标记
}

func (t *LazySegmentTree) Query(left, right int) int {
	// 对于不在线段树中的点，应按照题意来返回
	if t == nil || left > t.right || right < t.left {
		return 0
	}
	if left <= t.left && t.right <= right {
		return t.sum
	}

	t.spread()
	return t.lo.Query(left, right) + t.ro.Query(left, right)
}

// 也常叫做 push down
// 1. 动态开点
// 2. 把自身的 lazy 标记传递给子节点，并清除自身的 lazy 标记
func (t *LazySegmentTree) spread() {
	if t.left == t.right {
		return
	}

	mid := (t.left + t.right) / 2
	if t.lo == nil {
		t.lo = NewLazySegmentTree(t.left, mid)
	}
	if t.ro == nil {
		t.ro = NewLazySegmentTree(mid+1, t.right)
	}

	// 1. 必须先传递区间覆盖的 lazy 标记给子节点
	if t.todoUpdate != nil {
		t.lo.doUpdate(*t.todoUpdate)
		t.ro.doUpdate(*t.todoUpdate)
		t.todoUpdate = nil
	}

	// 2. 再处理区间增加的 lazy 标记
	if t.todoAdd != 0 {
		t.lo.doAdd(t.todoAdd)
		t.ro.doAdd(t.todoAdd)
		t.todoAdd = 0
	}
}

// 更新本区间结点的值，并将值保存为 lazy 标记，等到需要更新子节点时，再下放 lazy 标记
func (t *LazySegmentTree) doUpdate(value int) {
	t.todoAdd = 0
	t.todoUpdate = &value
	t.sum = value * (t.right - t.left + 1)

	// t.mx = max(t.mx, value) // 更新最值
}

// 更新本区间结点的值，并将值保存为 lazy 标记，等到需要更新子节点时，再下放 lazy 标记
func (t *LazySegmentTree) doAdd(add int) {
	t.todoAdd += add
	t.sum += add * (t.right - t.left + 1)
}

func (t *LazySegmentTree) operator(left, right int, op func(node *LazySegmentTree)) {
	// 当前节点已被区间 [left, right] 完整覆盖，不再继续递归
	if left <= t.left && t.right <= right {
		op(t)
		return
	}

	t.spread()

	mid := (t.left + t.right) / 2
	if left <= mid {
		t.lo.operator(left, right, op)
	}
	if mid < right {
		t.ro.operator(left, right, op)
	}
	t.sum = t.lo.sum + t.ro.sum // 根据题目要求修改
}

// Update 区间覆盖
// 把区间 [left, right] 的值都修改为 value
func (t *LazySegmentTree) Update(left, right int, value int) {
	t.operator(left, right, func(node *LazySegmentTree) {
		node.doUpdate(value)
	})
}

// Add 区间增加
// 区间 [left, right] 的值都加上 add
func (t *LazySegmentTree) Add(left, right int, add int) {
	t.operator(left, right, func(node *LazySegmentTree) {
		node.doAdd(add)
	})
}
