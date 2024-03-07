package segtree

// NewSumSegmentTreeWithNums
// 根据数组初始化区间和线段树，左右边界为：[1, len(nums)]
func NewSumSegmentTreeWithNums(nums []int) *SegmentTree {
	return NewSegmentTreeWithNums(nums, func(lo, ro int) int {
		return lo + ro
	})
}

// NewSegmentTreeWithNums
// 根据数组初始化线段树，左右边界为：[1, len(nums)]
// merge 是根据左右儿子节点，确定本节点的值
func NewSegmentTreeWithNums(nums []int, merge func(lo, ro int) int) *SegmentTree {
	var dfs func(node *SegmentTree)
	dfs = func(node *SegmentTree) {
		left, right := node.left, node.right
		if left == right {
			node.value = nums[left-1]
			return
		}

		mid := (left + right) >> 1 // 允许负数范围，需要 /2 下取整
		node.lo = NewSegmentTree(left, mid, merge)
		node.ro = NewSegmentTree(mid+1, right, merge)

		dfs(node.lo)
		dfs(node.ro)
		node.value = merge(node.lo.value, node.ro.value)
	}

	// 线段树的下标一般从 1 开始，即 [1, len(nums)]
	root := NewSegmentTree(1, len(nums), merge)
	dfs(root)
	return root
}

// NewSumSegmentTree 初始化区间和线段树
func NewSumSegmentTree(left, right int) *SegmentTree {
	return NewSegmentTree(left, right, func(lo, ro int) int {
		return lo + ro
	})
}

// NewSegmentTree 初始化线段树
// merge 是根据左右儿子节点，确定本节点的值
func NewSegmentTree(left, right int, merge func(lo, ro int) int) *SegmentTree {
	return &SegmentTree{
		left:  left,
		right: right,
		value: 0, // 默认值，根据题目要求修改

		lo: nil,
		ro: nil,

		merge: merge,
	}
}

// SegmentTree
// 动态开点线段树·其一·单点修改
type SegmentTree struct {
	left, right int // 该点的区间范围：[left,right]

	value int // 根据题目要求，可能求：区间最值 / 区间和

	lo, ro *SegmentTree // 左右儿子节点

	// 根据左右儿子的值，确定本节点的值
	//
	// 求区间最值时，设置为：func(lo, ro int) int { return max(lo, ro) }
	//
	// 求区间之和时，设置为：func(lo, ro int) int { return lo + ro }
	merge func(lo, ro int) int
}

// Query
// 根据配置的 merge 函数，查询区间 [left, right] 的值
func (t *SegmentTree) Query(left, right int) int {
	if t == nil || left > t.right || right < t.left {
		return 0
	}
	if left <= t.left && t.right <= right {
		return t.value
	}

	lo, ro := t.lo.Query(left, right), t.ro.Query(left, right)
	return t.merge(lo, ro)
}

// 也常叫做 push down
// 动态开点
func (t *SegmentTree) spread() {
	mid := (t.left + t.right) >> 1 // 允许负数范围，需要 /2 下取整
	if t.lo == nil {
		t.lo = NewSegmentTree(t.left, mid, t.merge)
	}
	if t.ro == nil {
		t.ro = NewSegmentTree(mid+1, t.right, t.merge)
	}
}

// 单点修改，没有延迟操作（只有区间修改需要延迟操作）
func (t *SegmentTree) operator(i int, op func(node *SegmentTree)) {
	if t.left == t.right {
		op(t)
		return
	}

	t.spread() // 动态开点

	mid := (t.left + t.right) >> 1 // 允许负数范围，需要 /2 下取整
	if i <= mid {
		t.lo.operator(i, op)
	} else {
		t.ro.operator(i, op)
	}
	t.value = t.merge(t.lo.value, t.ro.value) // 根据题目要求修改
}

// Set 单点修改，修改指定位置的值
//
// 闭包用于返回修改后的节点值
//
// 闭包的参数是节点的修改前的值，节点不存在时为 NewSegmentTree 里的默认值，闭包的返回值为节点修改后的值
func (t *SegmentTree) Set(i int, do func(old int) int) {
	t.operator(i, func(node *SegmentTree) {
		node.value = do(node.value) // 根据题目要求修改
	})
}
