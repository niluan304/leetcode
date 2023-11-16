package main

// 方法一：递归实现组合型枚举
// 思路与算法
// 从 n 个当中选 k 个的所有方案对应的枚举是组合型枚举。在「方法一」中我们用递归来实现组合型枚举。
// 首先我们先回忆一下如何用递归实现二进制枚举（子集枚举），假设我们需要找到一个长度为 n 的序列 a 的所有子序列，代码框架是这样的：
func leetcode1(n int, k int) (ans [][]int) {
	temp := []int{}
	var dfs func(int)
	dfs = func(cur int) {
		// 剪枝：temp 长度加上区间 [cur, n] 的长度小于 k，不可能构造出长度为 k 的 temp
		if len(temp)+(n-cur+1) < k {
			return
		}
		// 记录合法的答案
		if len(temp) == k {
			comb := make([]int, k)
			copy(comb, temp)
			ans = append(ans, comb)
			return
		}
		// 考虑选择当前位置
		temp = append(temp, cur)
		dfs(cur + 1)
		temp = temp[:len(temp)-1]
		// 考虑不选择当前位置
		dfs(cur + 1)
	}
	dfs(1)
	return
}

// 非递归（字典序法）实现组合型枚举
// 思路与算法
// 小贴士：这个方法理解起来比「方法一」复杂，建议读者遇到不理解的地方可以在草稿纸上举例模拟这个过程。
// 这里的非递归版不是简单的用栈模拟递归转化为非递归：我们希望通过合适的手段，消除递归栈带来的额外空间代价。
func leetcode2(n int, k int) (ans [][]int) {
	// 初始化
	// 将 temp 中 [0, k - 1] 每个位置 i 设置为 i + 1，即 [0, k - 1] 存 [1, k]
	// 末尾加一位 n + 1 作为哨兵
	temp := []int{}
	for i := 1; i <= k; i++ {
		temp = append(temp, i)
	}
	temp = append(temp, n+1)

	for j := 0; j < k; {
		comb := make([]int, k)
		copy(comb, temp[:k])
		ans = append(ans, comb)
		// 寻找第一个 temp[j] + 1 != temp[j + 1] 的位置 t
		// 我们需要把 [0, t - 1] 区间内的每个位置重置成 [1, t]
		for j = 0; j < k && temp[j]+1 == temp[j+1]; j++ {
			temp[j] = j + 1
		}
		// j 是第一个 temp[j] + 1 != temp[j + 1] 的位置
		temp[j]++
	}
	return
}
