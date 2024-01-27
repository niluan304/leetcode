package main

// 前缀和 + 哈希表
// - 时间复杂度：$\mathcal{O}(n)$。
// - 空间复杂度：$\mathcal{O}(n)$。
func checkSubarraySum(nums []int, k int) bool {
	var n = len(nums)
	var preSum = make([]int, n+1)
	for i, num := range nums {
		preSum[i+1] = preSum[i] + num
	}

	count := map[int]int{}
	for i, sum := range preSum {
		sum = ((sum % k) + k) % k
		if j, ok := count[sum]; ok {
			if i-j >= 2 {
				return true
			}
		} else {
			count[sum] = i
		}
	}

	return false
}

// checkSubarraySum 的空间优化版本
//
// 前缀和 + 哈希表
// - 时间复杂度：$\mathcal{O}(n)$。
// - 空间复杂度：$\mathcal{O}({\min(n,k)})$。
func checkSubarraySum2(nums []int, k int) bool {
	var sum = 0
	count := map[int]int{
		0: -1, // 初始化 preSum[0] = 0
	}

	for i, num := range nums {
		sum += num
		sum = ((sum % k) + k) % k // 负数取模

		if j, ok := count[sum]; ok {
			if i-j >= 2 {
				return true
			}
		} else {
			count[sum] = i // 记录余数首次出现的下标
		}
	}
	return false
}
