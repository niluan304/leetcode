package main

type NumArray struct {
	nums []int
	tree []int
}

func Constructor(nums []int) NumArray {
	var n = len(nums)
	tree := make([]int, n+1) // 树状数组，以 1 开始，以 n 结尾
	sum := make([]int, n+1)  // 前缀和数组
	for i := 1; i <= n; i++ {
		sum[i] = sum[i-1] + nums[i-1]
		j := i - lowbit(i)
		tree[i] = sum[i] - sum[j]
	}
	return NumArray{
		nums: nums,
		tree: tree,
	}
}

func Constructor2(nums []int) NumArray {
	tree := make([]int, len(nums)+1)
	for i, x := range nums {
		i++
		tree[i] += x
		if nxt := i + i&-i; nxt < len(tree) {
			tree[nxt] += tree[i]
		}
	}
	return NumArray{nums, tree}
}

func (a *NumArray) Update(index, val int) {
	delta := val - a.nums[index]
	a.nums[index] = val
	for i := index + 1; i < len(a.tree); i += lowbit(i) {
		a.tree[i] += delta
	}
}

func (a *NumArray) prefixSum(i int) (s int) {
	for ; i > 0; i &= i - 1 { // i -= i & -i 的另一种写法
		s += a.tree[i]
	}
	return
}

func (a *NumArray) SumRange(left, right int) int {
	return a.prefixSum(right+1) - a.prefixSum(left)
}

func lowbit(x int) int {
	return x & -x
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */

func Constructor3(nums []int) NumArray3 { return NumArray3{st: NewSegmentTree(nums)} }

type NumArray3 struct{ st *SegmentTree }

func (a *NumArray3) Update(index, val int)        { a.st.Update(index+1, val) }
func (a *NumArray3) SumRange(left, right int) int { return a.st.Query(left+1, right+1) }

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
