package seg_tree

// NewSegmentTree
// 根据数组初始化线段树
func NewSegmentTree(nums []int) *SegmentTree {
	root := &SegmentTree{}
	var dfs func(st *SegmentTree, left, right int)
	dfs = func(st *SegmentTree, left, right int) {
		*st = SegmentTree{
			left:   left,
			right:  right,
			lChild: &SegmentTree{},
			rChild: &SegmentTree{},
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

// SegmentTree
// 动态开点线段树·其一·单点修改
type SegmentTree struct {
	left, right, sum int

	lChild, rChild *SegmentTree
}

func (st *SegmentTree) Query(left, right int) int {
	if st == nil || left > st.right || right < st.left {
		return 0
	}
	if left <= st.left && st.right <= right {
		return st.sum
	}
	return st.lChild.Query(left, right) + st.rChild.Query(left, right)
}

// 动态开点
func (st *SegmentTree) spread() {
	mid := (st.left + st.right) / 2
	if st.lChild == nil {
		st.lChild = &SegmentTree{left: st.left, right: mid}
	}
	if st.rChild == nil {
		st.rChild = &SegmentTree{left: mid + 1, right: st.right}
	}
}

// 单点修改，没有延迟操作（只有区间修改需要延迟操作），因此可以整合 Add, Update 操作
func (st *SegmentTree) operator(i int, op func(st *SegmentTree)) {
	if st.left == st.right {
		op(st)
		return
	}

	st.spread() // 动态开点

	mid := (st.left + st.right) / 2
	if i <= mid {
		st.lChild.operator(i, op)
	} else {
		st.rChild.operator(i, op)
	}
	st.sum = st.lChild.sum + st.rChild.sum
}

func (st *SegmentTree) Update(i int, value int) {
	st.operator(i, func(st *SegmentTree) {
		st.sum = value
	})
}

func (st *SegmentTree) Add(i int, add int) {
	st.operator(i, func(st *SegmentTree) {
		st.sum += add
	})
}
