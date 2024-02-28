package main

// dfs + 贪心
//
// 这和线段树很像
//
// - 时间复杂度：O(n)。
// - 空间复杂度：O(n)。
func minIncrements(n int, cost []int) int {
	ans := 0
	var dfs func(i int) int
	dfs = func(i int) int {
		if i > n {
			return 0
		}

		left, right := dfs(2*i), dfs(2*i+1)
		if left < right {
			right, left = left, right // left, right = max(left, right), min(left, right)
		}
		ans += left - right
		return cost[i-1] + left
	}
	dfs(1) // 满二叉树的下标从 1 开始
	return ans
}

// 将 minIncrements 翻译为递推
// - 时间复杂度：O(n)。
// - 空间复杂度：O(1)。
func minIncrements2(n int, cost []int) int {
	ans := 0
	for i := n / 2; i > 0; i-- {
		left, right := cost[2*i-1], cost[2*i]
		if left < right {
			left, right = right, left // left, right = max(left, right), min(left, right)
		}
		ans += left - right
		cost[i-1] += left // 原地修改输入数组，也可以额外开一个数组
	}
	return ans
}
