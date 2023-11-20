package main

// 前缀和分解
// 时间复杂度：O(n)
// 空间复杂度：O(n)
// ## 思路
//
// 由于子数组的元素和等于两个前缀和的差，所以求出 $\textit{nums}$ 的前缀和，问题就变成 [121. 买卖股票的最佳时机](https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/) 了。本题子数组不能为空，相当于一定要交易一次。
//
// 我们可以一边遍历数组计算前缀和，一边维护前缀和的最小值（相当于股票最低价格），用当前的前缀和（卖出价格）减去前缀和的最小值（买入价格），就得到了以当前元素结尾的子数组和的最大值（利润），用它来更新答案的最大值（最大利润）。
//
// 请注意，由于题目要求子数组不能为空，应当先计算前缀和-最小前缀和，再更新最小前缀和。相当于不能在同一天买入股票又卖出股票。
//
// 如果先更新最小前缀和，再计算前缀和-最小前缀和，就会把空数组的元素和 $0$ 算入答案。
func maxSubArray(nums []int) int {
	var ans = nums[0]
	var mn = 0

	var n = len(nums)
	var preSum = make([]int, n+1)
	for i := 0; i < n; i++ {
		preSum[i+1] = preSum[i] + nums[i]
		ans = max(ans, preSum[i+1]-mn)
		mn = min(mn, preSum[i+1])
	}

	return ans
}

// 动态规划？
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func maxSubArray2(nums []int) int {
	var ans = nums[0]
	var n = len(nums)
	var dp = make([]int, n+1)
	for i, num := range nums {
		dp[i+1] = max(dp[i]+num, num) // 以 nums[i] 结尾的最大和是多少？
		ans = max(ans, dp[i+1])
	}
	return ans
}
