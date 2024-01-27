package main

// 前缀和 + 哈希表
// - 时间复杂度：$\mathcal{O}(n)$。
// - 空间复杂度：$\mathcal{O}(n)$。
func subarraySum(nums []int, k int) int {
	var n = len(nums)
	var preSum = make([]int, n+1)
	for i, num := range nums {
		preSum[i+1] = preSum[i] + num
	}

	var count = map[int]int{}
	var ans = 0
	for _, sum := range preSum {
		ans += count[sum-k] // 子数组是数组中元素的连续非空序列
		count[sum]++
	}
	return ans
}
