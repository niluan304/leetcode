package main

// 方法一：迭代法实现子集枚举
// 思路与算法
// 记原序列中元素的总数为 n。原序列中的每个数字 ai 的状态可能有两种，即「在子集中」和「不在子集中」。
// 我们用 1 表示「在子集中」，0 表示不在子集中，那么每一个子集可以对应一个长度为 n 的 0/1 序列，
// 第 i 位表示 ai 是否在子集中。
func leetcode(nums []int) (ans [][]int) {
	n := len(nums)
	for mask := 0; mask < 1<<n; mask++ {
		set := []int{}
		for i, v := range nums {
			if mask>>i&1 > 0 {
				set = append(set, v)
			}
		}
		ans = append(ans, append([]int(nil), set...))
	}
	return
}

// 方法二：递归法实现子集枚举
// 思路与算法
// 我们也可以用递归来实现子集枚举。
func leetcode2(nums []int) (ans [][]int) {
	set := []int{}
	var dfs func(int)
	dfs = func(cur int) {
		if cur == len(nums) {
			ans = append(ans, append([]int(nil), set...))
			return
		}
		set = append(set, nums[cur])
		dfs(cur + 1)
		set = set[:len(set)-1]
		dfs(cur + 1)
	}
	dfs(0)
	return
}
