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
