package seg_tree

// NewSegmentTreeWithNums
// 根据数组初始化线段树
func NewSegmentTreeWithNums(nums []int) *SegmentTree {
	root := &SegmentTree{}
	var dfs func(node *SegmentTree, left, right int)
	dfs = func(node *SegmentTree, left, right int) {
		*node = SegmentTree{
			left:  left,
			right: right,
			lo:    &SegmentTree{},
			ro:    &SegmentTree{},
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

// NewSegmentTree 初始化线段树
func NewSegmentTree(left, right int) *SegmentTree {
	return &SegmentTree{
		left:  left,
		right: right,

		sum: 0, // 根据题目要求修改

		lo: nil,
		ro: nil,
	}
}

// SegmentTree
// 动态开点线段树·其一·单点修改
type SegmentTree struct {
	left, right int // 该点的区间范围：[left,right]

	sum int // 根据题目要求，可能求：区间最值 / 区间和

	lo, ro *SegmentTree // 左右儿子节点
}

func (t *SegmentTree) Query(left, right int) int {
	if t == nil || left > t.right || right < t.left {
		return 0
	}
	if left <= t.left && t.right <= right {
		return t.sum
	}
	return t.lo.Query(left, right) + t.ro.Query(left, right)
}

// 也常叫做 push down
// 动态开点
func (t *SegmentTree) spread() {
	mid := (t.left + t.right) / 2
	if t.lo == nil {
		t.lo = NewSegmentTree(t.left, mid)
	}
	if t.ro == nil {
		t.ro = NewSegmentTree(mid+1, t.right)
	}
}

// 单点修改，没有延迟操作（只有区间修改需要延迟操作）
// 用于整合 Add, Update 操作
func (t *SegmentTree) operator(i int, op func(node *SegmentTree)) {
	if t.left == t.right {
		op(t)
		return
	}

	t.spread() // 动态开点

	mid := (t.left + t.right) / 2
	if i <= mid {
		t.lo.operator(i, op)
	} else {
		t.ro.operator(i, op)
	}
	t.sum = t.lo.sum + t.ro.sum // 根据题目要求修改
}

func (t *SegmentTree) Update(i int, value int) {
	t.operator(i, func(node *SegmentTree) {
		node.sum = value // 根据题目要求修改
		// node.mx = max(node.mx, value) // 最值
	})
}

func (t *SegmentTree) Add(i int, add int) {
	t.operator(i, func(node *SegmentTree) {
		node.sum += add
	})
}
