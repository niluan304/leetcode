package main

// ### 思路
//
// 按照题目要求，把 $s$ 中的字母映射到数字上，设映射后变成了数组 $a$，那么题目相当于求 $a$ 的 [53. 最大子数组和](https://leetcode.cn/problems/maximum-subarray/)（允许子数组为空）。
//
// ### 做法一：dp
func maximumCostSubstring(s string, chars string, vals []int) int {
	var cost [26]int
	for i, _ := range cost {
		cost[i] = i + 1
	}
	for i, char := range chars {
		cost[char-'a'] = vals[i]
	}

	var ans = 0
	var n = len(s)
	var cache = make([]*int, n+1)
	var dfs func(i int) int
	dfs = func(i int) int {
		if i == 0 {
			return 0
		}

		ptr := &cache[i]
		if *ptr != nil {
			return **ptr
		}

		var v = cost[s[i-1]-'a'] + max(0, dfs(i-1))
		*ptr = &v
		ans = max(ans, v)
		return v
	}

	dfs(n)
	return ans
}

// 把字符串中的每个字符换成它们的价值，题目就是在问：
// > 求一个子数组，使得子数组的和最大。
func maximumCostSubstring2(s string, chars string, vals []int) int {
	var cost [26]int
	for i, _ := range cost {
		cost[i] = i + 1
	}
	for i, char := range chars {
		cost[char-'a'] = vals[i]
	}

	var ans, n = 0, len(s)
	var dp = make([]int, n+1)
	for i := 0; i < n; i++ {
		dp[i+1] = cost[s[i]-'a'] + max(0, dp[i])
		ans = max(ans, dp[i+1])
	}

	//// 压缩空间的写法
	//var ans, n = 0, len(s)
	//var dp = 0
	//for i := 0; i < n; i++ {
	//	dp = cost[s[i]-'a'] + max(0, dp)
	//	ans = max(ans, dp)
	//}
	return ans
}

// ### 思路
//
// 按照题目要求，把 $s$ 中的字母映射到数字上，设映射后变成了数组 $a$，那么题目相当于求 $a$ 的 [53. 最大子数组和](https://leetcode.cn/problems/maximum-subarray/)（允许子数组为空）。
//
// ### 做法二：前缀和
// 求子数组的和，就想到了前缀和，从而想到找到最大的 $preSum[i]−preSum[j]$，就是题目结果了。
//
// 注意：
// 前缀和的理解下，[53. 最大子数组和] 又和 [121. 买卖股票的最佳时机] 有关联
//
// 由于子数组的元素和等于两个前缀和的差，所以求出 $\textit{nums}$ 的前缀和，问题就变成 [121. 买卖股票的最佳时机](https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/) 了。本题子数组不能为空，相当于一定要交易一次。
func maximumCostSubstring3(s string, chars string, vals []int) int {
	var cost [26]int
	for i, _ := range cost {
		cost[i] = i + 1
	}
	for i, char := range chars {
		cost[char-'a'] = vals[i]
	}

	var n = len(s)
	var prefixSum = make([]int, n+1)
	var mn, ans = 0, 0

	for i := 0; i < n; i++ {
		prefixSum[i+1] = prefixSum[i] + cost[s[i]-'a']
		ans = max(ans, prefixSum[i+1]-mn)
		mn = min(mn, prefixSum[i+1])
	}

	//// 压缩空间的写法
	//var preSum = 0
	//for i := 0; i < n; i++ {
	//	preSum += cost[s[i]-'a']
	//	ans = max(ans, preSum-mn)
	//	mn = min(mn, preSum)
	//}
	return ans
}
